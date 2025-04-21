package config

import (
	"time"
)

type ConfigDatabase struct {
	Addr                  string        `yaml:"addr" env:"DB_ADDR"`
	UserName              string        `yaml:"user_name" env:"DB_USER_NAME"`
	Password              string        `yaml:"password" env:"DB_PASSWORD"`
	Database              string        `yaml:"database" env:"DB_DATABASE"`
	Schema                string        `yaml:"schema" env:"DB_SCHEMA"`
	MaxIdleConnections    int           `yaml:"max_idle_connections" env:"DB_MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections    int           `yaml:"max_open_connections" env:"DB_MAX_OPEN_CONNECTIONS"`
	ConnectionMaxLifetime time.Duration `yaml:"connection_max_lifetime" env:"DB_CONNECTION_MAX_LIFETIME"`
	ReturnUTC             bool          `yaml:"return_utc" env:"DB_RETURN_UTC"`
}

func (c *ConfigDatabase) GetAddr() string {
	return c.Addr
}

func (c *ConfigDatabase) GetUserName() string {
	return c.UserName
}

func (c *ConfigDatabase) GetPassword() string {
	return c.Password
}

func (c *ConfigDatabase) GetDatabase() string {
	return c.Database
}

func (c *ConfigDatabase) GetSchema() string {
	return c.Schema
}

func (c *ConfigDatabase) GetMaxIdleConnections() int {
	return c.MaxIdleConnections
}

func (c *ConfigDatabase) GetMaxOpenConnections() int {
	return c.MaxOpenConnections
}

func (c *ConfigDatabase) GetConnectionMaxLifetime() time.Duration {
	return c.ConnectionMaxLifetime
}

func (c *ConfigDatabase) GetReturnUTC() bool {
	return c.ReturnUTC
}
