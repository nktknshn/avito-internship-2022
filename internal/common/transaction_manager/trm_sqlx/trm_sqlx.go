package trm_sqlx

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type transactionKeyType string

const transactionKey transactionKeyType = "sqlx_transaction"

func GetFromContext(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(transactionKey).(*sqlx.Tx)
	if ok {
		return tx, true
	}
	return nil, false
}

func PutToContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, transactionKey, tx)
}

func CreateIfNeeded(ctx context.Context, db *sqlx.DB) (context.Context, bool) {
	_, ok := GetFromContext(ctx)
	if ok {
		return ctx, true
	}
	tx, err := db.Beginx()
	if err != nil {
		return nil, false
	}
	ctx = PutToContext(ctx, tx)
	return ctx, true
}
