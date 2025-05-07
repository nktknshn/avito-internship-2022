package config

type ConfigLagging struct {
	Enabled bool `yaml:"enabled" env-required:"true" env:"LAGGING_ENABLED"`
}

func NewConfigLagging() *ConfigLagging {
	return &ConfigLagging{}
}

func (c *ConfigLagging) GetEnabled() bool {
	return c.Enabled
}
