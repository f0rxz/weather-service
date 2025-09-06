package weatherhandler

import (
	"weather_service/internal/usecase/weatherusecase"
	"weather_service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	weatherUC weatherusecase.WeatherUseCase
	logger    *logger.Logger
}

func NewHandler(weatherUC weatherusecase.WeatherUseCase) *Handler {
	return &Handler{weatherUC: weatherUC}
}

func (h *Handler) GetWeather(c *fiber.Ctx) error {
	city := c.Params("city")
	result, err := h.weatherUC.GetWeather(c.Context(), city)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}
