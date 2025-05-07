package config

import (
	"io"

	"github.com/nktknshn/avito-internship-2022/pkg/config_cleanenv"
)

type Config struct {
	HTTP     *ConfigHTTP     `yaml:"http" env-required:"true"`
	Postgres *ConfigPostgres `yaml:"postgres" env-required:"true"`
	JWT      *ConfigJWT      `yaml:"jwt" env-required:"true"`
	GRPC     *ConfigGRPC     `yaml:"grpc" env-required:"true"`
	Mode     string          `yaml:"mode" env-required:"true"`
	UseCases *ConfigUseCases `yaml:"use_cases" env-required:"true"`
	Lagging  *ConfigLagging  `yaml:"lagging" env-required:"true"`
	Jaeger   *ConfigJaeger   `yaml:"jaeger" env-required:"true"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetPostgres() *ConfigPostgres {
	return c.Postgres
}

func (c *Config) GetJWT() *ConfigJWT {
	return c.JWT
}

func (c *Config) GetHTTP() *ConfigHTTP {
	return c.HTTP
}

func (c *Config) GetGRPC() *ConfigGRPC {
	return c.GRPC
}

func (c *Config) GetMode() string {
	return c.Mode
}

func (c *Config) GetUseCases() *ConfigUseCases {
	return c.UseCases
}

func (c *Config) GetLagging() *ConfigLagging {
	return c.Lagging
}

func LoadConfigFromFile(path string) (*Config, error) {
	var cfg Config
	config_cleanenv.LoadConfig(path, &cfg)
	return &cfg, nil
}

func LoadConfigFromReader(reader io.Reader) (*Config, error) {
	var cfg Config
	config_cleanenv.ParseYAML(reader, &cfg)
	return &cfg, nil
}
