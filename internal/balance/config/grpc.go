package config

import "time"

type ConfigGRPC struct {
	Addr      string    `yaml:"addr" env-required:"true"`
	Keepalive Keepalive `yaml:"keepalive" env-required:"true"`
}

func NewConfigGRPC() *ConfigGRPC {
	return &ConfigGRPC{}
}

func (c *ConfigGRPC) GetAddr() string {
	return c.Addr
}

func (c *ConfigGRPC) GetKeepalive() *Keepalive {
	return &c.Keepalive
}

type Keepalive struct {
	Time                  time.Duration `yaml:"time" env-required:"true" env:"GRPC_KEEPALIVE_TIME"`
	Timeout               time.Duration `yaml:"timeout" env-required:"true" env:"GRPC_KEEPALIVE_TIMEOUT"`
	MaxConnectionIdle     time.Duration `yaml:"max_connection_idle" env-required:"true" env:"GRPC_KEEPALIVE_MAX_CONNECTION_IDLE"`
	MaxConnectionAge      time.Duration `yaml:"max_connection_age" env-required:"true" env:"GRPC_KEEPALIVE_MAX_CONNECTION_AGE"`
	MaxConnectionAgeGrace time.Duration `yaml:"max_connection_age_grace" env-required:"true" env:"GRPC_KEEPALIVE_MAX_CONNECTION_AGE_GRACE"`
}

func (c *Keepalive) GetTime() time.Duration {
	return c.Time
}

func (c *Keepalive) GetTimeout() time.Duration {
	return c.Timeout
}

func (c *Keepalive) GetMaxConnectionIdle() time.Duration {
	return c.MaxConnectionIdle
}

func (c *Keepalive) GetMaxConnectionAge() time.Duration {
	return c.MaxConnectionAge
}

func (c *Keepalive) GetMaxConnectionAgeGrace() time.Duration {
	return c.MaxConnectionAgeGrace
}
