package account

import (
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

var (
	ErrInsufficientBalance = errors.New("insufficent available balance")
	ErrInsufficientReserve = errors.New("insufficent reserved balance")
)

type AccountBalance struct {
	available amount.Amount
	reserved  amount.Amount
}

func NewAccountBalance(available amount.Amount, reserved amount.Amount) (AccountBalance, error) {
	return AccountBalance{
		available: available,
		reserved:  reserved,
	}, nil
}

func NewAccountBalanceFromValues(available int64, reserved int64) (AccountBalance, error) {
	_available, err := amount.NewAmount(available)
	if err != nil {
		return AccountBalance{}, err
	}
	_reserved, err := amount.NewAmount(reserved)
	if err != nil {
		return AccountBalance{}, err
	}
	return NewAccountBalance(_available, _reserved)
}

func NewAccountBalanceEmpty() AccountBalance {
	return AccountBalance{
		available: 0,
		reserved:  0,
	}
}

func (ac AccountBalance) GetAvailable() amount.Amount {
	return ac.available
}

func (ac AccountBalance) GetReserved() amount.Amount {
	return ac.reserved
}

func (ac AccountBalance) Deposit(a amount.AmountPositive) (AccountBalance, error) {
	return NewAccountBalance(ac.available.Add(a), ac.reserved)
}

func (ac AccountBalance) Reserve(a amount.AmountPositive) (AccountBalance, error) {
	if ac.available.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientBalance
	}

	newAvailable, err := ac.available.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}

	return NewAccountBalance(newAvailable, ac.reserved.Add(a))
}

func (ac AccountBalance) ReserveCancel(a amount.AmountPositive) (AccountBalance, error) {
	if ac.reserved.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientReserve
	}

	newReserved, err := ac.reserved.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}

	return NewAccountBalance(ac.available.Add(a), newReserved)
}

func (ac AccountBalance) ReserveConfirm(a amount.AmountPositive) (AccountBalance, error) {
	if ac.reserved.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientReserve
	}

	newReserved, err := ac.reserved.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}

	return NewAccountBalance(ac.available, newReserved)
}

// Withdraw withdraws an amount from the available balance
func (ac AccountBalance) Withdraw(a amount.AmountPositive) (AccountBalance, error) {
	if ac.available.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientBalance
	}
	newAvailable, err := ac.available.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}
	return NewAccountBalance(newAvailable, ac.reserved)
}
