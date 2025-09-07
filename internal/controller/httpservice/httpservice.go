package httpservice

import (
	"log"
	"weather_service/internal/controller/httpservice/handlers"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"
	zaplogger "weather_service/pkg/logger"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	handlers *handlers.Handlers
	logger   *zaplogger.Logger
}

func NewServer(logger *zaplogger.Logger, auth *authusecase.AuthUseCase, weather weatherusecase.WeatherUseCase) *Server {
	return &Server{
		handlers: handlers.NewHandlers(auth, weather, logger),
		logger:   logger,
	}
}

func (s Server) RunServer(port string) {
	app := fiber.New()
	app.Use(fiberlogger.New())

	s.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}

func (s Server) SetupRoutes(app *fiber.App) {
	app.Post("/auth/sign-up", s.handlers.AuthHandler.SignUp)
	app.Post("/auth/sign-in", s.handlers.AuthHandler.SignIn)

	app.Post("/auth/sign-out", s.handlers.AuthHandler.SignOut)

	app.Put("/auth/change-password", s.handlers.AuthHandler.ChangePassword)

	app.Get("/weather/:city", s.handlers.WeatherHandler.GetWeather)
}
