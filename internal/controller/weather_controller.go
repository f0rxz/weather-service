package controller

import (
	"weatherservice/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type WeatherController struct {
	weatherUC *usecase.WeatherUseCase
}

func NewWeatherController(weatherUC *usecase.WeatherUseCase) *WeatherController {
	return &WeatherController{weatherUC: weatherUC}
}

func (c *WeatherController) GetWeather(ctx *fiber.Ctx) error {
	city := ctx.Params("city")
	_, err := c.weatherUC.GetWeather(city)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error")
	}
	return ctx.SendString("Weather data for " + city)
}
