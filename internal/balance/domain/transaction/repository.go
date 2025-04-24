package transaction

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

type TransactionRepository interface {
	// upserts
	SaveTransactionDeposit(ctx context.Context, transaction *TransactionDeposit) (*TransactionDeposit, error)
	SaveTransactionSpend(ctx context.Context, transaction *TransactionSpend) (*TransactionSpend, error)
	SaveTransactionTransfer(ctx context.Context, transaction *TransactionTransfer) (*TransactionTransfer, error)
	// queries
	GetTransactionSpendByOrderID(ctx context.Context, userID domain.UserID, orderID domain.OrderID) ([]*TransactionSpend, error)
}
