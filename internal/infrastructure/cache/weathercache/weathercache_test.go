package weathercache

import (
	"context"
	"testing"
	"weather_service/config"
	"weather_service/internal/infrastructure/connectors"
	"weather_service/internal/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestWeatherCache_SetWeather(t *testing.T) {
	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	ch, err := connectors.ConnectRedis(context.Background(), cfg)
	require.NoError(t, err)

	weathercache := NewWeatherCache(ch)
	require.NotNil(t, weathercache)

	var weathervalue *models.WeatherResponse
	err = gofakeit.Struct(&weathervalue)
	require.NoError(t, err)

	err = weathercache.SetWeather(context.Background(), "Tula", weathervalue)
	require.NoError(t, err)
}

func TestWeatherCache_GetWeather(t *testing.T) {
	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	ch, err := connectors.ConnectRedis(context.Background(), cfg)
	require.NoError(t, err)

	weathercache := NewWeatherCache(ch)
	require.NotNil(t, weathercache)

	var weathervalue *models.WeatherResponse
	err = gofakeit.Struct(&weathervalue)
	require.NoError(t, err)

	err = weathercache.SetWeather(context.Background(), "Tula", weathervalue)
	require.NoError(t, err)

	response, err := weathercache.GetWeather(context.Background(), "Tula")
	require.Equal(t, weathervalue, response)
	require.NoError(t, err)

	response, err = weathercache.GetWeather(context.Background(), "Хуйляндия")
	require.Nil(t, response)
	require.ErrorIs(t, err, models.ErrNoCacheCity)
}
