package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	YP struct {
		Prefix string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DB       string
	}
	Peercast struct {
		Host string
		Port string
	}
}

var cfg Config

func FromFile(path string) (*Config, error) {
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		panic(err)
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
