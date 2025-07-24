package connectors

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"weather_service/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
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

func RunMigrations(cfg *config.Config) error {
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword +
		"@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return err
	}

	goose.SetBaseFS(os.DirFS(cfg.GOOSEMigrations))
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
