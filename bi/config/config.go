package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Http `yaml:"http"`
	Database
}

type Http struct {
	Address string `yaml:"address"`
}

type Database struct {
	Type string `env:"DB_TYPE"`
	Dsn  string `env:"DSN"`
}

func New() (*Config, error) {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.Database.Dsn = os.Getenv("DSN")
	cfg.Database.Type = os.Getenv("DB_TYPE")

	return &cfg, nil
}
