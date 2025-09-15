package httpservice

import (
	"log"
	"weather_service/internal/controller/httpservice/handlers"
	"weather_service/internal/usecase/weatherusecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	handlers *handlers.Handlers
	logger   *zap.Logger
}

func NewServer(logger *zap.Logger, weather weatherusecase.WeatherUseCase) *Server {
	return &Server{
		handlers: handlers.NewHandlers(weather, logger),
		logger:   logger,
	}
}

func (s Server) RunServer(port string) {
	app := fiber.New()

	s.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}

func (s Server) SetupRoutes(app *fiber.App) {
	app.Get("/weather/:city", s.handlers.WeatherHandler.GetWeather)
}
