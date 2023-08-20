package configs

import (
	"fmt"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerPort  string `env:"PORT"`
	Environment string `env:"ENV"`
	Debug       bool   `env:"DEBUG"`
	LogOutput   string `env:"LOG_OUTPUT"`
	LogLevel    string `env:"LOG_LEVEL"`
	PostgresDSN string `env:"POSTGRES_DSN"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
