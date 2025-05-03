package config

import "time"

type ConfigHTTP struct {
	Addr           string        `yaml:"addr" env-required:"true" env:"HTTP_ADDR"`
	ApiPrefix      string        `yaml:"api_prefix" env-required:"true" env:"HTTP_API_PREFIX"`
	HandlerTimeout time.Duration `yaml:"handler_timeout" env-required:"true" env:"HTTP_HANDLER_TIMEOUT"`
	Swagger        ConfigSwagger `yaml:"swagger"`
	Cors           ConfigCors    `yaml:"cors"`
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

func (c *ConfigHTTP) GetCors() *ConfigCors {
	return &c.Cors
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

type ConfigCors struct {
	AllowedOrigins []string `yaml:"allowed_origins" env-required:"true" env:"HTTP_CORS_ALLOWED_ORIGINS"`
}

func (c *ConfigCors) GetAllowedOrigins() []string {
	return c.AllowedOrigins
}
