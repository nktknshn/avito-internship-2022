package testing_pg

import (
	"time"
)

type config struct {
	Addr                  string
	User                  string
	Password              string
	Database              string
	Schema                string
	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime time.Duration
	UpMigrations          bool
	ReturnUTC             bool
}

func (cfg *config) GetAddr() string {
	return cfg.Addr
}

func (cfg *config) GetUser() string {
	return cfg.User
}

func (cfg *config) GetPassword() string {
	return cfg.Password
}

func (cfg *config) GetDatabase() string {
	return cfg.Database
}

func (cfg *config) GetSchema() string {
	return cfg.Schema
}

func (cfg *config) GetMaxIdleConnections() int {
	return cfg.MaxIdleConnections
}

func (cfg *config) GetMaxOpenConnections() int {
	return cfg.MaxOpenConnections
}

func (cfg *config) GetConnectionMaxLifetime() time.Duration {
	return cfg.ConnectionMaxLifetime
}

func (cfg *config) GetReturnUTC() bool {
	return cfg.ReturnUTC
}
