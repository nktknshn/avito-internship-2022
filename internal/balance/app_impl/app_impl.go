package app_impl

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/pkg/metrics_prometheus"
	"github.com/nktknshn/avito-internship-2022/pkg/password_hasher_argon"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
)

// NewApplication создает новую реализацию приложения
func NewApplication(ctx context.Context, cfg *config.Config) (*app.Application, error) {

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

	var (
		passwordHasher = password_hasher_argon.NewHasher()
	)

	// auth
	var (
		accountsRepository     = accounts_pg.New(db, trmsqlx.DefaultCtxGetter)
		transactionsRepository = transactions_pg.New(db, trmsqlx.DefaultCtxGetter)
		authRepository         = auth_pg.NewAuthRepository(db, trmsqlx.DefaultCtxGetter)
	)

	// account
	var (
		getBalance     = get_balance.New(trm, accountsRepository)
		deposit        = deposit.New(trm, accountsRepository, transactionsRepository)
		reserve        = reserve.New(trm, accountsRepository, transactionsRepository)
		reserveCancel  = reserve_cancel.New(trm, accountsRepository, transactionsRepository)
		reserveConfirm = reserve_confirm.New(trm, accountsRepository, transactionsRepository)
		transfer       = transfer.New(trm, accountsRepository, transactionsRepository)

		authSignup        = auth_signup.New(trm, passwordHasher, authRepository)
		authValidateToken = auth_validate_token.New(trm, authRepository)
	)

	// metrics
	logger := logging.NewSlog()
	metricsClient, err := metrics_prometheus.NewMetricsPrometheus("app")

	if err != nil {
		return nil, err
	}

	return &app.Application{
		// auth
		AuthValidateToken: decorator.DecorateQuery(authValidateToken, metricsClient, logger, "AuthValidateToken"),
		AuthSignup:        decorator.DecorateCommand(authSignup, metricsClient, logger, "AuthSignup"),
		//
		GetBalance:     decorator.DecorateQuery(getBalance, metricsClient, logger, "GetBalance"),
		Deposit:        decorator.DecorateCommand(deposit, metricsClient, logger, "Deposit"),
		Reserve:        decorator.DecorateCommand(reserve, metricsClient, logger, "Reserve"),
		ReserveCancel:  decorator.DecorateCommand(reserveCancel, metricsClient, logger, "ReserveCancel"),
		ReserveConfirm: decorator.DecorateCommand(reserveConfirm, metricsClient, logger, "ReserveConfirm"),
		Transfer:       decorator.DecorateCommand(transfer, metricsClient, logger, "Transfer"),
	}, nil
}
