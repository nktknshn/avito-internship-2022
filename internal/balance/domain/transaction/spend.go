package transaction

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
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
	AccountID domainAccount.AccountID
	UserID    domain.UserID
	OrderID   domainAccount.OrderID
	ProductID domainProduct.ProductID
	Amount    domainAmount.AmountPositive
	Status    TransactionSpendStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransactionSpendReserved(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	orderID domainAccount.OrderID,
	productID domainProduct.ProductID,
	amount domainAmount.AmountPositive,
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
		return ErrTransactionStatusMismatch
	}
	t.Status = TransactionSpendStatusCanceled
	t.UpdatedAt = now
	return nil
}

func (t *TransactionSpend) Confirm(now time.Time) error {
	if t.Status != TransactionSpendStatusReserved {
		return ErrTransactionStatusMismatch
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

	_accountID, err := domainAccount.NewAccountID(accountID)
	if err != nil {
		return nil, err
	}

	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return nil, err
	}

	_orderID, err := domainAccount.NewOrderID(orderID)
	if err != nil {
		return nil, err
	}

	_productID, err := domainProduct.NewProductID(productID)
	if err != nil {
		return nil, err
	}

	_status, err := NewTransactionSpendStatus(status)
	if err != nil {
		return nil, err
	}

	_amount, err := domainAmount.NewPositive(amount)
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
