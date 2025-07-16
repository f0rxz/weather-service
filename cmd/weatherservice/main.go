package main

import (
	"context"
	"log"
	"weather_service/config"
	"weather_service/internal/controller/http/handlers/authhandler"
	"weather_service/internal/controller/http/handlers/weatherhandler"
	"weather_service/internal/infrastructure/connectors"
	repo "weather_service/internal/infrastructure/repo/userrepo"
	authusecase "weather_service/internal/usecase/authusecase"
	weatherusecase "weather_service/internal/usecase/weatherusecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	db := connectors.ConnectPostgres(context.Background(), cfg)
	defer db.Close()

	userRepo := repo.NewUserRepository(db)

	authUC := authusecase.NewAuthUseCase(userRepo)
	weatherUC := weatherusecase.NewWeatherUseCase()

	authHandler := authhandler.NewHandler(authUC)
	weatherHandler := weatherhandler.NewHandler(weatherUC)

	app.Post("/auth/sign-up", authHandler.SignUp)
	app.Post("/auth/sign-in", authHandler.SignIn)
	app.Post("/auth/sign-out", authHandler.SignOut)
	app.Put("/auth/change-password", authHandler.ChangePassword)

	app.Get("/weather/:city", weatherHandler.GetWeather)

	log.Fatal(app.Listen(":8080"))
}
