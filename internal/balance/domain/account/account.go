package account

import (
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type OrderID int64

func (o OrderID) Value() int64 {
	return int64(o)
}

var (
	ErrInvalidOrderID = errors.New("invalid order id")
	ErrSameAccount    = errors.New("same account")
)

func NewOrderID(id int64) (OrderID, error) {
	if id <= 0 {
		return 0, ErrInvalidOrderID
	}
	return OrderID(id), nil
}

type AccountID int64

func (a AccountID) Value() int64 {
	return int64(a)
}

var (
	ErrInvalidAccountID = errors.New("invalid account id")
)

func NewAccountID(id int64) (AccountID, error) {
	if id <= 0 {
		return 0, ErrInvalidAccountID
	}
	return AccountID(id), nil
}

type Account struct {
	ID      AccountID
	UserID  domain.UserID
	Balance AccountBalance
}

func NewAccount(userID domain.UserID) (*Account, error) {
	return &Account{
		UserID:  userID,
		Balance: NewAccountBalanceEmpty(),
	}, nil
}

func (ac *Account) SetBalance(ab AccountBalance) error {
	ac.Balance = ab
	return nil
}

func (ac *Account) BalanceDeposit(amount amount.AmountPositive) error {
	newBalance, err := ac.Balance.Deposit(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserve(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.Reserve(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveCancel(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.ReserveCancel(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveConfirm(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.ReserveConfirm(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceWithdraw(amount amount.AmountPositive) error {
	newBalance, err := ac.Balance.Withdraw(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) Transfer(to *Account, amount amount.AmountPositive) error {
	if ac.ID == to.ID {
		return ErrSameAccount
	}
	err := ac.BalanceWithdraw(amount)
	if err != nil {
		return err
	}
	err = to.BalanceDeposit(amount)
	if err != nil {
		return err
	}
	return nil
}
