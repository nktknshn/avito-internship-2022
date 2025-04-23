package transactions_pg

import (
	"context"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/pkg/errors"
)

func (r *TransactionsRepository) SaveTransactionDeposit(ctx context.Context, transaction *domainTransaction.TransactionDeposit) (*domainTransaction.TransactionDeposit, error) {
	sq := `
		INSERT INTO transactions_deposit 
			(account_id, user_id, deposit_source, status, amount, created_at, updated_at) 
		VALUES 
			(:account_id, :user_id, :deposit_source, :status, :amount, :created_at, :updated_at)
		RETURNING *;
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	transactionDTO, err := toTransactionDepositDTO(transaction)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.toTransactionDepositDTO")
	}

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.BindNamed")
	}

	var newDTO transactionDepositDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.GetContext")
	}

	return fromTransactionDepositDTO(&newDTO)
}
