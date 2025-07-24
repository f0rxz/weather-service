package weatherhandler

import (
	"weather_service/internal/usecase/weatherusecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	weatherUC *weatherusecase.WeatherUseCase
}

func NewHandler(weatherUC *weatherusecase.WeatherUseCase) *Handler {
	return &Handler{weatherUC: weatherUC}
}

func (h *Handler) GetWeather(c *fiber.Ctx) error {
	city := c.Params("city")
	_, err := h.weatherUC.GetWeather(city)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error")
	}
	return c.SendString("Weather data for " + city)
}
