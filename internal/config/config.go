package config

type Config struct {
	Database *ConfigDatabase `yaml:"database"`
}

func (c *Config) GetDatabase() *ConfigDatabase {
	return c.Database
}
