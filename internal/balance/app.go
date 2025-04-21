package main

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/pkg/metrics_prometheus"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
)

func NewApplication(ctx context.Context, cfg *config.Config) (*app.Application, error) {

	db, err := sqlx_pg.Connect(ctx, cfg.GetDatabase())

	if err != nil {
		return nil, err
	}

	trmFactory := trmsqlx.NewFactory(db, sql.NewSavePoint())

	trm, err := manager.New(trmFactory)
	if err != nil {
		return nil, err
	}

	accountsRepository := accounts_pg.NewAccountsRepository(db, trmsqlx.DefaultCtxGetter)
	transactionsRepository := transactions_pg.NewTransactionsRepository(db, trmsqlx.DefaultCtxGetter)

	getBalance := get_balance.NewGetBalanceUseCase(trm, accountsRepository)

	deposit := deposit.NewDepositUseCase(trm, accountsRepository, transactionsRepository)
	reserve := reserve.NewReserveUseCase(trm, accountsRepository, transactionsRepository)
	reserveCancel := reserve_cancel.NewReserveCancelUseCase(trm, accountsRepository, transactionsRepository)
	reserveConfirm := reserve_confirm.NewReserveConfirmUseCase(trm, accountsRepository, transactionsRepository)
	transfer := transfer.NewTransferUseCase(trm, accountsRepository, transactionsRepository)

	// metrics
	logger := logging.NewSlog()
	metricsClient, err := metrics_prometheus.NewMetricsPrometheus("app")

	if err != nil {
		return nil, err
	}

	return &app.Application{
		GetBalance:     decorator.DecorateQuery(getBalance, metricsClient, logger, "GetBalance"),
		Deposit:        decorator.DecorateCommand(deposit, metricsClient, logger, "Deposit"),
		Reserve:        decorator.DecorateCommand(reserve, metricsClient, logger, "Reserve"),
		ReserveCancel:  decorator.DecorateCommand(reserveCancel, metricsClient, logger, "ReserveCancel"),
		ReserveConfirm: decorator.DecorateCommand(reserveConfirm, metricsClient, logger, "ReserveConfirm"),
		Transfer:       decorator.DecorateCommand(transfer, metricsClient, logger, "Transfer"),
	}, nil
}
