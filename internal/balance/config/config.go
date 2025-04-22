package config

type Config struct {
	Postgres *ConfigPostgres `yaml:"postgres"`
}

func (c *Config) GetPostgres() *ConfigPostgres {
	return c.Postgres
}
