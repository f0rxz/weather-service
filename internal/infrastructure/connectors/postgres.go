package connectors

import (
	"context"

	"weatherservice/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectPostgres(cfg *config.Config) *pgxpool.Pool {
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword +
		"@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return pool
}
