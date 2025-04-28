package report_transactions

import (
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type ReportTransactionsPage struct {
	AccountID    domainAccount.AccountID
	UserID       domain.UserID
	Transactions []Transaction
	Cursor       Cursor
	HasMore      bool
}

// Limit
type Limit int

func NewLimit(limit int) (Limit, error) {
	if limit <= 0 {
		return 0, errors.New("limit must be greater than 0")
	}
	return Limit(limit), nil
}

const (
	ReportTransactionsLimitDefault Limit = 100
)

// Реализуется в репозитории
type Cursor any

func NewCursor(cursor any) (Cursor, error) {
	return cursor, nil
}

// domainTransaction.TransactionDeposit
// domainTransaction.TransactionSpend
// domainTransaction.TransactionTransfer
type Transaction any

type SortingDirection string

func NewSortingDirection(direction string) (SortingDirection, error) {
	switch SortingDirection(direction) {
	case SortingDirectionAsc, SortingDirectionDesc:
		return SortingDirection(direction), nil
	default:
		return "", errors.New("invalid sorting direction")
	}
}

const (
	SortingDirectionAsc  SortingDirection = "asc"
	SortingDirectionDesc SortingDirection = "desc"
)

type Sorting string

const (
	SortingCreatedAt Sorting = "created_at"
	SortingUpdatedAt Sorting = "updated_at"
	SortingAmount    Sorting = "amount"
	SortingStatus    Sorting = "status"
)

func NewSorting(sorting string) (Sorting, error) {
	switch Sorting(sorting) {
	case SortingCreatedAt, SortingUpdatedAt, SortingAmount, SortingStatus:
		return Sorting(sorting), nil
	default:
		return "", errors.New("invalid sorting")
	}
}

type TransactionType string

const (
	TransactionTypeSpend    TransactionType = "spend"
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeTransfer TransactionType = "transfer"
)

func NewTransactionType(transactionType string) (TransactionType, error) {
	switch TransactionType(transactionType) {
	case TransactionTypeSpend, TransactionTypeDeposit, TransactionTypeTransfer:
		return TransactionType(transactionType), nil
	default:
		return "", errors.New("invalid transaction type")
	}
}
