package handlers

import (
	"weather_service/internal/controller/http/handlers/authhandler"
	"weather_service/internal/controller/http/handlers/weatherhandler"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"
)

type Handlers struct {
	authhandler    *authhandler.Handler
	weatherhandler *weatherhandler.Handler
}

func NewHandlers(auth *authusecase.AuthUseCase, weather *weatherusecase.WeatherUseCase) *Handlers {
	return &Handlers{
		authhandler:    authhandler.NewHandler(auth),
		weatherhandler: weatherhandler.NewHandler(weather),
	}
}
