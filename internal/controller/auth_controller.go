package controller

import (
	"weatherservice/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUC *usecase.AuthUseCase
}

func NewAuthController(authUC *usecase.AuthUseCase) *AuthController {
	return &AuthController{authUC: authUC}
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	// Заглушка обработчика
	return ctx.SendStatus(fiber.StatusOK)
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {
	// Заглушка обработчика
	return ctx.SendStatus(fiber.StatusOK)
}

func (c *AuthController) SignOut(ctx *fiber.Ctx) error {
	// Заглушка обработчика
	return ctx.SendStatus(fiber.StatusOK)
}

func (c *AuthController) ChangePassword(ctx *fiber.Ctx) error {
	// Заглушка обработчика
	return ctx.SendStatus(fiber.StatusOK)
}
