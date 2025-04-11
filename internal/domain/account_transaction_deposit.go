package domain

import (
	"errors"
	"time"
)

type DepositSource string

func (s DepositSource) Value() string {
	return string(s)
}

func NewDepositSource(s string) (DepositSource, error) {
	return DepositSource(s), nil
}

type AccountTransactionDepositStatus string

func (s AccountTransactionDepositStatus) Value() string {
	return string(s)
}

const (
	AccountTransactionDepositStatusUnknown = ""
	AccountTransactionDepositStatusDone    = "done"
)

var (
	ErrInvalidAccountTransactionDepositStatus = errors.New("invalid deposit transaction status")
)

func NewAccountTransactionDepositStatus(s string) (AccountTransactionDepositStatus, error) {
	switch s {
	case AccountTransactionDepositStatusDone:
		return AccountTransactionDepositStatusDone, nil
	}
	return "", ErrInvalidAccountTransactionDepositStatus
}

type AccountTransactionDeposit struct {
	ID            AccountTransactionDepositID
	AccountID     AccountID
	UserID        UserID
	DepositSource DepositSource
	Status        AccountTransactionDepositStatus
	Amount        AmountPositive
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewAccountTransactionDeposit(accountID AccountID, userID UserID, source DepositSource, amount AmountPositive, now time.Time) (*AccountTransactionDeposit, error) {
	return &AccountTransactionDeposit{
		UserID:        userID,
		DepositSource: source,
		Amount:        amount,
		Status:        AccountTransactionDepositStatusDone,
		CreatedAt:     now,
		UpdatedAt:     now,
		AccountID:     accountID,
	}, nil
}

func NewAccountTransactionDepositFromValues(
	id int64,
	accountID int64,
	userID int64,
	depositSource string,
	status string,
	amount int64,
	createdAt time.Time,
	updatedAt time.Time,
) (*AccountTransactionDeposit, error) {
	_id, err := NewAccountTransactionDepositID(id)
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

	_depositSource, err := NewDepositSource(depositSource)
	if err != nil {
		return nil, err
	}

	_status, err := NewAccountTransactionDepositStatus(status)
	if err != nil {
		return nil, err
	}

	_amount, err := NewAmountPositive(amount)
	if err != nil {
		return nil, err
	}

	return &AccountTransactionDeposit{
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
