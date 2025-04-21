package transactions_pg

import (
	"context"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

func (r *TransactionsRepository) SaveTransactionTransfer(ctx context.Context, transaction *domainTransaction.TransactionTransfer) (*domainTransaction.TransactionTransfer, error) {
	sq := `
		INSERT INTO transactions_transfer 
			(account_id, user_id, order_id, product_id, status, amount, created_at, updated_at) 
		VALUES 
			(:account_id, :user_id, :order_id, :product_id, :status, :amount, :created_at, :updated_at)
		RETURNING *;
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	transactionDTO, err := toTransactionTransferDTO(transaction)
	if err != nil {
		return nil, err
	}

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, err
	}

	var newDTO transactionTransferDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, err
	}

	return fromTransactionTransferDTO(&newDTO)
}
