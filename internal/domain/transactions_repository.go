package domain

import "context"

type AccountTransactionRepository interface {
	SaveTransactionDeposit(ctx context.Context, transaction *AccountTransactionDeposit) (*AccountTransactionDeposit, error)
	SaveTransactionSpend(ctx context.Context, transaction *AccountTransactionSpend) (*AccountTransactionSpend, error)
	GetTransactionSpendByOrderID(ctx context.Context, userID UserID, orderID OrderID) ([]*AccountTransactionSpend, error)
}
