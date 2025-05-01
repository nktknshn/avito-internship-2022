package account

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrInsufficientAvailableBalance = domainError.New("insufficent available balance")
	ErrInsufficientReserveBalance   = domainError.New("insufficent reserved balance")
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

func NewAccountBalanceZero() AccountBalance {
	return AccountBalance{
		available: amount.Zero(),
		reserved:  amount.Zero(),
	}
}

func NewAccountBalanceFromValues(available int64, reserved int64) (AccountBalance, error) {
	_available, err := amount.New(available)
	if err != nil {
		return AccountBalance{}, err
	}
	_reserved, err := amount.New(reserved)
	if err != nil {
		return AccountBalance{}, err
	}
	return NewAccountBalance(_available, _reserved)
}

func NewAccountBalanceEmpty() AccountBalance {
	return AccountBalance{
		available: amount.Zero(),
		reserved:  amount.Zero(),
	}
}

func (ac AccountBalance) GetAvailable() amount.Amount {
	return ac.available
}

func (ac AccountBalance) GetReserved() amount.Amount {
	return ac.reserved
}

func (ac AccountBalance) Deposit(a amount.AmountPositive) (AccountBalance, error) {
	available, err := ac.available.Add(a)
	if err != nil {
		return AccountBalance{}, err
	}
	return NewAccountBalance(available, ac.reserved)
}

func (ac AccountBalance) Reserve(a amount.AmountPositive) (AccountBalance, error) {
	if ac.available.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientAvailableBalance
	}

	newAvailable, err := ac.available.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}

	reserved, err := ac.reserved.Add(a)
	if err != nil {
		return AccountBalance{}, err
	}

	return NewAccountBalance(newAvailable, reserved)
}

func (ac AccountBalance) ReserveCancel(a amount.AmountPositive) (AccountBalance, error) {
	if ac.reserved.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientReserveBalance
	}

	newReserved, err := ac.reserved.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}

	newAvailable, err := ac.available.Add(a)
	if err != nil {
		return AccountBalance{}, err
	}

	return NewAccountBalance(newAvailable, newReserved)
}

func (ac AccountBalance) ReserveConfirm(a amount.AmountPositive) (AccountBalance, error) {
	if ac.reserved.LessThanPositive(a) {
		return AccountBalance{}, ErrInsufficientReserveBalance
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
		return AccountBalance{}, ErrInsufficientAvailableBalance
	}
	newAvailable, err := ac.available.Sub(a)
	if err != nil {
		return AccountBalance{}, err
	}
	return NewAccountBalance(newAvailable, ac.reserved)
}
