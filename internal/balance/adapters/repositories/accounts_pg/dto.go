package accounts_pg

import (
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type accountDTO struct {
	Id               int64 `db:"id"`
	UserId           int64 `db:"user_id"`
	BalanceAvailable int64 `db:"balance_available"`
	BalanceReserved  int64 `db:"balance_reserved"`
}

func fromAccountDTO(a *accountDTO) (*domainAccount.Account, error) {
	return domainAccount.NewAccountFromValues(
		a.Id,
		a.UserId,
		a.BalanceAvailable,
		a.BalanceReserved,
	)
}

func toAccountDTO(a *domainAccount.Account) (*accountDTO, error) {
	return &accountDTO{
		Id:               a.ID.Value(),
		UserId:           a.UserID.Value(),
		BalanceAvailable: a.Balance.GetAvailable().Value(),
		BalanceReserved:  a.Balance.GetReserved().Value(),
	}, nil
}
