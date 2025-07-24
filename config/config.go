package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost          string `envconfig:"DB_HOST" default:"localhost"`
	DBPort          string `envconfig:"DB_PORT" default:"5432"`
	DBUser          string `envconfig:"DB_USER" default:"postgres"`
	DBPassword      string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName          string `envconfig:"DB_NAME" default:"weather_db"`
	GOOSEMigrations string `envconfig:"GOOSE_MIGRATIONS" default:"./migrations"`
	APIKey          string `envconfig:"API_KEY" default:""`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
