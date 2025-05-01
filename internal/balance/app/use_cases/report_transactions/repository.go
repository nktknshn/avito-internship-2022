package report_transactions

import (
	"context"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

var (
	ErrCursorInvalid        = useCaseError.New("report transactions cursor is invalid")
	ErrSortingCursorInvalid = useCaseError.New("report transactions sorting and cursor combination is invalid")
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
