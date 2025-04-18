package transactions_pg

import (
	"time"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type transactionDepositDTO struct {
	ID            int64     `db:"id"`
	AccountID     int64     `db:"account_id"`
	UserID        int64     `db:"user_id"`
	DepositSource string    `db:"deposit_source"`
	Status        string    `db:"status"`
	Amount        int64     `db:"amount"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func toTransactionDepositDTO(transaction *domainTransaction.TransactionDeposit) (*transactionDepositDTO, error) {
	return &transactionDepositDTO{
		ID:            transaction.ID.Value(),
		AccountID:     transaction.AccountID.Value(),
		UserID:        transaction.UserID.Value(),
		DepositSource: transaction.DepositSource.Value(),
		Status:        transaction.Status.Value(),
		Amount:        transaction.Amount.Value(),
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}, nil
}

func fromTransactionDepositDTO(dto *transactionDepositDTO) (*domainTransaction.TransactionDeposit, error) {
	return domainTransaction.NewTransactionDepositFromValues(
		dto.ID,
		dto.AccountID,
		dto.UserID,
		dto.DepositSource,
		dto.Status,
		dto.Amount,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
