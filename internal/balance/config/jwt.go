package config

import "time"

type ConfigJWT struct {
	Secret string        `yaml:"secret"`
	TTL    time.Duration `yaml:"ttl"`
}

func NewConfigJWT() *ConfigJWT {
	return &ConfigJWT{}
}

func (c *ConfigJWT) GetSecret() string {
	return c.Secret
}

func (c *ConfigJWT) GetTTL() time.Duration {
	return c.TTL
}
