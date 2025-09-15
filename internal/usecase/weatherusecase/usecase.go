package weatherusecase

import (
	"context"
	"errors"
	"weather_service/internal/infrastructure/cache/weathercache"
	"weather_service/internal/models"
	"weather_service/internal/service/weatherservice"

	"go.uber.org/zap"
)

type WeatherUseCase interface {
	GetWeather(ctx context.Context, logger *zap.Logger, city string) (*models.WeatherResponse, error)
}

type weatherUseCase struct {
	weatherservice weatherservice.Service
	weathercache   weathercache.Cache
}

func NewWeatherUseCase(weatherservice weatherservice.Service, weathercache weathercache.Cache) WeatherUseCase {
	return &weatherUseCase{
		weatherservice: weatherservice,
		weathercache:   weathercache,
	}
}

func (uc *weatherUseCase) GetWeather(ctx context.Context, logger *zap.Logger, city string) (*models.WeatherResponse, error) {
	logger = logger.With(zap.String("WeatherUseCase", "GetWeather"))

	logger.Info("Starting to recieve weather in usecase layer.")
	value, err := uc.weathercache.GetWeather(ctx, logger, city)
	if err != nil && !errors.Is(err, models.ErrNoCacheCity) {
		logger.Error("Error in usecase layer while getting cache", zap.Error(err))
		return nil, err
	}
	if value != nil {
		return value, nil
	}

	value, err = uc.weatherservice.GetWeather(ctx, logger, city)
	if err != nil {
		logger.Error("Error while requesting weather service in usecase layer", zap.Error(err))
		return nil, err
	}

	if err = uc.weathercache.SetWeather(ctx, logger, city, value); err != nil {
		logger.Error("Error while recording weather city info in cache in usecase layer", zap.Error(err))
		return nil, err
	}

	return value, nil
}
