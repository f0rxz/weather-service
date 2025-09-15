package handlers

import (
	"weather_service/internal/controller/httpservice/handlers/weatherhandler"
	"weather_service/internal/usecase/weatherusecase"

	"go.uber.org/zap"
)

type Handlers struct {
	WeatherHandler *weatherhandler.Handler
}

func NewHandlers(weather weatherusecase.WeatherUseCase, logger *zap.Logger) *Handlers {
	return &Handlers{
		WeatherHandler: weatherhandler.NewHandler(logger, weather),
	}
}
