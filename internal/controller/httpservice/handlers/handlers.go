package handlers

import (
	"weather_service/internal/controller/httpservice/handlers/authhandler"
	"weather_service/internal/controller/httpservice/handlers/weatherhandler"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"
	"weather_service/pkg/logger"
)

type Handlers struct {
	WeatherHandler *weatherhandler.Handler
	AuthHandler    *authhandler.Handler
}

func NewHandlers(auth *authusecase.AuthUseCase, weather weatherusecase.WeatherUseCase, logger *logger.Logger) *Handlers {
	return &Handlers{
		WeatherHandler: weatherhandler.NewHandler(logger, weather),
		AuthHandler:    authhandler.NewHandler(auth),
	}
}
