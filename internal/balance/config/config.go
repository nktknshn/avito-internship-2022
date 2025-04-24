package config

type Config struct {
	Postgres *ConfigPostgres `yaml:"postgres"`
	JWT      *ConfigJWT      `yaml:"jwt"`
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
