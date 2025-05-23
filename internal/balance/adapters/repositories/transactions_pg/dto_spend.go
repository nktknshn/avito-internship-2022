package transactions_pg

import (
	"time"

	"github.com/google/uuid"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type transactionSpendDTO struct {
	ID           uuid.UUID `db:"id"`
	AccountID    int64     `db:"account_id"`
	UserID       int64     `db:"user_id"`
	OrderID      int64     `db:"order_id"`
	ProductID    int64     `db:"product_id"`
	ProductTitle string    `db:"product_title"`
	Status       string    `db:"status"`
	Amount       int64     `db:"amount"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func toTransactionSpendDTO(transaction *domainTransaction.TransactionSpend) *transactionSpendDTO {
	return &transactionSpendDTO{
		ID:           transaction.ID.Value(),
		AccountID:    transaction.AccountID.Value(),
		UserID:       transaction.UserID.Value(),
		OrderID:      transaction.OrderID.Value(),
		ProductID:    transaction.ProductID.Value(),
		ProductTitle: transaction.ProductTitle.Value(),
		Status:       transaction.Status.Value(),
		Amount:       transaction.Amount.Value(),
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
	}
}

func fromTransactionSpendDTO(dto *transactionSpendDTO) (*domainTransaction.TransactionSpend, error) {
	t, err := domainTransaction.NewTransactionSpendFromValues(
		dto.ID,
		dto.AccountID,
		dto.UserID,
		dto.OrderID,
		dto.ProductID,
		dto.ProductTitle,
		dto.Amount,
		dto.Status,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
	if err != nil {
		// не возвращаем доменную ошибку, так как это внутренняя ошибка адаптера
		return nil, domainError.Strip(err)
	}
	return t, nil
}
