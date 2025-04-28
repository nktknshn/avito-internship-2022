package transactions_pg

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

var sqlQuery = `
	WITH combined_transactions AS (
            -- Deposit transactions
            SELECT 
                'deposit' as transaction_type,
                td.id,
                user_id,
                account_id,
				CAST(NULL as bigint) as order_id,
				CAST(NULL as bigint) as product_id,
                CAST(NULL as bigint) as to_account_id,
				CAST(NULL as bigint) as from_account_id,
                amount,
                status,
                created_at,
				updated_at
            FROM transactions_deposit td
            WHERE user_id = :user_id

            UNION ALL

            -- Spend transactions
            SELECT 
                'spend' as transaction_type,
                ts.id,
                user_id,
                account_id,
				order_id,
				product_id,
                CAST(NULL as bigint) as to_account_id,
                CAST(NULL as bigint) as from_account_id,
                amount,
                status,
                created_at,
				updated_at
            FROM transactions_spend ts
            WHERE user_id = :user_id

            UNION ALL

            -- Transfer transactions (as sender)
            SELECT 
                'transfer' as transaction_type,
                tt.id,
                a1.user_id,
				CAST(NULL as bigint) as order_id,
				CAST(NULL as bigint) as product_id,
				CAST(NULL as bigint) as account_id,
                from_account_id,
                to_account_id,
                amount,
                status,
                created_at,
				updated_at
            FROM transactions_transfer tt
            JOIN accounts a1 ON tt.from_account_id = a1.id
            WHERE a1.user_id = :user_id

            UNION ALL

            -- Transfer transactions (as receiver)
            SELECT 
                'transfer' as transaction_type,
                tt.id,
                a2.user_id,
				CAST(NULL as bigint) as account_id,
				CAST(NULL as bigint) as order_id,
				CAST(NULL as bigint) as product_id,
                to_account_id,
                from_account_id,
                amount,
                status,
                created_at,
				updated_at
            FROM transactions_transfer tt
            JOIN accounts a2 ON tt.to_account_id = a2.id
            WHERE a2.user_id = :user_id
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
