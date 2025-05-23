package transaction

import (
	"time"

	"github.com/google/uuid"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type TransactionTransferID uuid.UUID

func NewTransactionTransferID(id uuid.UUID) (TransactionTransferID, error) {
	return TransactionTransferID(id), nil
}
func (id TransactionTransferID) Value() uuid.UUID {
	return uuid.UUID(id)
}

type TransactionTransferStatus string

func NewTransactionTransferStatus(s string) (TransactionTransferStatus, error) {
	switch s {
	case TransactionTransferStatusConfirmed.Value():
		return TransactionTransferStatusConfirmed, nil
	default:
		return "", ErrInvalidTransactionStatus
	}
}

func (s TransactionTransferStatus) Value() string {
	return string(s)
}

const (
	TransactionTransferStatusConfirmed TransactionTransferStatus = "confirmed"
)

type TransactionTransfer struct {
	ID            TransactionTransferID
	FromAccountID domainAccount.AccountID
	ToAccountID   domainAccount.AccountID
	Amount        domainAmount.AmountPositive
	Status        TransactionTransferStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewTransactionTransfer(
	fromAccountID domainAccount.AccountID,
	toAccountID domainAccount.AccountID,
	amount domainAmount.AmountPositive,
	now time.Time,
) (*TransactionTransfer, error) {
	return &TransactionTransfer{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
		Status:        TransactionTransferStatusConfirmed,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

func NewTransactionTransferFromValues(
	id uuid.UUID,
	fromAccountID int64,
	toAccountID int64,
	amount int64,
	status string,
	createdAt time.Time,
	updatedAt time.Time,
) (*TransactionTransfer, error) {

	_id, err := NewTransactionTransferID(id)
	if err != nil {
		return nil, err
	}

	_fromAccountID, err := domainAccount.NewAccountID(fromAccountID)
	if err != nil {
		return nil, err
	}

	_toAccountID, err := domainAccount.NewAccountID(toAccountID)
	if err != nil {
		return nil, err
	}

	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return nil, err
	}

	_status, err := NewTransactionTransferStatus(status)
	if err != nil {
		return nil, err
	}

	return &TransactionTransfer{
		ID:            _id,
		FromAccountID: _fromAccountID,
		ToAccountID:   _toAccountID,
		Amount:        _amount,
		Status:        _status,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}
