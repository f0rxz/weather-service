package usecase

import (
	"weatherservice/internal/infrastructure/repo"
)

type AuthUseCase struct {
	userRepo *repo.UserRepository
}

func NewAuthUseCase(userRepo *repo.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) SignUp(login, password string) error {
	// Заглушка для бизнес-логики
	return nil
}

func (uc *AuthUseCase) SignIn(login, password string) error {
	// Заглушка для бизнес-логики
	return nil
}

func (uc *AuthUseCase) SignOut() error {
	// Заглушка для бизнес-логики
	return nil
}

func (uc *AuthUseCase) ChangePassword(userID int, newPassword string) error {
	// Заглушка для бизнес-логики
	return nil
}
