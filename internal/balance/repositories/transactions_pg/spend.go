package transactions_pg

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

func (r *TransactionsRepository) GetTransactionSpendByOrderID(ctx context.Context, userID domain.UserID, orderID domainAccount.OrderID) ([]*domainTransaction.TransactionSpend, error) {
	sq := `
		SELECT id, account_id, user_id, order_id, product_id, status, amount, created_at, updated_at 
		FROM transactions_spend 
		WHERE user_id = ? AND order_id = ? FOR UPDATE;
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	var transactions []*transactionSpendDTO

	err := tr.SelectContext(ctx, &transactions, r.db.Rebind(sq), userID, orderID)

	if err != nil {
		return nil, err
	}

	result := make([]*domainTransaction.TransactionSpend, len(transactions))

	for i, transaction := range transactions {
		result[i], err = fromTransactionSpendDTO(transaction)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r *TransactionsRepository) SaveTransactionSpend(ctx context.Context, transaction *domainTransaction.TransactionSpend) (*domainTransaction.TransactionSpend, error) {
	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	transactionDTO, err := toTransactionSpendDTO(transaction)
	if err != nil {
		return nil, err
	}

	var newDTO *transactionSpendDTO

	if transactionDTO.ID == 0 {
		newDTO, err = r.createTransactionSpend(ctx, tr, transactionDTO)
	} else {
		newDTO, err = r.updateTransactionSpend(ctx, tr, transactionDTO)
	}

	if err != nil {
		return nil, err
	}

	return fromTransactionSpendDTO(newDTO)
}

func (r *TransactionsRepository) createTransactionSpend(ctx context.Context, tr trmsqlx.Tr, transactionDTO *transactionSpendDTO) (*transactionSpendDTO, error) {
	sq := `
		INSERT INTO transactions_spend 
			(account_id, user_id, order_id, product_id, status, amount, created_at, updated_at) 
		VALUES 
			(:account_id, :user_id, :order_id, :product_id, :status, :amount, :created_at, :updated_at)
		RETURNING *;
	`

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, err
	}

	var newDTO transactionSpendDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)

	if err != nil {
		return nil, err
	}

	return &newDTO, nil
}

func (r *TransactionsRepository) updateTransactionSpend(ctx context.Context, tr trmsqlx.Tr, transactionDTO *transactionSpendDTO) (*transactionSpendDTO, error) {
	sq := `
		UPDATE transactions_spend 
		SET 
			status = :status, 
			updated_at = :updated_at 
		WHERE id = :id 
		RETURNING *;
	`

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, err
	}

	var newDTO transactionSpendDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, err
	}

	return &newDTO, nil
}
