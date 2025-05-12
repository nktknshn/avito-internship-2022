package config

type ConfigJaeger struct {
	Host        string `yaml:"host"         env-required:"true" env:"JAEGER_HOST"`
	ServiceName string `yaml:"service_name" env-required:"true" env:"JAEGER_SERVICE_NAME"`
	LogSpans    bool   `yaml:"log_spans"    env-required:"true" env:"JAEGER_LOG_SPANS"`
}

func (c ConfigJaeger) GetHost() string {
	return c.Host
}

func (c ConfigJaeger) GetServiceName() string {
	return c.ServiceName
}

func (c ConfigJaeger) GetLogSpans() bool {
	return c.LogSpans
}
