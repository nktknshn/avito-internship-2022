package report_transactions

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type ReportTransactionsUseCase struct {
	repo ReportTransactionsRepository
}

func New(repo ReportTransactionsRepository) *ReportTransactionsUseCase {

	if repo == nil {
		panic("repo is nil")
	}

	return &ReportTransactionsUseCase{repo: repo}
}

func (u *ReportTransactionsUseCase) Handle(ctx context.Context, in In) (Out, error) {

	out := Out{
		Transactions: make([]outTransaction, 0),
	}

	page, err := u.repo.GetTransactionsByUserID(ctx, in.userID, GetTransactionsQuery{
		Limit:            in.limit,
		Cursor:           in.cursor,
		Sorting:          in.sorting,
		SortingDirection: in.sortingDirection,
	})

	if err != nil {
		return Out{}, err
	}

	out.Cursor = page.Cursor
	out.HasMore = page.HasMore

	for _, transaction := range page.Transactions {
		switch transaction := transaction.(type) {
		case *domainTransaction.TransactionSpend:
			out.Transactions = append(out.Transactions, &OutTransactionSpend{
				ID:           transaction.ID.Value(),
				AccountID:    transaction.AccountID.Value(),
				OrderID:      transaction.OrderID.Value(),
				ProductID:    transaction.ProductID.Value(),
				ProductTitle: transaction.ProductTitle.Value(),
				Amount:       transaction.Amount.Value(),
				Status:       transaction.Status.Value(),
				CreatedAt:    transaction.CreatedAt,
				UpdatedAt:    transaction.UpdatedAt,
			})
		case *domainTransaction.TransactionDeposit:
			out.Transactions = append(out.Transactions, &OutTransactionDeposit{
				ID:        transaction.ID.Value(),
				Source:    transaction.DepositSource.Value(),
				Amount:    transaction.Amount.Value(),
				Status:    transaction.Status.Value(),
				CreatedAt: transaction.CreatedAt,
				UpdatedAt: transaction.UpdatedAt,
			})
		case *domainTransaction.TransactionTransfer:
			out.Transactions = append(out.Transactions, &OutTransactionTransfer{
				ID:        transaction.ID.Value(),
				From:      transaction.FromAccountID.Value(),
				To:        transaction.ToAccountID.Value(),
				Amount:    transaction.Amount.Value(),
				Status:    transaction.Status.Value(),
				CreatedAt: transaction.CreatedAt,
				UpdatedAt: transaction.UpdatedAt,
			})
		}
	}

	return out, nil
}

func (u *ReportTransactionsUseCase) GetName() string {
	return use_cases.NameReportTransactions
}
