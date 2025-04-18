package domain

import (
	"errors"
	"time"
)

type TransactionSpendStatus string

func (s TransactionSpendStatus) Value() string {
	return string(s)
}

const (
	TransactionSpendStatusReserved  = "reserved"
	TransactionSpendStatusCanceled  = "canceled"
	TransactionSpendStatusConfirmed = "confirmed"
)

func NewTransactionSpendStatus(status string) (TransactionSpendStatus, error) {
	switch status {
	case TransactionSpendStatusReserved, TransactionSpendStatusCanceled, TransactionSpendStatusConfirmed:
		return TransactionSpendStatus(status), nil
	default:
		return "", ErrInvalidTransactionStatus
	}
}

type TransactionSpendID int64

func (id TransactionSpendID) Value() int64 {
	return int64(id)
}

func NewTransactionSpendID(id int64) (TransactionSpendID, error) {
	if id < 0 {
		return 0, ErrInvalidTransactionID
	}
	return TransactionSpendID(id), nil
}

type TransactionSpend struct {
	ID        TransactionSpendID
	AccountID AccountID
	UserID    UserID
	OrderID   OrderID
	ProductID ProductID
	Amount    AmountPositive
	Status    TransactionSpendStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransactionSpendReserved(
	accountID AccountID,
	userID UserID,
	orderID OrderID,
	productID ProductID,
	amount AmountPositive,
	now time.Time,
) (*TransactionSpend, error) {
	return &TransactionSpend{
		AccountID: accountID,
		UserID:    userID,
		OrderID:   orderID,
		ProductID: productID,
		Amount:    amount,
		Status:    TransactionSpendStatusReserved,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (t *TransactionSpend) Cancel(now time.Time) error {
	if t.Status != TransactionSpendStatusReserved {
		return errors.New("cannot cancel")
	}
	t.Status = TransactionSpendStatusCanceled
	t.UpdatedAt = now
	return nil
}

func (t *TransactionSpend) Confirm(now time.Time) error {
	if t.Status != TransactionSpendStatusReserved {
		return errors.New("cannot cancel")
	}
	t.Status = TransactionSpendStatusConfirmed
	t.UpdatedAt = now
	return nil
}

func NewTransactionSpendFromValues(
	id int64,
	accountID int64,
	userID int64,
	orderID int64,
	productID int64,
	amount int64,
	status string,
	createdAt time.Time,
	updatedAt time.Time,
) (*TransactionSpend, error) {
	_id, err := NewTransactionSpendID(id)
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

	_status, err := NewTransactionSpendStatus(status)
	if err != nil {
		return nil, err
	}

	_amount, err := NewAmountPositive(amount)
	if err != nil {
		return nil, err
	}

	return &TransactionSpend{
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
