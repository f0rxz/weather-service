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

func NewHandler(logger *logger.Logger, weatherUC weatherusecase.WeatherUseCase) *Handler {
	return &Handler{
		weatherUC: weatherUC,
		logger:    logger,
	}
}

func (h *Handler) GetWeather(c *fiber.Ctx) error {
	city := c.Params("city")
	result, err := h.weatherUC.GetWeather(c.Context(), city)
	if err != nil {
		h.logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	h.logger.Info("Request weather completed successfully.")
	return c.JSON(result)
}
