package transactions_pg

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type TransactionsRepository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func NewTransactionsRepository(db *sqlx.DB, c *trmsqlx.CtxGetter) *TransactionsRepository {
	return &TransactionsRepository{db: db, getter: c}
}

var _ domainTransaction.TransactionRepository = &TransactionsRepository{}
