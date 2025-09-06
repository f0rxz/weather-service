package authhandler

import (
	"weather_service/internal/usecase/authusecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	authUC *authusecase.AuthUseCase
}

func NewHandler(authUC *authusecase.AuthUseCase) *Handler {
	return &Handler{authUC: authUC}
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	// Заглушка обработчика
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	// Заглушка обработчика
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) SignOut(c *fiber.Ctx) error {
	// Заглушка обработчика
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	// Заглушка обработчика
	return c.SendStatus(fiber.StatusOK)
}
