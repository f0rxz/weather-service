package weatherusecase_test

import (
	"context"
	"errors"
	"testing"
	"weather_service/internal/models"
	"weather_service/internal/usecase/weatherusecase"
	"weather_service/mocks"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestGetWeather_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCache := mocks.NewMockCache(ctrl)
	mockService := mocks.NewMockService(ctrl)

	expected := &models.WeatherResponse{
		Location: models.Location{
			Name:      "Paris",
			Region:    "Ile-de-France",
			Country:   "France",
			Localtime: "2025-09-07 15:00",
		},
		Current: models.Current{
			TempC:       25.0,
			IsDay:       1,
			Condition:   models.Condition{Text: "Partly cloudy"},
			WindKPH:     12.5,
			FeelslikeC:  26.3,
			LastUpdated: "2025-09-07 14:50",
		},
	}

	mockCache.
		EXPECT().
		GetWeather(ctx, zap.NewNop(), "Paris").
		Return(expected, nil)

	useCase := weatherusecase.NewWeatherUseCase(mockService, mockCache)

	result, err := useCase.GetWeather(ctx, nil, "Paris")
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestGetWeather_CacheMiss_ThenSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCache := mocks.NewMockCache(ctrl)
	mockService := mocks.NewMockService(ctrl)

	expected := &models.WeatherResponse{
		Location: models.Location{
			Name:      "Tokyo",
			Region:    "Kanto",
			Country:   "Japan",
			Localtime: "2025-09-07 18:00",
		},
		Current: models.Current{
			TempC:       30.0,
			IsDay:       1,
			Condition:   models.Condition{Text: "Sunny"},
			WindKPH:     15.0,
			FeelslikeC:  33.0,
			LastUpdated: "2025-09-07 17:50",
		},
	}

	gomock.InOrder(
		mockCache.EXPECT().GetWeather(ctx, zap.NewNop(), "Tokyo").Return(nil, models.ErrNoCacheCity),
		mockService.EXPECT().GetWeather(ctx, zap.NewNop(), "Tokyo").Return(expected, nil),
		mockCache.EXPECT().SetWeather(ctx, zap.NewNop(), "Tokyo", expected).Return(nil),
	)

	useCase := weatherusecase.NewWeatherUseCase(mockService, mockCache)

	result, err := useCase.GetWeather(ctx, nil, "Tokyo")
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestGetWeather_CacheMiss_ServiceFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCache := mocks.NewMockCache(ctrl)
	mockService := mocks.NewMockService(ctrl)

	mockCache.EXPECT().GetWeather(ctx, zap.NewNop(), "London").Return(nil, models.ErrNoCacheCity)
	mockService.EXPECT().GetWeather(ctx, zap.NewNop(), "London").Return(nil, errors.New("service error"))

	useCase := weatherusecase.NewWeatherUseCase(mockService, mockCache)

	result, err := useCase.GetWeather(ctx, zap.NewNop(), "London")
	require.Error(t, err)
	require.Nil(t, result)
	require.EqualError(t, err, "service error")
}

func TestGetWeather_CacheReturnsUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCache := mocks.NewMockCache(ctrl)
	mockService := mocks.NewMockService(ctrl)

	mockCache.EXPECT().GetWeather(ctx, zap.NewNop(), "Berlin").Return(nil, errors.New("cache failure"))

	useCase := weatherusecase.NewWeatherUseCase(mockService, mockCache)

	result, err := useCase.GetWeather(ctx, zap.NewNop(), "Berlin")
	require.Error(t, err)
	require.Nil(t, result)
	require.EqualError(t, err, "cache failure")
}

func TestGetWeather_CacheSetFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockCache := mocks.NewMockCache(ctrl)
	mockService := mocks.NewMockService(ctrl)

	expected := &models.WeatherResponse{
		Location: models.Location{
			Name:      "Oslo",
			Region:    "Oslo",
			Country:   "Norway",
			Localtime: "2025-09-07 13:00",
		},
		Current: models.Current{
			TempC:       22.0,
			IsDay:       1,
			Condition:   models.Condition{Text: "Cloudy"},
			WindKPH:     10.5,
			FeelslikeC:  21.0,
			LastUpdated: "2025-09-07 12:45",
		},
	}

	gomock.InOrder(
		mockCache.EXPECT().GetWeather(ctx, zap.NewNop(), "Oslo").Return(nil, models.ErrNoCacheCity),
		mockService.EXPECT().GetWeather(ctx, zap.NewNop(), "Oslo").Return(expected, nil),
		mockCache.EXPECT().SetWeather(ctx, zap.NewNop(), "Oslo", expected).Return(errors.New("cache write failed")),
	)

	useCase := weatherusecase.NewWeatherUseCase(mockService, mockCache)

	result, err := useCase.GetWeather(ctx, zap.NewNop(), "Oslo")
	require.Error(t, err)
	require.Nil(t, result)
	require.EqualError(t, err, "cache write failed")
}
