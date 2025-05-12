package config

type ConfigLagging struct {
	Enabled *bool `yaml:"enabled" env-required:"true" env:"LAGGING_ENABLED"`
}

func (c ConfigLagging) GetEnabled() bool {
	if c.Enabled == nil {
		return false
	}
	return *c.Enabled
}
