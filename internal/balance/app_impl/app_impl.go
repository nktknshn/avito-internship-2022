package app_impl

import (
	"context"
	"net/http"
	"time"

	"github.com/avito-tech/go-transaction-manager/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/pkg/file_exporter_http"
	"github.com/nktknshn/avito-internship-2022/pkg/metrics_prometheus"
	"github.com/nktknshn/avito-internship-2022/pkg/password_hasher_argon"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
	"github.com/nktknshn/avito-internship-2022/pkg/token_generator_jwt"
)

type Application struct {
	app            *app.Application
	metricsHandler http.Handler
	logger         logging.Logger

	config                 *config.Config
	revenueExporterCleanup func()
	revenueExporterHandler http.Handler
}

func (a *Application) GetApp() *app.Application {
	return a.app
}

func (a *Application) GetRevenueExporterHandler() http.Handler {
	return a.revenueExporterHandler
}

func (a *Application) GetMetricsHandler() http.Handler {
	return a.metricsHandler
}

func (a *Application) GetLogger() logging.Logger {
	return a.logger
}

func (a *Application) GetRevenueExporterCleanup() func() {
	return a.revenueExporterCleanup
}

func (a *Application) GetConfig() *config.Config {
	return a.config
}

func (a *Application) RunRevenueExporterCleanup(ctx context.Context) {
	a.logger.Info("Running revenue exporter cleanup goroutine")
	go func() {
		for {
			select {
			case <-ctx.Done():
				a.logger.Info("Revenue exporter cleanup goroutine finished")
				return
			case <-time.After(a.config.GetUseCases().GetReportRevenueExport().GetTTL()):
				a.logger.Info("Running revenue exporter cleanup")
				a.revenueExporterCleanup()
			}
		}
	}()
}

// NewApplication создает новую реализацию приложения для использования в адаптерах
func NewApplication(ctx context.Context, cfg *config.Config) (*Application, error) {

	db, err := sqlx_pg.Connect(ctx, cfg.GetPostgres())
	if err != nil {
		return nil, err
	}

	err = sqlx_pg.Migrate(ctx, db.DB, cfg.GetPostgres().GetMigrationsDir())
	if err != nil {
		return nil, err
	}

	trmFactory := trmsqlx.NewFactory(db, sql.NewSavePoint())
	trm, err := manager.New(trmFactory)
	if err != nil {
		return nil, err
	}

	// logs & metrics
	logger := logging.NewSlog()
	metricsClient, err := metrics_prometheus.NewMetricsPrometheus("app_balance")

	if err != nil {
		return nil, err
	}

	// auth
	var (
		passwordHasher = password_hasher_argon.New()
		tokenGenerator = token_generator_jwt.NewTokenGeneratorJWT[domainAuth.AuthUserTokenClaims](
			[]byte(cfg.GetJWT().GetSecret()),
			cfg.GetJWT().GetTTL(),
		)

		tokenValidator = token_generator_jwt.NewTokenValidatorJWT[domainAuth.AuthUserTokenClaims](
			[]byte(cfg.GetJWT().GetSecret()),
		)

		authRepository = auth_pg.New(db, trmsqlx.DefaultCtxGetter)

		authSignup        = auth_signup.New(trm, passwordHasher, authRepository)
		authSignin        = auth_signin.New(trm, passwordHasher, tokenGenerator, authRepository)
		authValidateToken = auth_validate_token.New(trm, tokenValidator, authRepository)
	)

	exporter, err := file_exporter_http.New(
		file_exporter_http.Config{
			Folder: cfg.GetUseCases().GetReportRevenueExport().GetFolder(),
			TTL:    cfg.GetUseCases().GetReportRevenueExport().GetTTL(),
			URL:    cfg.GetUseCases().GetReportRevenueExport().GetURL(),
			Zip:    cfg.GetUseCases().GetReportRevenueExport().GetZip(),
		},
	)

	if err != nil {
		return nil, err
	}

	exporterCleanup := func() {
		err := exporter.Cleanup()
		if err != nil {
			logger.Error("exporter.Cleanup()", "error", err)
		}
	}

	// balance
	var (
		accountsRepository     = accounts_pg.New(db, trmsqlx.DefaultCtxGetter)
		transactionsRepository = transactions_pg.New(db, trmsqlx.DefaultCtxGetter)

		getBalance          = get_balance.New(trm, accountsRepository)
		deposit             = deposit.New(trm, accountsRepository, transactionsRepository)
		reserve             = reserve.New(trm, accountsRepository, transactionsRepository)
		reserveCancel       = reserve_cancel.New(trm, accountsRepository, transactionsRepository)
		reserveConfirm      = reserve_confirm.New(trm, accountsRepository, transactionsRepository)
		transfer            = transfer.New(trm, accountsRepository, transactionsRepository)
		reportTransactions  = report_transactions.New(transactionsRepository)
		reportRevenue       = report_revenue.New(transactionsRepository)
		reportRevenueExport = report_revenue_export.New(exporter, transactionsRepository)
	)

	return &Application{
		app: &app.Application{
			// auth
			AuthSignin:        decorator.Decorate1(authSignin, metricsClient, logger),
			AuthSignup:        decorator.Decorate0(authSignup, metricsClient, logger),
			AuthValidateToken: decorator.Decorate1(authValidateToken, metricsClient, logger),
			// balance
			GetBalance:          decorator.Decorate1(getBalance, metricsClient, logger),
			Deposit:             decorator.Decorate0(deposit, metricsClient, logger),
			Reserve:             decorator.Decorate0(reserve, metricsClient, logger),
			ReserveCancel:       decorator.Decorate0(reserveCancel, metricsClient, logger),
			ReserveConfirm:      decorator.Decorate0(reserveConfirm, metricsClient, logger),
			Transfer:            decorator.Decorate0(transfer, metricsClient, logger),
			ReportTransactions:  decorator.Decorate1(reportTransactions, metricsClient, logger),
			ReportRevenue:       decorator.Decorate1(reportRevenue, metricsClient, logger),
			ReportRevenueExport: decorator.Decorate1(reportRevenueExport, metricsClient, logger),
		},
		config:                 cfg,
		metricsHandler:         metricsClient.GetHandler(),
		logger:                 logger,
		revenueExporterHandler: exporter.GetHandler(),
		revenueExporterCleanup: exporterCleanup,
	}, nil
}
