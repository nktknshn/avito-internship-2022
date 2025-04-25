package config_cleanenv

import "github.com/ilyakaznacheev/cleanenv"

func LoadConfig(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}

func LoadConfigType[T any](path string) (T, error) {
	var cfg T
	err := LoadConfig(path, &cfg)
	return cfg, err
}
