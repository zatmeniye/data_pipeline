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
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	Database string `env:"DATABASE_DATABASE"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Schema   string `env:"DATABASE_SCHEMA"`
	Type     string `env:"DATABASE_TYPE"`
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

	cfg.Database.User = os.Getenv("DATABASE_USER")
	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	cfg.Database.Host = os.Getenv("DATABASE_HOST")
	cfg.Database.Port = os.Getenv("DATABASE_PORT")
	cfg.Database.Database = os.Getenv("DATABASE_DATABASE")
	cfg.Database.Schema = os.Getenv("DATABASE_SCHEMA")
	cfg.Database.Type = os.Getenv("DATABASE_TYPE")

	return &cfg, nil
}
