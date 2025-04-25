package config

type ConfigGRPC struct {
	Addr string `yaml:"addr" env-required:"true"`
}

func NewConfigGRPC() *ConfigGRPC {
	return &ConfigGRPC{}
}

func (c *ConfigGRPC) GetAddr() string {
	return c.Addr
}
