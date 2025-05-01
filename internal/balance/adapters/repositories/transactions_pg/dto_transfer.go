package transactions_pg

import (
	"time"

	"github.com/google/uuid"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type transactionTransferDTO struct {
	ID            uuid.UUID `db:"id"`
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
	t, err := domainTransaction.NewTransactionTransferFromValues(
		dto.ID,
		dto.FromAccountID,
		dto.ToAccountID,
		dto.Amount,
		dto.Status,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
	if err != nil {
		return nil, domainError.Strip(err)
	}
	return t, nil
}
