package transaction

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type TransactionTransferID int64

func (id TransactionTransferID) Value() int64 {
	return int64(id)
}

func NewTransactionTransferID(id int64) (TransactionTransferID, error) {
	if id < 0 {
		return 0, ErrInvalidTransactionID
	}
	return TransactionTransferID(id), nil
}

type TransactionTransferStatus string

func NewTransactionTransferStatus(s string) (TransactionTransferStatus, error) {
	switch s {
	case TransactionTransferStatusConfirmed:
		return TransactionTransferStatus(s), nil
	default:
		return "", ErrInvalidTransactionStatus
	}
}

func (s TransactionTransferStatus) Value() string {
	return string(s)
}

const (
	TransactionTransferStatusConfirmed = "confirmed"
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
	id int64,
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

	_fromAccountID, err := account.NewAccountID(fromAccountID)
	if err != nil {
		return nil, err
	}

	_toAccountID, err := account.NewAccountID(toAccountID)
	if err != nil {
		return nil, err
	}

	_amount, err := domainAmount.NewAmountPositive(amount)
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
