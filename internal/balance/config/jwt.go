package config

import "time"

type ConfigJWT struct {
	Secret string        `yaml:"secret" env-required:"true" env:"JWT_SECRET"`
	TTL    time.Duration `yaml:"ttl"    env-required:"true" env:"JWT_TTL"`
}

func (c ConfigJWT) GetSecret() string {
	return c.Secret
}

func (c ConfigJWT) GetTTL() time.Duration {
	return c.TTL
}
