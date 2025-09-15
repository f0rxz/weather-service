package weatherhandler

import (
	"weather_service/internal/usecase/weatherusecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	weatherUC weatherusecase.WeatherUseCase
	logger    *zap.Logger
}

func NewHandler(logger *zap.Logger, weatherUC weatherusecase.WeatherUseCase) *Handler {
	return &Handler{
		weatherUC: weatherUC,
		logger:    logger,
	}
}

func (h *Handler) GetWeather(c *fiber.Ctx) error {
	h.logger = h.logger.With(zap.String("Handler", "GetWeather"))

	city := c.Params("city")
	h.logger.Info("Request in service layer was called.")
	result, err := h.weatherUC.GetWeather(c.Context(), h.logger, city)
	if err != nil {
		h.logger.Error("Error while requesting in service was occurred", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}
