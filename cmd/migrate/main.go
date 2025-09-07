package main

import (
	"weather_service/config"
	"weather_service/internal/infrastructure/connectors"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	if err = connectors.RunMigrations(cfg); err != nil {
		panic(err)
	}
}
