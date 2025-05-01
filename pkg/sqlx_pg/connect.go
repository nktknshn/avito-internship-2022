package sqlx_pg

import (
	"context"
	"errors"
	"time"

	pfgxStdlib "github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

type PostgresCfg interface {
	GetAddr() string
	GetUserName() string
	GetPassword() string
	GetDatabase() string
	GetSchema() string
	GetMaxIdleConnections() int
	GetMaxOpenConnections() int
	GetConnectionMaxLifetime() time.Duration
	GetReturnUTC() bool
}

func Connect(ctx context.Context, cfg PostgresCfg) (*sqlx.DB, error) {

	url := "postgres://" + cfg.GetUserName() + ":" +
		cfg.GetPassword() + "@" +
		cfg.GetAddr() + "/" +
		cfg.GetDatabase()

	// устанавливаем scan location для timestampz всегда в UTC, чтобы время из базы приходило в UTC
	// https://github.com/jackc/pgx/issues/1195
	// https://github.com/jackc/pgx/pull/1948

	pgxConfig, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	opts := []pfgxStdlib.OptionOpenDB{}

	if cfg.GetReturnUTC() {
		opts = append(opts, timeZoneOption)
	}

	db := pfgxStdlib.OpenDB(*pgxConfig, opts...)

	conn := sqlx.NewDb(db, "pgx")

	tz, err := GetTimezone(ctx, conn)

	if err != nil {
		return nil, err
	}

	if tz != "UTC" {
		return nil, errors.New("postgres timezone is not UTC")
	}

	retryCount := 1

	for range retryCount {
		err = conn.PingContext(ctx)
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}

	conn.SetMaxOpenConns(cfg.GetMaxOpenConnections())
	conn.SetMaxIdleConns(cfg.GetMaxIdleConnections())
	conn.SetConnMaxLifetime(cfg.GetConnectionMaxLifetime())

	return conn, nil
}
