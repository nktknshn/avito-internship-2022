package sqlx_pg

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func GetPgError(err error) (*pgconn.PgError, bool) {

	if err == nil {
		return nil, false
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr, true
	}
	return nil, false
}

func IsDuplicateKeyError(err error) bool {

	if pgErr, ok := GetPgError(err); ok {
		return pgErr.Code == "23505"
	}

	return false
}
