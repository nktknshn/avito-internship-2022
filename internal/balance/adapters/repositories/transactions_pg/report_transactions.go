package transactions_pg

import (
	"context"
	"fmt"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	"github.com/nktknshn/avito-internship-2022/internal/common/select_query_builder"
)

var cteQuery = `combined_transactions AS (
	SELECT
		'deposit' as transaction_type,
		td.id,
		td.user_id,
		td.account_id,
		td.deposit_source,
		CAST(NULL as bigint) as order_id,
		CAST(NULL as bigint) as product_id,
		CAST(NULL as text) as product_title,
		CAST(NULL as bigint) as to_account_id,
		CAST(NULL as bigint) as from_account_id,
		td.amount,
		td.status,
		td.created_at,
		td.updated_at
	FROM transactions_deposit td
	WHERE td.user_id = :user_id

	UNION ALL

	SELECT
		'spend' as transaction_type,
		ts.id,
		ts.user_id,
		ts.account_id,
		CAST(NULL as varchar) as deposit_source,
		ts.order_id,
		ts.product_id,
		ts.product_title,
		CAST(NULL as bigint) as to_account_id,
		CAST(NULL as bigint) as from_account_id,
		ts.amount,
		ts.status,
		ts.created_at,
		ts.updated_at
	FROM transactions_spend ts
	WHERE ts.user_id = :user_id

	UNION ALL

	SELECT
		'transfer' as transaction_type,
		tt.id,
		a1.user_id,
		tt.from_account_id as account_id,
		CAST(NULL as varchar) as deposit_source,
		CAST(NULL as bigint) as order_id,
		CAST(NULL as bigint) as product_id,
		CAST(NULL as text) as product_title,
		tt.to_account_id,
		tt.from_account_id,
		tt.amount,
		tt.status,
		tt.created_at,
		tt.updated_at
	FROM transactions_transfer tt
	JOIN accounts a1 ON tt.from_account_id = a1.id OR tt.to_account_id = a1.id
	WHERE a1.user_id = :user_id
)
`

var sqlQuery = `SELECT * FROM combined_transactions t`

func (r *TransactionsRepository) setReportTransactionsCursor(
	qb *select_query_builder.SelectQueryBuilder,
	sorting report_transactions.Sorting,
	sortingDirection report_transactions.SortingDirection,
	queryCursor report_transactions.Cursor,
	queryArgs map[string]any,
) error {
	if queryCursor == nil {
		return nil
	}

	cursor, err := unmarshalCursor(queryCursor)

	if err != nil {
		return err
	}

	// валидируем комбинацию сортировки и курсора
	if !cursor.IsZero() && !sorting.IsZero() {
		if cursor.IsAmount() && !sorting.IsAmount() {
			return report_transactions.ErrSortingCursorInvalid
		} else if cursor.IsUpdatedAt() && !sorting.IsUpdatedAt() {
			return report_transactions.ErrSortingCursorInvalid
		}
	}

	if cursor.IsAmount() {
		if sortingDirection.IsAsc() {
			qb.And("(t.amount = :cursor_amount AND t.id > :cursor_id) OR t.amount > :cursor_amount")
		} else {
			qb.And("(t.amount = :cursor_amount AND t.id < :cursor_id) OR t.amount < :cursor_amount")
		}
		queryArgs["cursor_amount"] = cursor.Amount.Amount
		queryArgs["cursor_id"] = cursor.Amount.ID
	}

	if cursor.IsUpdatedAt() {
		if sortingDirection.IsAsc() {
			qb.And("(t.updated_at = :cursor_updated_at AND t.id > :cursor_id) OR t.updated_at > :cursor_updated_at")
		} else {
			qb.And("(t.updated_at = :cursor_updated_at AND t.id < :cursor_id) OR t.updated_at < :cursor_updated_at")
		}
		queryArgs["cursor_updated_at"] = cursor.UpdatedAt.UpdatedAt
		queryArgs["cursor_id"] = cursor.UpdatedAt.ID
	}

	return nil
}

func (r *TransactionsRepository) setReportTransactionsSorting(qb *select_query_builder.SelectQueryBuilder, sorting report_transactions.Sorting, sortingDirection report_transactions.SortingDirection) {

	if sortingDirection.IsAsc() {
		qb.Order = "t.updated_at ASC, t.id ASC"
	} else {
		qb.Order = "t.updated_at DESC, t.id DESC"
	}

	if sorting.IsAmount() {
		if sortingDirection.IsAsc() {
			qb.Order = "t.amount ASC, t.id ASC"
		} else {
			qb.Order = "t.amount DESC, t.id DESC"
		}
		return
	}

	if sorting.IsUpdatedAt() {
		if sortingDirection.IsAsc() {
			qb.Order = "t.updated_at ASC, t.id ASC"
		} else {
			qb.Order = "t.updated_at DESC, t.id DESC"
		}
		return
	}
}

func (r *TransactionsRepository) GetTransactionsByUserID(ctx context.Context, userID domain.UserID, query report_transactions.GetTransactionsQuery) (report_transactions.ReportTransactionsPage, error) {

	qb := select_query_builder.New()

	qb.Ctes = []string{cteQuery}
	qb.Query = sqlQuery

	queryArgs := map[string]any{
		"user_id": userID,
		"limit":   query.Limit.Value(),
	}

	if query.Limit.Value() > 0 {
		qb.Limit = fmt.Sprintf("%d", query.Limit)
	}

	sorting := report_transactions.SortingUpdatedAt
	sortingDirection := report_transactions.SortingDirectionDesc

	if !query.SortingDirection.IsZero() {
		sortingDirection = query.SortingDirection
	}

	if !query.Sorting.IsZero() {
		sorting = query.Sorting
	}

	r.setReportTransactionsCursor(qb, sorting, sortingDirection, query.Cursor, queryArgs)
	r.setReportTransactionsSorting(qb, sorting, sortingDirection)

	transactions := []reportTransactionDTO{}

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := tr.BindNamed(qb.Build(), queryArgs)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, err
	}

	err = tr.SelectContext(ctx, &transactions, sql, args...)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, err
	}

	if len(transactions) == 0 {
		return report_transactions.ReportTransactionsPage{
			Transactions: []report_transactions.Transaction{},
			HasMore:      false,
			Cursor:       nil,
		}, nil
	}

	result := report_transactions.ReportTransactionsPage{
		Transactions: make([]report_transactions.Transaction, len(transactions)),
	}

	for i, transaction := range transactions {
		model, err := fromReportTransactionDTO(&transaction)
		if err != nil {
			return report_transactions.ReportTransactionsPage{}, err
		}
		result.Transactions[i] = model

		if sorting.IsAmount() {
			nextCursor, err := marshalCursor(&cursorUnion{
				Amount: &cursorAmount{
					Amount: transaction.Amount,
					ID:     transaction.ID,
				},
			})
			if err != nil {
				return report_transactions.ReportTransactionsPage{}, err
			}
			result.Cursor = nextCursor
		}

		if sorting.IsUpdatedAt() {
			nextCursor, err := marshalCursor(&cursorUnion{
				UpdatedAt: &cursorUpdatedAt{
					UpdatedAt: transaction.UpdatedAt,
					ID:        transaction.ID,
				},
			})
			if err != nil {
				return report_transactions.ReportTransactionsPage{}, err
			}
			result.Cursor = nextCursor
		}
	}

	result.HasMore = len(transactions) == int(query.Limit)

	return result, nil
}
