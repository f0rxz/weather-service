package main

import (
	"context"
	"log"
	"weather_service/config"
	"weather_service/internal/controller/http/handlers/authhandler"
	"weather_service/internal/controller/http/handlers/weatherhandler"
	"weather_service/internal/infrastructure/connectors"
	"weather_service/internal/infrastructure/repo/userrepo"
	"weather_service/internal/usecase/authusecase"
	"weather_service/internal/usecase/weatherusecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := connectors.ConnectPostgres(context.Background(), cfg)
	defer db.Close()

	userRepo := userrepo.NewUserRepository(db)

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
