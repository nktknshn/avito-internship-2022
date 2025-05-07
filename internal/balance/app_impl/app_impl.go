package app_impl

import (
	"context"
	"net/http"
	"time"

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

	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
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

func NewApplicationFromDeps(ctx context.Context, deps *AppDeps) (*Application, error) {

	// balance
	var (
		authSignin        = auth_signin.New(deps.Trm, deps.PasswordHasher, deps.TokenGenerator, deps.Repositories.AuthRepository)
		authSignup        = auth_signup.New(deps.Trm, deps.PasswordHasher, deps.Repositories.AuthRepository)
		authValidateToken = auth_validate_token.New(deps.Trm, deps.TokenGenerator, deps.Repositories.AuthRepository)

		getBalance          = get_balance.New(deps.Trm, deps.Repositories.AccountsRepository)
		deposit             = deposit.New(deps.Trm, deps.Repositories.AccountsRepository, deps.Repositories.TransactionsRepository)
		reserve             = reserve.New(deps.Trm, deps.Repositories.AccountsRepository, deps.Repositories.TransactionsRepository)
		reserveCancel       = reserve_cancel.New(deps.Trm, deps.Repositories.AccountsRepository, deps.Repositories.TransactionsRepository)
		reserveConfirm      = reserve_confirm.New(deps.Trm, deps.Repositories.AccountsRepository, deps.Repositories.TransactionsRepository)
		transfer            = transfer.New(deps.Trm, deps.Repositories.AccountsRepository, deps.Repositories.TransactionsRepository)
		reportTransactions  = report_transactions.New(deps.Repositories.TransactionsRepository)
		reportRevenue       = report_revenue.New(deps.Repositories.TransactionsRepository)
		reportRevenueExport = report_revenue_export.New(deps.FileExporter, deps.Repositories.TransactionsRepository)
	)

	exporterCleanup := func() {
		err := deps.FileExporter.Cleanup()
		if err != nil {
			deps.Logger.Error("exporter.Cleanup()", "error", err)
		}
	}

	return &Application{
		app: &app.Application{
			// auth
			AuthSignin:        decorator.Decorate1(authSignin, deps.MetricsClient, deps.Logger),
			AuthSignup:        decorator.Decorate0(authSignup, deps.MetricsClient, deps.Logger),
			AuthValidateToken: decorator.Decorate1(authValidateToken, deps.MetricsClient, deps.Logger),
			// balance
			GetBalance:          decorator.Decorate1(getBalance, deps.MetricsClient, deps.Logger),
			Deposit:             decorator.Decorate0(deposit, deps.MetricsClient, deps.Logger),
			Reserve:             decorator.Decorate0(reserve, deps.MetricsClient, deps.Logger),
			ReserveCancel:       decorator.Decorate0(reserveCancel, deps.MetricsClient, deps.Logger),
			ReserveConfirm:      decorator.Decorate0(reserveConfirm, deps.MetricsClient, deps.Logger),
			Transfer:            decorator.Decorate0(transfer, deps.MetricsClient, deps.Logger),
			ReportTransactions:  decorator.Decorate1(reportTransactions, deps.MetricsClient, deps.Logger),
			ReportRevenue:       decorator.Decorate1(reportRevenue, deps.MetricsClient, deps.Logger),
			ReportRevenueExport: decorator.Decorate1(reportRevenueExport, deps.MetricsClient, deps.Logger),
		},
		metricsHandler:         deps.MetricsClient.GetHandler(),
		logger:                 deps.Logger,
		revenueExporterHandler: deps.FileExporter.GetHandler(),
		revenueExporterCleanup: exporterCleanup,
	}, nil
}

// NewApplication создает новую реализацию приложения для использования в адаптерах
func NewApplication(ctx context.Context, cfg *config.Config) (*Application, error) {
	deps, err := NewDeps(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return NewApplicationFromDeps(ctx, deps)
}
