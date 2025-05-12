package config

import "time"

type ConfigHTTP struct {
	Addr           string        `yaml:"addr"            env-required:"true" env:"HTTP_ADDR"`
	APIPrefix      string        `yaml:"api_prefix"      env-required:"true" env:"HTTP_API_PREFIX"`
	HandlerTimeout time.Duration `yaml:"handler_timeout" env-required:"true" env:"HTTP_HANDLER_TIMEOUT"`
	ReadTimeout    time.Duration `yaml:"read_timeout"    env-required:"true" env:"HTTP_READ_TIMEOUT"`
	WriteTimeout   time.Duration `yaml:"write_timeout"   env-required:"true" env:"HTTP_WRITE_TIMEOUT"`
	Swagger        ConfigSwagger `yaml:"swagger"`
	Cors           ConfigCors    `yaml:"cors"`
	TLS            ConfigTLS     `yaml:"tls"`
}

func (c ConfigHTTP) GetAddr() string {
	return c.Addr
}

func (c ConfigHTTP) GetAPIPrefix() string {
	return c.APIPrefix
}

func (c ConfigHTTP) GetHandlerTimeout() time.Duration {
	return c.HandlerTimeout
}

func (c ConfigHTTP) GetReadTimeout() time.Duration {
	return c.ReadTimeout
}

func (c ConfigHTTP) GetWriteTimeout() time.Duration {
	return c.WriteTimeout
}

func (c ConfigHTTP) GetSwagger() *ConfigSwagger {
	return &c.Swagger
}

func (c ConfigHTTP) GetCORS() *ConfigCors {
	return &c.Cors
}

func (c ConfigHTTP) GetTLS() *ConfigTLS {
	return &c.TLS
}

type ConfigSwagger struct {
	Enabled *bool  `yaml:"enabled" env-required:"true" env:"HTTP_SWAGGER_ENABLED"`
	Path    string `yaml:"path"    env-required:"true" env:"HTTP_SWAGGER_PATH"`
}

func (c ConfigSwagger) GetPath() string {
	return c.Path
}

func (c ConfigSwagger) GetEnabled() bool {
	if c.Enabled == nil {
		return false
	}
	return *c.Enabled
}

type ConfigCors struct {
	AllowedOrigins []string `yaml:"allowed_origins" env-required:"true" env:"HTTP_CORS_ALLOWED_ORIGINS"`
}

func (c ConfigCors) GetAllowedOrigins() []string {
	return c.AllowedOrigins
}

type ConfigTLS struct {
	Enabled  *bool  `yaml:"enabled"   env-required:"true" env:"HTTP_TLS_ENABLED"`
	CertFile string `yaml:"cert_file" env-required:"true" env:"HTTP_TLS_CERT_FILE"`
	KeyFile  string `yaml:"key_file"  env-required:"true" env:"HTTP_TLS_KEY_FILE"`
}

func (c ConfigTLS) GetEnabled() bool {
	if c.Enabled == nil {
		return false
	}
	return *c.Enabled
}

func (c ConfigTLS) GetCertFile() string {
	return c.CertFile
}

func (c ConfigTLS) GetKeyFile() string {
	return c.KeyFile
}
