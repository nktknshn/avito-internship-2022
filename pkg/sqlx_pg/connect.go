package sqlx_pg

import (
	"context"
	"time"

	pfgxStdlib "github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

type postgresCfgI interface {
	GetAddr() string
	GetUserName() string
	GetPassword() string
	GetDatabase() string
	GetSchema() string
	GetMaxIdleConnections() int
	GetMaxOpenConnections() int
	GetConnectionMaxLifetime() time.Duration
	GetReturnUTC() bool
	// GetUpMigrations() db.DbUpMigrations
	// GetMigrationDir() db.DbMigrationsDirectory
}

func Connect(ctx context.Context, cfg postgresCfgI) (*sqlx.DB, error) {

	url := "postgres://" + cfg.GetUserName() + ":" +
		cfg.GetPassword() + "@" +
		cfg.GetAddr() + "/" +
		cfg.GetDatabase()

	// устанавливаем scan location для timestampz всегда в UTC, чтобы время из базы приходило в UTC
	// https://github.com/jackc/pgx/issues/1195
	// https://github.com/jackc/pgx/pull/1948
	timeZoneOption := pfgxStdlib.OptionAfterConnect(func(ctx context.Context, c *pgx.Conn) error {
		c.TypeMap().RegisterType(&pgtype.Type{
			Name:  "timestamptz",
			OID:   pgtype.TimestamptzOID,
			Codec: &pgtype.TimestamptzCodec{ScanLocation: time.UTC},
		})
		return nil
	})

	pgxConfig, err := pgx.ParseConfig(url)
	if err != nil {
		// log.Errorf("pgx.ParseConfig(url): %v", err)
		return nil, err
	}

	opts := []pfgxStdlib.OptionOpenDB{}

	if cfg.GetReturnUTC() {
		opts = append(opts, timeZoneOption)
	}

	db := pfgxStdlib.OpenDB(*pgxConfig, opts...)

	conn := sqlx.NewDb(db, "pgx")

	err = conn.PingContext(ctx)
	if err != nil {
		// log.Errorf("db.PingContext(ctx): %v", err)
		return nil, err
	}
	conn.SetMaxOpenConns(cfg.GetMaxOpenConnections())
	conn.SetMaxIdleConns(cfg.GetMaxIdleConnections())
	conn.SetConnMaxLifetime(cfg.GetConnectionMaxLifetime())

	//накатываем миграции
	// if cfg.GetUpMigrations().Bool() {
	// 	err = c.UpMigrations(ctx, cfg.GetMigrationDir())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return conn, nil
}
