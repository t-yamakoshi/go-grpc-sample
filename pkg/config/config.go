package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	MYSQL_HOST     string `env:"MYSQL_HOST" envDefault:"localhost"`
	MYSQL_PORT     string `env:"MYSQL_PORT" envDefault:"3306"`
	MYSQL_USER     string `env:"MYSQL_USER" envDefault:"root"`
	MYSQL_PASSWORD string `env:"MYSQL_PASS"`
	MYSQL_DB       string `env:"MYSQL_DB" envDefault:"todo"`
	APP_PORT       int    `env:"APP_PORT" envDefault:"443"`
}

func NewConfig() (*Config, error) {
	if err := loadEnvFile(".env"); err != nil {
		return nil, fmt.Errorf("failed to load env file: %w", err)
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env vars: %w", err)
	}
	return &cfg, nil
}

func loadEnvFile(filepath string) error {
	err := godotenv.Load(filepath)
	if err != nil {
		return fmt.Errorf("failed to load env file: %w", err)
	}

	return nil
}
