package domain

import (
	"errors"
)

type AccountBalance struct {
	Available Amount
	Reserved  Amount
}

func NewAccountBalance(available Amount, reserved Amount) (AccountBalance, error) {
	return AccountBalance{
		Available: available,
		Reserved:  reserved,
	}, nil
}

func NewAccountBalanceEmpty() AccountBalance {
	return AccountBalance{
		Available: 0,
		Reserved:  0,
	}
}

var (
	ErrInsufficientBalance = errors.New("insufficent available balance")
	ErrInsufficientReserve = errors.New("insufficent reserved balance")
)

func (ac AccountBalance) Deposit(amount AmountPositive) (AccountBalance, error) {
	return AccountBalance{
		Available: ac.Available.Add(amount),
		Reserved:  ac.Reserved,
	}, nil
}

func (ac AccountBalance) Reserve(amount AmountPositive) (AccountBalance, error) {
	if ac.Available.Value() < amount.Value() {
		return AccountBalance{}, ErrInsufficientBalance
	}
	return AccountBalance{
		Available: ac.Available.Sub(amount),
		Reserved:  ac.Reserved.Add(amount),
	}, nil
}

func (ac AccountBalance) ReserveCancel(amount AmountPositive) (AccountBalance, error) {
	if ac.Reserved.Value() < amount.Value() {
		return AccountBalance{}, ErrInsufficientReserve
	}

	return AccountBalance{
		Available: ac.Available.Add(amount),
		Reserved:  ac.Reserved.Sub(amount),
	}, nil
}

func (ac AccountBalance) ReserveConfirm(amount AmountPositive) (AccountBalance, error) {
	if ac.Reserved.Value() < amount.Value() {
		return AccountBalance{}, ErrInsufficientReserve
	}

	return AccountBalance{
		Available: ac.Available,
		Reserved:  ac.Reserved.Sub(amount),
	}, nil
}
