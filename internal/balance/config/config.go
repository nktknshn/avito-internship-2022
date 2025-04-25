package config

type Config struct {
	HTTP     *ConfigHTTP     `yaml:"http" env-required:"true"`
	Postgres *ConfigPostgres `yaml:"postgres" env-required:"true"`
	JWT      *ConfigJWT      `yaml:"jwt" env-required:"true"`
	GRPC     *ConfigGRPC     `yaml:"grpc" env-required:"true"`
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
