package app

import (
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
	"github.com/nktknshn/avito-internship-2022/internal/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/reserve_confirm"
)

type Repositories struct {
	AccountRepository     *accounts_pg.AccountsRepository
	TransactionRepository *transactions_pg.TransactionsRepository
}

type UseCases struct {
	GetBalance     decorator.UseCaseQueryHandler[get_balance.In, get_balance.Out]
	Deposit        decorator.UseCaseCommandHandler[deposit.In]
	Reserve        decorator.UseCaseCommandHandler[reserve.In]
	ReserveCancel  decorator.UseCaseCommandHandler[reserve_cancel.In]
	ReserveConfirm decorator.UseCaseCommandHandler[reserve_confirm.In]
}

type App struct {
	metrics      metrics.Metrics
	logger       logging.Logger
	repositories Repositories
	useCases     UseCases
}

func NewApp() (*App, error) {
	return &App{
		repositories: Repositories{},
		useCases:     UseCases{},
	}, nil
}
