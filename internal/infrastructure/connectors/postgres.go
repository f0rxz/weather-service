package connectors

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"weather_service/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func RunMigrations(cfg *config.Config) error {
	db, err := sql.Open("pgx", cfg.PostgresDsn)
	if err != nil {
		return err
	}

	abs, err := filepath.Abs(cfg.GooseMigrations)
	if err != nil {
		return err
	}

	goose.SetBaseFS(os.DirFS(abs))
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
