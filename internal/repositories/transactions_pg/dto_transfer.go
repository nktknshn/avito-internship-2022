package transactions_pg

import (
	"time"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type transactionTransferDTO struct {
	ID            int64     `db:"id"`
	FromAccountID int64     `db:"from_account_id"`
	ToAccountID   int64     `db:"to_account_id"`
	Amount        int64     `db:"amount"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func toTransactionTransferDTO(transaction *domainTransaction.TransactionTransfer) (*transactionTransferDTO, error) {
	return &transactionTransferDTO{
		ID:            transaction.ID.Value(),
		FromAccountID: transaction.FromAccountID.Value(),
		ToAccountID:   transaction.ToAccountID.Value(),
		Amount:        transaction.Amount.Value(),
		Status:        transaction.Status.Value(),
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}, nil
}

func fromTransactionTransferDTO(dto *transactionTransferDTO) (*domainTransaction.TransactionTransfer, error) {
	return domainTransaction.NewTransactionTransferFromValues(
		dto.ID,
		dto.FromAccountID,
		dto.ToAccountID,
		dto.Amount,
		dto.Status,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
