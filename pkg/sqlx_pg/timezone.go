package sqlx_pg

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	pfgxStdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var timeZoneOption = pfgxStdlib.OptionAfterConnect(func(_ context.Context, c *pgx.Conn) error {
	c.TypeMap().RegisterType(&pgtype.Type{
		Name:  "timestamptz",
		OID:   pgtype.TimestamptzOID,
		Codec: &pgtype.TimestamptzCodec{ScanLocation: time.UTC},
	})
	return nil
})

func GetTimezone(ctx context.Context, conn *sqlx.DB) (string, error) {
	query := "select current_setting('TIMEZONE')"
	var timezone string
	err := conn.GetContext(ctx, &timezone, query)
	if err != nil {
		return "", err
	}
	return timezone, nil
}

func IsUTC(tz string) bool {
	return tz == "Etc/UTC" || tz == "UTC"
}
