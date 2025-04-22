package sqlx_pg

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Migrate(ctx context.Context, conn *sql.DB, folder string) error {
	return goose.UpContext(ctx, conn, folder)
}
