package transaction

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrInvalidTransactionID     = domainError.New("invalid account transaction id")
	ErrInvalidTransactionStatus = domainError.New("invalid account transaction status")
)

type DepositSource string

func (s DepositSource) Value() string {
	return string(s)
}

func NewDepositSource(s string) (DepositSource, error) {
	return DepositSource(s), nil
}

type TransactionDepositStatus string

func (s TransactionDepositStatus) Value() string {
	return string(s)
}

const (
	TransactionDepositStatusConfirmed = "confirmed"
)

var (
	ErrInvalidTransactionDepositStatus = domainError.New("invalid deposit transaction status")
)

func NewTransactionDepositStatus(s string) (TransactionDepositStatus, error) {
	switch s {
	case TransactionDepositStatusConfirmed:
		return TransactionDepositStatusConfirmed, nil
	}
	return "", ErrInvalidTransactionDepositStatus
}

type TransactionDepositID int64

func (id TransactionDepositID) Value() int64 {
	return int64(id)
}

func NewTransactionDepositID(id int64) (TransactionDepositID, error) {
	if id < 0 {
		return 0, ErrInvalidTransactionID
	}
	return TransactionDepositID(id), nil
}

type TransactionDeposit struct {
	ID            TransactionDepositID
	AccountID     domainAccount.AccountID
	UserID        domain.UserID
	DepositSource DepositSource
	Status        TransactionDepositStatus
	Amount        domainAmount.AmountPositive
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewTransactionDeposit(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	source DepositSource,
	amount domainAmount.AmountPositive,
	now time.Time,
) (*TransactionDeposit, error) {
	return &TransactionDeposit{
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
	id int64,
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
