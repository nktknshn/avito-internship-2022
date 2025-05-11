package transactions_pg

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

func (r *TransactionsRepository) GetTransactionSpendByOrderID(
	ctx context.Context,
	userID domain.UserID,
	orderID domain.OrderID,
) ([]*domainTransaction.TransactionSpend, error) {
	sq := `
		SELECT 
			id, 
			account_id, 
			user_id, 
			order_id, 
			product_id, 
			product_title, 
			status, 
			amount, 
			created_at, 
			updated_at 
		FROM transactions_spend 
		WHERE user_id = ? AND order_id = ? FOR UPDATE;
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("TransactionsRepository.GetTransactionSpendByOrderID: tr is nil")
	}

	var transactions []*transactionSpendDTO

	err := tr.SelectContext(ctx, &transactions, tr.Rebind(sq), userID, orderID)

	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.GetTransactionSpendByOrderID.SelectContext")
	}

	result := make([]*domainTransaction.TransactionSpend, len(transactions))

	for i, transaction := range transactions {
		result[i], err = fromTransactionSpendDTO(transaction)
		if err != nil {
			return nil, errors.Wrap(err, "TransactionsRepository.GetTransactionSpendByOrderID.fromTransactionSpendDTO")
		}
	}
	return result, nil
}

func (r *TransactionsRepository) SaveTransactionSpend(
	ctx context.Context,
	transaction *domainTransaction.TransactionSpend,
) (*domainTransaction.TransactionSpend, error) {
	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("TransactionsRepository.SaveTransactionSpend: tr is nil")
	}

	transactionDTO := toTransactionSpendDTO(transaction)

	var newDTO *transactionSpendDTO
	var err error

	if transactionDTO.ID == uuid.Nil {
		newDTO, err = r.createTransactionSpend(ctx, tr, transactionDTO)
	} else {
		newDTO, err = r.updateTransactionSpend(ctx, tr, transactionDTO)
	}

	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionSpend")
	}

	res, err := fromTransactionSpendDTO(newDTO)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.SaveTransactionSpend.fromTransactionSpendDTO")
	}

	return res, nil
}

func (r *TransactionsRepository) createTransactionSpend(
	ctx context.Context,
	tr trmsqlx.Tr,
	transactionDTO *transactionSpendDTO,
) (*transactionSpendDTO, error) {
	sq := `
		INSERT INTO transactions_spend 
			(account_id, user_id, order_id, product_id, product_title, status, amount, created_at, updated_at) 
		VALUES 
			(:account_id, :user_id, :order_id, :product_id, :product_title, :status, :amount, :created_at, :updated_at)
		RETURNING *;
	`

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.createTransactionSpend.BindNamed")
	}

	var newDTO transactionSpendDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)

	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.createTransactionSpend.GetContext")
	}

	return &newDTO, nil
}

func (r *TransactionsRepository) updateTransactionSpend(
	ctx context.Context,
	tr trmsqlx.Tr,
	transactionDTO *transactionSpendDTO,
) (*transactionSpendDTO, error) {
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
		return nil, errors.Wrap(err, "TransactionsRepository.updateTransactionSpend.BindNamed")
	}

	var newDTO transactionSpendDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, errors.Wrap(err, "TransactionsRepository.updateTransactionSpend.GetContext")
	}

	return &newDTO, nil
}
