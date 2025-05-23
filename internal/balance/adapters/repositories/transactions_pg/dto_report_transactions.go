package transactions_pg

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

type reportTransactionDTO struct {
	ID              uuid.UUID      `db:"id"`
	UserID          int64          `db:"user_id"`
	OrderID         sql.NullInt64  `db:"order_id"`
	ProductID       sql.NullInt64  `db:"product_id"`
	ProductTitle    sql.NullString `db:"product_title"`
	DepositSource   sql.NullString `db:"deposit_source"`
	AccountID       sql.NullInt64  `db:"account_id"`
	ToAccountID     sql.NullInt64  `db:"to_account_id"`
	FromAccountID   sql.NullInt64  `db:"from_account_id"`
	TransactionType string         `db:"transaction_type"`
	Amount          int64          `db:"amount"`
	Status          string         `db:"status"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}

func fromReportTransactionDTO(dto *reportTransactionDTO) (report_transactions.Transaction, error) {
	switch dto.TransactionType {
	case "deposit":
		transactionDeposit, err := fromTransactionDepositDTO(&transactionDepositDTO{
			ID:            dto.ID,
			UserID:        dto.UserID,
			AccountID:     dto.AccountID.Int64,
			Amount:        dto.Amount,
			DepositSource: dto.DepositSource.String,
			Status:        dto.Status,
			CreatedAt:     dto.CreatedAt,
			UpdatedAt:     dto.UpdatedAt,
		})
		if err != nil {
			return nil, domainError.Strip(err)
		}
		return report_transactions.Transaction(transactionDeposit), nil
	case "spend":
		transactionSpend, err := fromTransactionSpendDTO(&transactionSpendDTO{
			ID:           dto.ID,
			UserID:       dto.UserID,
			OrderID:      dto.OrderID.Int64,
			ProductID:    dto.ProductID.Int64,
			ProductTitle: dto.ProductTitle.String,
			AccountID:    dto.AccountID.Int64,
			Amount:       dto.Amount,
			Status:       dto.Status,
			CreatedAt:    dto.CreatedAt,
			UpdatedAt:    dto.UpdatedAt,
		})
		if err != nil {
			return nil, domainError.Strip(err)
		}
		return report_transactions.Transaction(transactionSpend), nil
	case "transfer":
		transactionTransfer, err := fromTransactionTransferDTO(&transactionTransferDTO{
			ID:            dto.ID,
			ToAccountID:   dto.ToAccountID.Int64,
			FromAccountID: dto.FromAccountID.Int64,
			Amount:        dto.Amount,
			Status:        dto.Status,
			CreatedAt:     dto.CreatedAt,
			UpdatedAt:     dto.UpdatedAt,
		})
		if err != nil {
			return nil, domainError.Strip(err)
		}
		return report_transactions.Transaction(transactionTransfer), nil
	}

	return nil, errors.New("invalid transaction type")
}
