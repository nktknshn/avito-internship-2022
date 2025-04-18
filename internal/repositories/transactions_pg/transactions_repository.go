package transactions_pg

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	domain "github.com/nktknshn/avito-internship-2022/internal/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type TransactionsRepository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewTransactionsRepository(db *sqlx.DB, c *trmsqlx.CtxGetter) *TransactionsRepository {
	return &TransactionsRepository{db: db, getter: c}
}

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
		return nil, err
	}

	sq, args, err := tr.BindNamed(sq, transactionDTO)
	if err != nil {
		return nil, err
	}

	var newDTO transactionDepositDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, err
	}

	return fromTransactionDepositDTO(&newDTO)
}

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

var _ domainTransaction.TransactionRepository = &TransactionsRepository{}
