package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var config *Config

func C() *Config {
	if config == nil {
		panic("load config first")
	}
	return config
}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}

func LoadConfigFromToml(filepath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return err
	}
	return nil
}
