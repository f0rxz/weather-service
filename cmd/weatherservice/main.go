package main

import (
	"context"
	"weather_service/config"
	"weather_service/internal/controller/httpservice"
	"weather_service/internal/infrastructure/cache/weathercache"
	"weather_service/internal/infrastructure/connectors"
	"weather_service/internal/infrastructure/repo/userrepo"
	"weather_service/internal/service/weatherservice"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"
	"weather_service/pkg/logger"
)

func main() {
	logger, err := logger.NewZapLogger()
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	db, err := connectors.ConnectPostgres(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ch, err := connectors.ConnectRedis(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	weatherservice := weatherservice.NewService(cfg.ApiKey, nil)
	weathercache := weathercache.NewWeatherCache(ch)

	userRepo := userrepo.NewUserRepository(db)

	authUC := authusecase.NewAuthUseCase(userRepo)
	weatherUC := weatherusecase.NewWeatherUseCase(weatherservice, weathercache)

	s := httpservice.NewServer(logger, authUC, weatherUC)
	s.RunServer(":8080")
}
