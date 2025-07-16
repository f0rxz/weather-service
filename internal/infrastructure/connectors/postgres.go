package connectors

import (
	"context"
	"fmt"
	"os"

	"weather_service/config"

	"github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

func ConnectPostgres(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword +
		"@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	return pool
}

func RunMigrations(pool *pgxpool.Pool, migrationsDir string) error {
	// Конвертируем pgxpool в *sql.DB
	conn := pool.Config().ConnConfig
	db := stdlib.OpenDB(*conn)

	goose.SetBaseFS(os.DirFS(migrationsDir))
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
