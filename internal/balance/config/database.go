package config

import (
	"time"
)

type ConfigPostgres struct {
	Addr                  string        `yaml:"addr"                    env:"POSTGRES_ADDR"                    env-required:"true"`
	User                  string        `yaml:"user"                    env:"POSTGRES_USER"                    env-required:"true"`
	Password              string        `yaml:"password"                env:"POSTGRES_PASSWORD"                env-required:"true"`
	Database              string        `yaml:"database"                env:"POSTGRES_DB"                      env-required:"true"`
	Schema                string        `yaml:"schema"                  env:"POSTGRES_SCHEMA"                  env-required:"true"`
	MaxIdleConnections    int           `yaml:"max_idle_connections"    env:"POSTGRES_MAX_IDLE_CONNECTIONS"    env-required:"true"`
	MaxOpenConnections    int           `yaml:"max_open_connections"    env:"POSTGRES_MAX_OPEN_CONNECTIONS"    env-required:"true"`
	ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime" env:"POSTGRES_CONNECTION_MAX_LIFETIME" env-required:"true"`
	ReturnUTC             bool          `yaml:"return_utc"              env:"POSTGRES_RETURN_UTC"              env-required:"true"`
	MigrationsDir         string        `yaml:"migrations_dir"          env:"POSTGRES_MIGRATIONS_DIR"          env-required:"true"`
}

func (c ConfigPostgres) GetAddr() string {
	return c.Addr
}

func (c ConfigPostgres) GetUser() string {
	return c.User
}

func (c ConfigPostgres) GetPassword() string {
	return c.Password
}

func (c ConfigPostgres) GetDatabase() string {
	return c.Database
}

func (c ConfigPostgres) GetSchema() string {
	return c.Schema
}

func (c ConfigPostgres) GetMaxIdleConnections() int {
	return c.MaxIdleConnections
}

func (c ConfigPostgres) GetMaxOpenConnections() int {
	return c.MaxOpenConnections
}

func (c ConfigPostgres) GetConnectionMaxLifetime() time.Duration {
	return c.ConnectionMaxLifetime
}

func (c ConfigPostgres) GetReturnUTC() bool {
	return c.ReturnUTC
}

func (c ConfigPostgres) GetMigrationsDir() string {
	return c.MigrationsDir
}
