package config_cleanenv

import "github.com/ilyakaznacheev/cleanenv"

func LoadConfig(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}
