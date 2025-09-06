package weatherusecase

import (
	"context"
	"testing"
	"weather_service/config"
	"weather_service/internal/infrastructure/cache/weathercache"
	"weather_service/internal/infrastructure/connectors"
	"weather_service/internal/service/weatherservice"

	"github.com/stretchr/testify/require"
)

func TestNewWeatherUseCase(t *testing.T) {
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

	weatherusecase := NewWeatherUseCase(weatherservice, weathercache)

	require.NotNil(t, weatherusecase)
}

func TestWeatherUseCase_GetWeather(t *testing.T) {

}
