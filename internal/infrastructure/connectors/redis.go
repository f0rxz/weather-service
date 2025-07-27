package connectors

import (
	"context"
	"weather_service/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	options, err := redis.ParseURL(cfg.RedisDsn)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(options), nil
}
