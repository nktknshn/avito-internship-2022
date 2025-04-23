package config

type Config struct {
	Postgres *ConfigPostgres `yaml:"postgres"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetPostgres() *ConfigPostgres {
	return c.Postgres
}
