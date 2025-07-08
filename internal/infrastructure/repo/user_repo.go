package repo

import (
	"context"
	"weatherservice/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(context.Background(),
		"INSERT INTO users (login, password, salt, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Login, user.Password, user.Salt, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) UpdatePassword(id int, newPassword, newSalt string) error {
	_, err := r.db.Exec(context.Background(),
		"UPDATE users SET password = $1, salt = $2 WHERE id = $3",
		newPassword, newSalt, id)
	return err
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(context.Background(),
		"SELECT id, login, password, salt, created_at, updated_at FROM users WHERE login = $1", login).
		Scan(&user.ID, &user.Login, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
