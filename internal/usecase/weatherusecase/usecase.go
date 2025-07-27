package weatherusecase

import (
	"context"
	"errors"
	"weather_service/internal/infrastructure/cache/weathercache"
	"weather_service/internal/models"
	"weather_service/internal/service/weatherservice"
)

type WeatherUseCase struct {
	weatherservice *weatherservice.Service
	weathercache   *weathercache.Cache
}

func NewWeatherUseCase(weatherservice *weatherservice.Service, weathercache *weathercache.Cache) *WeatherUseCase {
	return &WeatherUseCase{
		weatherservice: weatherservice,
		weathercache:   weathercache,
	}
}

func (uc *WeatherUseCase) GetWeather(ctx context.Context, city string) (*models.WeatherResponse, error) {
	value, err := uc.weathercache.GetWeather(ctx, city)
	if err != nil && !errors.Is(err, models.ErrNoCacheCity) {
		return nil, err
	}

	if value != nil {
		return value, nil
	}

	value, err = uc.weatherservice.GetWeather(ctx, city)
	if err != nil {
		return nil, err
	}

	if err = uc.weathercache.SetWeather(ctx, city, value); err != nil {
		return nil, err
	}

	return value, nil
}
