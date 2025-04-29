package transactions_pg

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

var sqlQuery = `
	WITH combined_transactions AS (
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
	SELECT *
	FROM combined_transactions
`

func (r *TransactionsRepository) GetTransactionsByUserID(ctx context.Context, userID domain.UserID, query report_transactions.GetTransactionsQuery) (report_transactions.ReportTransactionsPage, error) {
	// q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	transactions := []reportTransactionDTO{}

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	sql, args, err := tr.BindNamed(sqlQuery, map[string]interface{}{
		"user_id": userID,
	})

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, err
	}

	err = tr.SelectContext(ctx, &transactions, sql, args...)

	if err != nil {
		return report_transactions.ReportTransactionsPage{}, err
	}

	models := make([]report_transactions.Transaction, len(transactions))

	for i, transaction := range transactions {
		model, err := fromReportTransactionDTO(&transaction)
		if err != nil {
			return report_transactions.ReportTransactionsPage{}, err
		}
		models[i] = model
	}

	return report_transactions.ReportTransactionsPage{
		Transactions: models,
	}, nil
}
