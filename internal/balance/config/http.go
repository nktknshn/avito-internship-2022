package config

import "time"

type ConfigHTTP struct {
	Addr           string        `yaml:"addr" env-required:"true" env:"HTTP_ADDR"`
	ApiPrefix      string        `yaml:"api_prefix" env-required:"true" env:"HTTP_API_PREFIX"`
	HandlerTimeout time.Duration `yaml:"handler_timeout" env-required:"true" env:"HTTP_HANDLER_TIMEOUT"`
	Swagger        ConfigSwagger `yaml:"swagger"`
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

func (c *ConfigHTTP) GetSwagger() *ConfigSwagger {
	return &c.Swagger
}

type ConfigSwagger struct {
	Enabled bool   `yaml:"enabled" env-required:"true" env:"HTTP_SWAGGER_ENABLED"`
	Path    string `yaml:"path" env-required:"true" env:"HTTP_SWAGGER_PATH"`
}

func (c *ConfigSwagger) GetPath() string {
	return c.Path
}

func (c *ConfigSwagger) GetEnabled() bool {
	return c.Enabled
}
