package accounts_pg

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	domain "github.com/nktknshn/avito-internship-2022/internal/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/domain/amount"
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

func fromAccountDTO(a *accountDTO) (*domainAccount.Account, error) {
	id, err := domainAccount.NewAccountID(a.Id)
	if err != nil {
		return nil, err
	}
	userID, err := domain.NewUserID(a.UserId)
	if err != nil {
		return nil, err
	}
	balanceAvailable, err := domainAmount.NewAmount(a.BalanceAvailable)
	if err != nil {
		return nil, err
	}
	balanceReserved, err := domainAmount.NewAmount(a.BalanceReserved)
	if err != nil {
		return nil, err
	}
	accountBalance, err := domainAccount.NewAccountBalance(balanceAvailable, balanceReserved)
	if err != nil {
		return nil, err
	}

	return &domainAccount.Account{
		ID:      id,
		UserID:  userID,
		Balance: accountBalance,
	}, nil
}

func toAccountDTO(a *domainAccount.Account) (*accountDTO, error) {
	return &accountDTO{
		Id:               a.ID.Value(),
		UserId:           a.UserID.Value(),
		BalanceAvailable: a.Balance.GetAvailable().Value(),
		BalanceReserved:  a.Balance.GetReserved().Value(),
	}, nil
}
