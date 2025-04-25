package config

import (
	"time"
)

type ConfigPostgres struct {
	Addr                  string        `yaml:"addr" env:"DB_ADDR" env-required:"true"`
	UserName              string        `yaml:"user_name" env:"DB_USER_NAME" env-required:"true"`
	Password              string        `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	Database              string        `yaml:"database" env:"DB_DATABASE" env-required:"true"`
	Schema                string        `yaml:"schema" env:"DB_SCHEMA" env-required:"true"`
	MaxIdleConnections    int           `yaml:"max_idle_connections" env:"DB_MAX_IDLE_CONNECTIONS" env-required:"true"`
	MaxOpenConnections    int           `yaml:"max_open_connections" env:"DB_MAX_OPEN_CONNECTIONS" env-required:"true"`
	ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime" env:"DB_CONNECTION_MAX_LIFETIME" env-required:"true"`
	ReturnUTC             bool          `yaml:"return_utc" env:"DB_RETURN_UTC" env-required:"true"`
	MigrationsDir         string        `yaml:"migrations_dir" env:"DB_MIGRATIONS_DIR" env-required:"true"`
}

func (c *ConfigPostgres) GetAddr() string {
	return c.Addr
}

func (c *ConfigPostgres) GetUserName() string {
	return c.UserName
}

func (c *ConfigPostgres) GetPassword() string {
	return c.Password
}

func (c *ConfigPostgres) GetDatabase() string {
	return c.Database
}

func (c *ConfigPostgres) GetSchema() string {
	return c.Schema
}

func (c *ConfigPostgres) GetMaxIdleConnections() int {
	return c.MaxIdleConnections
}

func (c *ConfigPostgres) GetMaxOpenConnections() int {
	return c.MaxOpenConnections
}

func (c *ConfigPostgres) GetConnectionMaxLifetime() time.Duration {
	return c.ConnectionMaxLifetime
}

func (c *ConfigPostgres) GetReturnUTC() bool {
	return c.ReturnUTC
}

func (c *ConfigPostgres) GetMigrationsDir() string {
	return c.MigrationsDir
}
