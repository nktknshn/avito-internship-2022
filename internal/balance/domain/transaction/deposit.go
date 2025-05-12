package transaction

import (
	"time"

	"github.com/google/uuid"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrInvalidTransactionID     = domainError.New("invalid account transaction id")
	ErrInvalidTransactionStatus = domainError.New("invalid account transaction status")
	ErrInvalidDepositSource     = domainError.New("invalid deposit source")
)

type TransactionDepositSource string

func (s TransactionDepositSource) Value() string {
	return string(s)
}

func NewDepositSource(s string) (TransactionDepositSource, error) {
	if s == "" {
		return "", ErrInvalidDepositSource
	}
	return TransactionDepositSource(s), nil
}

type TransactionDepositStatus string

func (s TransactionDepositStatus) Value() string {
	return string(s)
}

const (
	TransactionDepositStatusConfirmed TransactionDepositStatus = "confirmed"
)

var (
	ErrInvalidTransactionDepositStatus = domainError.New("invalid deposit transaction status")
)

func NewTransactionDepositStatus(s string) (TransactionDepositStatus, error) {
	if s == TransactionDepositStatusConfirmed.Value() {
		return TransactionDepositStatusConfirmed, nil
	}
	return "", ErrInvalidTransactionDepositStatus
}

type TransactionDepositID uuid.UUID

func NewTransactionDepositID(id uuid.UUID) (TransactionDepositID, error) {
	return TransactionDepositID(id), nil
}

func (id TransactionDepositID) Value() uuid.UUID {
	return uuid.UUID(id)
}

var TransactionDepositIDNil = TransactionDepositID(uuid.Nil)

type TransactionDeposit struct {
	ID            TransactionDepositID
	AccountID     domainAccount.AccountID
	UserID        domain.UserID
	DepositSource TransactionDepositSource
	Status        TransactionDepositStatus
	Amount        domainAmount.AmountPositive
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewTransactionDeposit(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	source TransactionDepositSource,
	amount domainAmount.AmountPositive,
	now time.Time,
) (*TransactionDeposit, error) {
	return &TransactionDeposit{
		ID:            TransactionDepositIDNil,
		UserID:        userID,
		DepositSource: source,
		Amount:        amount,
		Status:        TransactionDepositStatusConfirmed,
		CreatedAt:     now,
		UpdatedAt:     now,
		AccountID:     accountID,
	}, nil
}

func NewTransactionDepositFromValues(
	id uuid.UUID,
	accountID int64,
	userID int64,
	depositSource string,
	status string,
	amount int64,
	createdAt time.Time,
	updatedAt time.Time,
) (*TransactionDeposit, error) {

	_id, err := NewTransactionDepositID(id)
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

	_depositSource, err := NewDepositSource(depositSource)
	if err != nil {
		return nil, err
	}

	_status, err := NewTransactionDepositStatus(status)
	if err != nil {
		return nil, err
	}

	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return nil, err
	}

	return &TransactionDeposit{
		ID:            _id,
		AccountID:     _accountID,
		UserID:        _userID,
		DepositSource: _depositSource,
		Status:        _status,
		Amount:        _amount,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}
