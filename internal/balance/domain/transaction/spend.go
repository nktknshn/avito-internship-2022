package transaction

import (
	"time"

	"github.com/google/uuid"
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
	TransactionSpendStatusReserved  TransactionSpendStatus = "reserved"
	TransactionSpendStatusCanceled  TransactionSpendStatus = "canceled"
	TransactionSpendStatusConfirmed TransactionSpendStatus = "confirmed"
)

func NewTransactionSpendStatus(status string) (TransactionSpendStatus, error) {
	switch TransactionSpendStatus(status) {
	case TransactionSpendStatusReserved, TransactionSpendStatusCanceled, TransactionSpendStatusConfirmed:
		return TransactionSpendStatus(status), nil
	default:
		return "", ErrInvalidTransactionStatus
	}
}

type TransactionSpendID uuid.UUID

func (id TransactionSpendID) Value() uuid.UUID {
	return uuid.UUID(id)
}

func NewTransactionSpendID(id uuid.UUID) (TransactionSpendID, error) {
	return TransactionSpendID(id), nil
}

func NewTransactionSpendIDFromString(id string) (TransactionSpendID, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return TransactionSpendID{}, err
	}
	return TransactionSpendID(uuid), nil
}

type TransactionSpend struct {
	ID        TransactionSpendID
	AccountID domainAccount.AccountID
	UserID    domain.UserID
	OrderID   domain.OrderID
	ProductID domainProduct.ProductID
	Amount    domainAmount.AmountPositive
	Status    TransactionSpendStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransactionSpendReserved(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	orderID domain.OrderID,
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
	id uuid.UUID,
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

	_orderID, err := domain.NewOrderID(orderID)
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
