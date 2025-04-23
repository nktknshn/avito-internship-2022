package transfer

import (
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type In struct {
	From   domainAccount.AccountID
	To     domainAccount.AccountID
	Amount domainAmount.AmountPositive
}

func NewInFromValues(from int64, to int64, amount int64) (In, error) {
	_from, err := domainAccount.NewAccountID(from)
	if err != nil {
		return In{}, err
	}
	_to, err := domainAccount.NewAccountID(to)
	if err != nil {
		return In{}, err
	}
	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{From: _from, To: _to, Amount: _amount}, nil
}
