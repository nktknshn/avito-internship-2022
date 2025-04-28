package report_transactions

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrReportTransactionsCursorInvalid = domainError.New("report transactions cursor invalid")
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
