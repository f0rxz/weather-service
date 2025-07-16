package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost          string `envconfig:"DB_HOST" required:"true"`
	DBPort          string `envconfig:"DB_PORT" default:"5432"`
	DBUser          string `envconfig:"DB_USER" required:"true"`
	DBPassword      string `envconfig:"DB_PASSWORD" required:"true"`
	DBName          string `envconfig:"DB_NAME" required:"true"`
	GOOSEMigrations string `envconfig:"GOOSE_MIGRATIONS" raquired:"true"`
}

func LoadConfig() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		fmt.Errorf("Failed to load configuration: %v", err)
	}
	return &cfg
}
