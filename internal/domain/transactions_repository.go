package domain

import (
	"context"
)

type TransactionRepository interface {
	SaveTransactionDeposit(ctx context.Context, transaction *TransactionDeposit) (*TransactionDeposit, error)
	SaveTransactionSpend(ctx context.Context, transaction *TransactionSpend) (*TransactionSpend, error)
	GetTransactionSpendByOrderID(ctx context.Context, userID UserID, orderID OrderID) ([]*TransactionSpend, error)
}
