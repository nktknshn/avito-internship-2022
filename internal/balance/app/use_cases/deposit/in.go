package deposit

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type In struct {
	userID domain.UserID
	amount domainAmount.AmountPositive
	source domainTransaction.TransactionDepositSource
}

func NewInFromValues(userID int64, amount int64, source string) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_source, err := domainTransaction.NewDepositSource(source)
	if err != nil {
		return In{}, err
	}

	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		userID: _userID,
		source: _source,
		amount: _amount,
	}, nil
}
