package transactions_pg

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

func (r *TransactionsRepository) GetTransactionsByUserID(ctx context.Context, userID domain.UserID, query report_transactions.GetTransactionsQuery) (report_transactions.ReportTransactionsPage, error) {
	// q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	result := report_transactions.ReportTransactionsPage{}
	return result, nil

}
