package config

import (
	"os"
)

type Config struct {
	Secret  string `yaml:"secret"`
	Address string `yaml:"address"`
	Migrations string
	DBPassword string
}

func ParseConfig() *Config {
	config := &Config{}
	config.Secret = os.Getenv("SECRET")
	config.Address = os.Getenv("ADDRESS")
	config.Migrations = os.Getenv("MIGRATIONS")
	config.DBPassword = os.Getenv("DB_PASSWORD")

	return config
}
