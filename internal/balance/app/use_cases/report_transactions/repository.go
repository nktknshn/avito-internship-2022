package report_transactions

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrCursorInvalid        = domainError.New("report transactions cursor is invalid")
	ErrSortingCursorInvalid = domainError.New("report transactions sorting and cursor combination is invalid")
)

type GetTransactionsQuery struct {
	Limit            Limit
	Cursor           Cursor
	Sorting          Sorting
	SortingDirection SortingDirection
}

type ReportTransactionsRepository interface {
	GetTransactionsByUserID(ctx context.Context, userID domain.UserID, query GetTransactionsQuery) (ReportTransactionsPage, error)
}
