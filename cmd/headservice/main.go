package main

import (
	"log"
	"weatherservice/config"
	"weatherservice/internal/controller"
	"weatherservice/internal/infrastructure/connectors"
	"weatherservice/internal/infrastructure/repo"
	"weatherservice/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	db := connectors.ConnectPostgres(cfg)
	defer db.Close()

	userRepo := repo.NewUserRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo)
	weatherUC := usecase.NewWeatherUseCase()

	authController := controller.NewAuthController(authUC)
	weatherController := controller.NewWeatherController(weatherUC)

	app.Post("/auth/sign-up", authController.SignUp)
	app.Post("/auth/sign-in", authController.SignIn)
	app.Post("/auth/sign-out", authController.SignOut)
	app.Put("/auth/change-password", authController.ChangePassword)

	app.Get("/weather/:city", weatherController.GetWeather)

	log.Fatal(app.Listen(":8080"))
}
