package handlers

import (
	"weather_service/internal/controller/httpservice/handlers/authhandler"
	"weather_service/internal/controller/httpservice/handlers/weatherhandler"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"
)

type Handlers struct {
	WeatherHandler *weatherhandler.Handler
	AuthHandler    *authhandler.Handler
}

func NewHandlers(auth *authusecase.AuthUseCase, weather *weatherusecase.WeatherUseCase) *Handlers {
	return &Handlers{
		WeatherHandler: weatherhandler.NewHandler(weather),
		AuthHandler:    authhandler.NewHandler(auth),
	}
}
