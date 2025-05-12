package transactions_pg

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

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

func reportTransactionsValidateCursorSorting(
	cursor *cursorUnion,
	sorting report_transactions.Sorting,
) error {
	// валидируем комбинацию сортировки и курсора

	if cursor.IsZero() || sorting.IsZero() {
		return nil
	}

	if cursor.IsAmount() && !sorting.IsAmount() {
		return errors.New("cursor is amount, sorting is not amount")
	}

	if cursor.IsUpdatedAt() && !sorting.IsUpdatedAt() {
		return errors.New("cursor is updated_at, sorting is not updated_at")
	}

	return nil
}

func reportTransactionsSetCursor(
	qb *select_query_builder.SelectQueryBuilder,
	sorting report_transactions.Sorting,
	sortingDirection report_transactions.SortingDirection,
	queryCursor report_transactions.Cursor,
	queryArgs map[string]any,
) error {
	if queryCursor == report_transactions.CursorEmpty {
		return nil
	}

	cursor, err := unmarshalCursor(queryCursor)

	if err != nil {
		return errors.Wrap(report_transactions.ErrCursorInvalid, err.Error())
	}

	err = reportTransactionsValidateCursorSorting(cursor, sorting)

	if err != nil {
		return errors.Wrap(report_transactions.ErrSortingCursorInvalid, err.Error())
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

func reportTransactionsSetSorting(
	qb *select_query_builder.SelectQueryBuilder,
	sorting report_transactions.Sorting,
	sortingDirection report_transactions.SortingDirection,
) {

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

var (
	defaultSorting          = report_transactions.SortingUpdatedAt
	defaultSortingDirection = report_transactions.SortingDirectionDesc
)

func (r *TransactionsRepository) GetTransactionsByUserID(
	ctx context.Context,
	userID domain.UserID,
	query report_transactions.GetTransactionsQuery,
) (report_transactions.ReportTransactionsPage, error) {

	qb := select_query_builder.New()

	qb.Ctes = []string{cteQuery}
	qb.Query = sqlQuery

	queryArgs := map[string]any{
		"user_id": userID,
	}

	queryLimit := query.Limit.Value()

	if queryLimit > 0 {
		qb.Limit = strconv.FormatUint(queryLimit+1, 10)
		queryArgs["limit"] = queryLimit + 1
	}

	sorting := defaultSorting
	sortingDirection := defaultSortingDirection

	if !query.SortingDirection.IsZero() {
		sortingDirection = query.SortingDirection
	}

	if !query.Sorting.IsZero() {
		sorting = query.Sorting
	}

	err := reportTransactionsSetCursor(qb, sorting, sortingDirection, query.Cursor, queryArgs)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, err
	}

	reportTransactionsSetSorting(qb, sorting, sortingDirection)

	transactions := []reportTransactionDTO{}

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return report_transactions.ReportTransactionsPage{}, errors.New("TransactionsRepository.GetTransactionsByUserID: tr is nil")
	}

	sql, args, err := tr.BindNamed(qb.Build(), queryArgs)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, errors.Wrap(err, "TransactionsRepository.GetTransactionsByUserID.BindNamed")
	}

	err = tr.SelectContext(ctx, &transactions, sql, args...)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, errors.Wrap(
			err,
			"TransactionsRepository.GetTransactionsByUserID.SelectContext",
		)
	}

	if len(transactions) == 0 {
		return report_transactions.ReportTransactionsPage{
			Transactions: []report_transactions.Transaction{},
			HasMore:      false,
			Cursor:       report_transactions.CursorEmpty,
		}, nil
	}

	var result report_transactions.ReportTransactionsPage

	if queryLimit > 0 {
		resultLength := min(len(transactions), int(queryLimit))
		result.Transactions = make([]report_transactions.Transaction, resultLength)
		result.HasMore = len(transactions) > int(queryLimit)
		transactions = transactions[:resultLength]
	} else {
		result.Transactions = make([]report_transactions.Transaction, len(transactions))
		result.HasMore = false
	}

	for i, transaction := range transactions {
		model, err := fromReportTransactionDTO(&transaction)
		if err != nil {
			return report_transactions.ReportTransactionsPage{}, errors.Wrap(
				err,
				"TransactionsRepository.GetTransactionsByUserID.fromReportTransactionDTO",
			)
		}

		result.Transactions[i] = model

		if i < len(transactions)-1 {
			continue
		}

		// устанавливаем курсор
		err = reportTransactionsSetCursorResult(&result, sorting, transactions[i])

		if err != nil {
			return report_transactions.ReportTransactionsPage{}, err
		}
	}

	return result, nil
}

func reportTransactionsSetCursorResult(
	result *report_transactions.ReportTransactionsPage,
	sorting report_transactions.Sorting,
	transaction reportTransactionDTO,
) error {
	if sorting.IsAmount() {
		nextCursor, err := marshalCursor(&cursorUnion{
			Amount: &cursorAmount{
				Amount: transaction.Amount,
				ID:     transaction.ID,
			},
		})
		if err != nil {
			return errors.Wrap(
				err,
				"TransactionsRepository.GetTransactionsByUserID.marshalCursor",
			)
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
			return errors.Wrap(
				err,
				"TransactionsRepository.GetTransactionsByUserID.marshalCursor",
			)
		}
		result.Cursor = nextCursor
	}

	return nil
}
