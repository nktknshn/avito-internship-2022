package app_impl

import (
	"context"
	"net/http"

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
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/pkg/metrics_prometheus"
	"github.com/nktknshn/avito-internship-2022/pkg/password_hasher_argon"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
	"github.com/nktknshn/avito-internship-2022/pkg/token_generator_jwt"
)

type Application struct {
	app.Application
	MetricsHandler http.Handler
	Logger         logging.Logger
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

	// balance
	var (
		accountsRepository     = accounts_pg.New(db, trmsqlx.DefaultCtxGetter)
		transactionsRepository = transactions_pg.New(db, trmsqlx.DefaultCtxGetter)

		getBalance         = get_balance.New(trm, accountsRepository)
		deposit            = deposit.New(trm, accountsRepository, transactionsRepository)
		reserve            = reserve.New(trm, accountsRepository, transactionsRepository)
		reserveCancel      = reserve_cancel.New(trm, accountsRepository, transactionsRepository)
		reserveConfirm     = reserve_confirm.New(trm, accountsRepository, transactionsRepository)
		transfer           = transfer.New(trm, accountsRepository, transactionsRepository)
		reportTransactions = report_transactions.New(transactionsRepository)
		reportRevenue      = report_revenue.New(transactionsRepository)
	)

	// logs & metrics
	logger := logging.NewSlog()
	metricsClient, err := metrics_prometheus.NewMetricsPrometheus("app_balance")

	if err != nil {
		return nil, err
	}

	return &Application{
		Application: app.Application{
			// auth
			AuthSignin:        decorator.DecorateQuery(authSignin, metricsClient, logger),
			AuthSignup:        decorator.DecorateCommand(authSignup, metricsClient, logger),
			AuthValidateToken: decorator.DecorateQuery(authValidateToken, metricsClient, logger),
			// balance
			GetBalance:         decorator.DecorateQuery(getBalance, metricsClient, logger),
			Deposit:            decorator.DecorateCommand(deposit, metricsClient, logger),
			Reserve:            decorator.DecorateCommand(reserve, metricsClient, logger),
			ReserveCancel:      decorator.DecorateCommand(reserveCancel, metricsClient, logger),
			ReserveConfirm:     decorator.DecorateCommand(reserveConfirm, metricsClient, logger),
			Transfer:           decorator.DecorateCommand(transfer, metricsClient, logger),
			ReportTransactions: decorator.DecorateQuery(reportTransactions, metricsClient, logger),
			ReportRevenue:      decorator.DecorateQuery(reportRevenue, metricsClient, logger),
		},
		MetricsHandler: metricsClient.GetHandler(),
		Logger:         logger,
	}, nil
}
