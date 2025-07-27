package weatherservice

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"weather_service/config"
	"weather_service/internal/models"
	"weather_service/mocks"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetWeather(t *testing.T) {
	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRT := mocks.NewMockRoundTripper(ctrl)

	service := NewService(cfg.ApiKey, mockRT)
	require.NotNil(t, service)

	weathermodel, err := json.Marshal(models.WeatherResponse{
		Location: models.Location{Name: "Tula"},
	})
	require.NoError(t, err)

	mockRT.EXPECT().
		RoundTrip(gomock.Any()).
		Return(&http.Response{Body: io.NopCloser(bytes.NewReader(weathermodel))}, nil).
		Times(1)

	result, err := service.GetWeather(context.Background(), "Tula")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "Tula", result.Location.Name)

	mockRT.EXPECT().
		RoundTrip(gomock.Any()).
		Return(&http.Response{StatusCode: 400}, nil).
		Times(1)

	result, err = service.GetWeather(context.Background(), "Artemland")
	require.ErrorIs(t, models.ErrNoLocation, err)
	require.Nil(t, result)
}
