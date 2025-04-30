package report_transactions

import (
	"errors"
)

type ReportTransactionsPage struct {
	Transactions []Transaction
	Cursor       Cursor
	HasMore      bool
}

// Limit
type Limit uint64

func (l Limit) Value() uint64 {
	return uint64(l)
}

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

var CursorEmpty = Cursor(nil)

func NewCursor(cursor any) (Cursor, error) {
	return cursor, nil
}

// domainTransaction.TransactionDeposit
// domainTransaction.TransactionSpend
// domainTransaction.TransactionTransfer
type Transaction any

type SortingDirection string

func (s SortingDirection) IsZero() bool {
	return s == ""
}

func (s SortingDirection) IsAsc() bool {
	return s == SortingDirectionAsc
}

func (s SortingDirection) IsDesc() bool {
	return s == SortingDirectionDesc
}

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
	SortingUpdatedAt Sorting = "updated_at"
	SortingAmount    Sorting = "amount"
)

func (s Sorting) IsZero() bool {
	return s == ""
}

func (s Sorting) IsAmount() bool {
	return s == SortingAmount
}

func (s Sorting) IsUpdatedAt() bool {
	return s == SortingUpdatedAt
}

func NewSorting(sorting string) (Sorting, error) {
	switch Sorting(sorting) {
	case SortingUpdatedAt, SortingAmount:
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
