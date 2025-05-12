package config

import (
	"io"

	"github.com/nktknshn/avito-internship-2022/pkg/config_cleanenv"
)

type Config struct {
	HTTP     ConfigHTTP     `yaml:"http"`
	Postgres ConfigPostgres `yaml:"postgres"`
	JWT      ConfigJWT      `yaml:"jwt"`
	GRPC     ConfigGRPC     `yaml:"grpc"`
	Mode     string         `yaml:"mode"`
	UseCases ConfigUseCases `yaml:"use_cases"`
	Lagging  ConfigLagging  `yaml:"lagging"`
	Jaeger   ConfigJaeger   `yaml:"jaeger"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) GetPostgres() ConfigPostgres {
	return c.Postgres
}

func (c Config) GetJWT() ConfigJWT {
	return c.JWT
}

func (c Config) GetHTTP() ConfigHTTP {
	return c.HTTP
}

func (c Config) GetGRPC() ConfigGRPC {
	return c.GRPC
}

func (c Config) GetMode() string {
	return c.Mode
}

func (c Config) GetUseCases() ConfigUseCases {
	return c.UseCases
}

func (c Config) GetLagging() ConfigLagging {
	return c.Lagging
}

func (c Config) GetJaeger() ConfigJaeger {
	return c.Jaeger
}

func LoadConfigFromFile(path string) (*Config, error) {
	var cfg Config
	err := config_cleanenv.LoadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfigFromReader(reader io.Reader) (*Config, error) {
	var cfg Config
	err := config_cleanenv.ParseYAML(reader, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
