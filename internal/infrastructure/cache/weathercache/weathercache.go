package weathercache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"weather_service/internal/models"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache interface {
	SetWeather(ctx context.Context, logger *zap.Logger, city string, weathervalue *models.WeatherResponse) error
	GetWeather(ctx context.Context, logger *zap.Logger, city string) (*models.WeatherResponse, error)
}

type redisCache struct {
	cache *redis.Client
}

func NewWeatherCache(cache *redis.Client) Cache {
	return &redisCache{
		cache: cache,
	}
}

func (c redisCache) SetWeather(ctx context.Context, logger *zap.Logger, city string, weathervalue *models.WeatherResponse) error {
	logger = logger.With(zap.String("redisCache", "SetWeather"))

	logger.Info("Started setting cache data.")
	value, err := json.Marshal(weathervalue)
	if err != nil {
		logger.Error("Error in cache layer while json marshal", zap.Error(err))
		return err
	}
	if err := c.cache.Set(ctx, city, value, time.Minute*30).Err(); err != nil {
		logger.Error("Error while setting cache value in cache layer", zap.Error(err))
		return err
	}

	return nil
}

func (c redisCache) GetWeather(ctx context.Context, logger *zap.Logger, city string) (*models.WeatherResponse, error) {
	logger = logger.With(zap.String("RedisCache", "GetWeather"))

	logger.Info("Started getting cache data.")
	value, err := c.cache.Get(ctx, city).Result()
	if err != nil {
		logger.Error("Error in cache layer while getting cache data", zap.Error(err))
		if errors.Is(err, redis.Nil) {
			logger.Error("Error in cache layer while getting cache data and key doesnt exist", zap.Error(err))
			return nil, models.ErrNoCacheCity
		}
		return nil, err
	}

	response := models.WeatherResponse{}
	if err := json.Unmarshal([]byte(value), &response); err != nil {
		logger.Error("Error in cache layer while unmarshaling json data", zap.Error(err))
		return nil, err
	}

	return &response, nil
}
