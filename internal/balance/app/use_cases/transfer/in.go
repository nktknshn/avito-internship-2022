package transfer

import (
	domainUser "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type In struct {
	FromUserID domainUser.UserID
	ToUserID   domainUser.UserID
	Amount     domainAmount.AmountPositive
}

func NewInFromValues(fromUserID int64, toUserID int64, amount int64) (In, error) {
	_from, err := domainUser.NewUserID(fromUserID)
	if err != nil {
		return In{}, err
	}
	_to, err := domainUser.NewUserID(toUserID)
	if err != nil {
		return In{}, err
	}
	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{FromUserID: _from, ToUserID: _to, Amount: _amount}, nil
}
