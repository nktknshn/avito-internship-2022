package domain

import (
	"errors"
	"time"
)

type AccountTransactionSpendStatus string

func (s AccountTransactionSpendStatus) Value() string {
	return string(s)
}

const (
	AccountTransactionSpendStatusUnknown   = ""
	AccountTransactionSpendStatusReserved  = "reserved"
	AccountTransactionSpendStatusCanceled  = "canceled"
	AccountTransactionSpendStatusConfirmed = "confirmed"
)

func NewAccountTransactionSpendStatus(status string) (AccountTransactionSpendStatus, error) {
	switch status {
	case AccountTransactionSpendStatusUnknown, AccountTransactionSpendStatusReserved, AccountTransactionSpendStatusCanceled, AccountTransactionSpendStatusConfirmed:
		return AccountTransactionSpendStatus(status), nil
	default:
		return "", ErrInvalidAccountTransactionStatus
	}
}

type AccountTransactionSpend struct {
	ID        AccountTransactionSpendID
	AccountID AccountID
	UserID    UserID
	OrderID   OrderID
	ProductID ProductID
	Amount    AmountPositive
	Status    AccountTransactionSpendStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccountTransactionSpendReserved(
	accountID AccountID,
	userID UserID,
	orderID OrderID,
	productID ProductID,
	amount AmountPositive,
	now time.Time,
) (*AccountTransactionSpend, error) {
	return &AccountTransactionSpend{
		AccountID: accountID,
		UserID:    userID,
		OrderID:   orderID,
		ProductID: productID,
		Amount:    amount,
		Status:    AccountTransactionSpendStatusReserved,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (t *AccountTransactionSpend) Cancel(now time.Time) error {
	if t.Status != AccountTransactionSpendStatusReserved {
		return errors.New("cannot cancel")
	}
	t.Status = AccountTransactionSpendStatusCanceled
	t.UpdatedAt = now
	return nil
}

func (t *AccountTransactionSpend) Confirm(now time.Time) error {
	if t.Status != AccountTransactionSpendStatusReserved {
		return errors.New("cannot cancel")
	}
	t.Status = AccountTransactionSpendStatusConfirmed
	t.UpdatedAt = now
	return nil
}

func NewAccountTransactionSpendFromValues(
	id int64,
	accountID int64,
	userID int64,
	orderID int64,
	productID int64,
	amount int64,
	status string,
	createdAt time.Time,
	updatedAt time.Time,
) (*AccountTransactionSpend, error) {
	_id, err := NewAccountTransactionSpendID(id)
	if err != nil {
		return nil, err
	}

	_accountID, err := NewAccountID(accountID)
	if err != nil {
		return nil, err
	}

	_userID, err := NewUserID(userID)
	if err != nil {
		return nil, err
	}

	_orderID, err := NewOrderID(orderID)
	if err != nil {
		return nil, err
	}

	_productID, err := NewProductID(productID)
	if err != nil {
		return nil, err
	}

	_status, err := NewAccountTransactionSpendStatus(status)
	if err != nil {
		return nil, err
	}

	_amount, err := NewAmountPositive(amount)
	if err != nil {
		return nil, err
	}

	return &AccountTransactionSpend{
		ID:        _id,
		AccountID: _accountID,
		UserID:    _userID,
		OrderID:   _orderID,
		ProductID: _productID,
		Amount:    _amount,
		Status:    _status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
