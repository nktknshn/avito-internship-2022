package domain

import "errors"

type OrderID int64

func (o OrderID) Value() int64 {
	return int64(o)
}

var (
	ErrInvalidOrderID = errors.New("invalid order id")
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
	UserID  UserID
	Balance AccountBalance
}

func NewAccount(userID UserID) (*Account, error) {
	return &Account{
		UserID:  userID,
		Balance: NewAccountBalanceEmpty(),
	}, nil
}

func (ac *Account) SetBalance(ab AccountBalance) error {
	ac.Balance = ab
	return nil
}

func (ac *Account) BalanceDeposit(amount AmountPositive) error {
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

func (ac *Account) BalanceReserve(amount AmountPositive) error {
	newBalance, err := ac.Balance.Reserve(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveCancel(amount AmountPositive) error {
	newBalance, err := ac.Balance.ReserveCancel(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveConfirm(amount AmountPositive) error {
	newBalance, err := ac.Balance.ReserveConfirm(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}
