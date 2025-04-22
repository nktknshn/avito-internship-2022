package config

import (
	"time"
)

type ConfigPostgres struct {
	Addr                  string        `yaml:"addr" env:"DB_ADDR"`
	UserName              string        `yaml:"user_name" env:"DB_USER_NAME"`
	Password              string        `yaml:"password" env:"DB_PASSWORD"`
	Database              string        `yaml:"database" env:"DB_DATABASE"`
	Schema                string        `yaml:"schema" env:"DB_SCHEMA"`
	MaxIdleConnections    int           `yaml:"max_idle_connections" env:"DB_MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections    int           `yaml:"max_open_connections" env:"DB_MAX_OPEN_CONNECTIONS"`
	ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime" env:"DB_CONNECTION_MAX_LIFETIME"`
	ReturnUTC             bool          `yaml:"return_utc" env:"DB_RETURN_UTC"`
	MigrationsDir         string        `yaml:"migrations_dir" env:"DB_MIGRATIONS_DIR"`
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
