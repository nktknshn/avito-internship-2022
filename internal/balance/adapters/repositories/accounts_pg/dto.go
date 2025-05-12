package accounts_pg

import (
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

type accountDTO struct {
	ID               int64 `db:"id"`
	UserID           int64 `db:"user_id"`
	BalanceAvailable int64 `db:"balance_available"`
	BalanceReserved  int64 `db:"balance_reserved"`
}

func fromAccountDTO(a *accountDTO) (*domainAccount.Account, error) {
	acc, err := domainAccount.NewAccountFromValues(
		a.ID,
		a.UserID,
		a.BalanceAvailable,
		a.BalanceReserved,
	)
	if err != nil {
		return nil, domainError.Strip(err)
	}
	return acc, nil
}

func toAccountDTO(a *domainAccount.Account) *accountDTO {
	return &accountDTO{
		ID:               a.ID.Value(),
		UserID:           a.UserID.Value(),
		BalanceAvailable: a.Balance.GetAvailable().Value(),
		BalanceReserved:  a.Balance.GetReserved().Value(),
	}
}
