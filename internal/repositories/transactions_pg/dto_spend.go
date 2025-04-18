package transactions_pg

import (
	"time"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type transactionSpendDTO struct {
	ID        int64     `db:"id"`
	AccountID int64     `db:"account_id"`
	UserID    int64     `db:"user_id"`
	OrderID   int64     `db:"order_id"`
	ProductID int64     `db:"product_id"`
	Status    string    `db:"status"`
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func toTransactionSpendDTO(transaction *domainTransaction.TransactionSpend) (*transactionSpendDTO, error) {
	return &transactionSpendDTO{
		ID:        transaction.ID.Value(),
		AccountID: transaction.AccountID.Value(),
		UserID:    transaction.UserID.Value(),
		OrderID:   transaction.OrderID.Value(),
		ProductID: transaction.ProductID.Value(),
		Status:    transaction.Status.Value(),
		Amount:    transaction.Amount.Value(),
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}, nil
}

func fromTransactionSpendDTO(dto *transactionSpendDTO) (*domainTransaction.TransactionSpend, error) {
	return domainTransaction.NewTransactionSpendFromValues(
		dto.ID,
		dto.AccountID,
		dto.UserID,
		dto.OrderID,
		dto.ProductID,
		dto.Amount,
		dto.Status,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
