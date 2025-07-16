package repo

import (
	"context"
	"weather_service/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := pgx.NamedArgs{
		"login":      user.Login,
		"password":   user.Password,
		"salt":       user.Salt,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	_, err := r.db.Exec(ctx,
		`INSERT INTO users (login, password, salt, created_at, updated_at) 
		 VALUES (@login, @password, @salt, @created_at, @updated_at)`,
		args,
	)
	return err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id int, newPassword, newSalt string) error {
	args := pgx.NamedArgs{
		"password": newPassword,
		"salt":     newSalt,
		"id":       id,
	}

	_, err := r.db.Exec(ctx,
		`UPDATE users 
		 SET password = @password, salt = @salt 
		 WHERE id = @id`,
		args,
	)
	return err
}

func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	user := &models.User{}
	args := pgx.NamedArgs{"login": login}

	err := r.db.QueryRow(ctx,
		`SELECT id, login, password, salt, created_at, updated_at 
		 FROM users 
		 WHERE login = @login`,
		args,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}
