package deposit

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/domain/amount"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type In struct {
	UserID domain.UserID
	Amount domainAmount.AmountPositive
	Source domainTransaction.DepositSource
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

	_amount, err := domainAmount.NewAmountPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		UserID: _userID,
		Source: _source,
		Amount: _amount,
	}, nil
}
