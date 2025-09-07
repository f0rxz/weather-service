package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDsn     string
	GooseMigrations string
	ApiKey          string
	RedisDsn        string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		PostgresDsn:     getEnv("POSTGRES_DSN", "postgresql://postgres:postgres@localhost:5432/weather_db?sslmode=disable"),
		GooseMigrations: getEnv("GOOSE_MIGRATIONS", "./migrations"),
		ApiKey:          getEnv("API_KEY", ""),
		RedisDsn:        getEnv("REDIS_DSN", "redis://localhost:6379/0"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
