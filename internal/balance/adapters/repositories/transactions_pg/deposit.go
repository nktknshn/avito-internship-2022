package transactions_pg

import (
	"context"

	"github.com/pkg/errors"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

func (r *TransactionsRepository) SaveTransactionDeposit(
	ctx context.Context,
	transaction *domainTransaction.TransactionDeposit,
) (*domainTransaction.TransactionDeposit, error) {
	sq := `
		INSERT INTO transactions_deposit 
			(account_id, user_id, deposit_source, status, amount, created_at, updated_at) 
		VALUES 
			(:account_id, :user_id, :deposit_source, :status, :amount, :created_at, :updated_at)
		RETURNING *;
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("TransactionsRepository.SaveTransactionDeposit: tr is nil")
	}

	transactionDTO := toTransactionDepositDTO(transaction)

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.BindNamed")
	}

	var newDTO transactionDepositDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.GetContext")
	}

	t, err := fromTransactionDepositDTO(&newDTO)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionDeposit.fromTransactionDepositDTO")
	}
	return t, nil
}
