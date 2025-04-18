package accounts_pg

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

func NewAccountsRepository(db *sqlx.DB, c *trmsqlx.CtxGetter) *AccountsRepository {
	return &AccountsRepository{db: db, getter: c}
}

type accountDTO struct {
	Id               int64 `db:"id"`
	UserId           int64 `db:"user_id"`
	BalanceAvailable int64 `db:"balance_available"`
	BalanceReserved  int64 `db:"balance_reserved"`
}

func fromAccountDTO(a *accountDTO) (*domain.Account, error) {
	id, err := domain.NewAccountID(a.Id)
	if err != nil {
		return nil, err
	}
	userID, err := domain.NewUserID(a.UserId)
	if err != nil {
		return nil, err
	}
	balanceAvailable, err := domain.NewAmount(a.BalanceAvailable)
	if err != nil {
		return nil, err
	}
	balanceReserved, err := domain.NewAmount(a.BalanceReserved)
	if err != nil {
		return nil, err
	}
	accountBalance, err := domain.NewAccountBalance(balanceAvailable, balanceReserved)
	if err != nil {
		return nil, err
	}
	return &domain.Account{
		ID:      id,
		UserID:  userID,
		Balance: accountBalance,
	}, nil
}

func toAccountDTO(a *domain.Account) (*accountDTO, error) {
	return &accountDTO{
		Id:               a.ID.Value(),
		UserId:           a.UserID.Value(),
		BalanceAvailable: a.Balance.Available.Value(),
		BalanceReserved:  a.Balance.Reserved.Value(),
	}, nil
}
