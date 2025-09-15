package main

import (
	"context"
	"weather_service/config"
	"weather_service/internal/controller/httpservice"
	"weather_service/internal/infrastructure/cache/weathercache"
	"weather_service/internal/infrastructure/connectors"
	"weather_service/internal/service/weatherservice"
	"weather_service/internal/usecase/weatherusecase"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	ch, err := connectors.ConnectRedis(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	weatherservice := weatherservice.NewService(cfg.ApiKey, nil)
	weathercache := weathercache.NewWeatherCache(ch)

	weatherUC := weatherusecase.NewWeatherUseCase(weatherservice, weathercache)

	s := httpservice.NewServer(logger, weatherUC)
	s.RunServer(":8080")
}
