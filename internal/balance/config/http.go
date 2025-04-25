package config

import "time"

type ConfigHTTP struct {
	Addr           string        `yaml:"addr" env-required:"true" env:"HTTP_ADDR"`
	ApiPrefix      string        `yaml:"api_prefix" env-required:"true" env:"HTTP_API_PREFIX"`
	HandlerTimeout time.Duration `yaml:"handler_timeout" env-required:"true" env:"HTTP_HANDLER_TIMEOUT"`
}

func NewConfigHTTP() *ConfigHTTP {
	return &ConfigHTTP{}
}

func (c *ConfigHTTP) GetAddr() string {
	return c.Addr
}

func (c *ConfigHTTP) GetApiPrefix() string {
	return c.ApiPrefix
}

func (c *ConfigHTTP) GetHandlerTimeout() time.Duration {
	return c.HandlerTimeout
}
