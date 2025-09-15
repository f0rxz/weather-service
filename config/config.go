package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PostgresDsn     string `envconfig:"POSTGRES_DSN" default:"postgresql://postgres:postgres@localhost:5432/weather_db?sslmode=disable"`
	GooseMigrations string `envconfig:"GOOSE_MIGRATIONS" default:"./migrations"`
	ApiKey          string `envconfig:"API_KEY" default:""`
	RedisDsn        string `envconfig:"REDIS_DSN" default:"redis://localhost:6379/0"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
