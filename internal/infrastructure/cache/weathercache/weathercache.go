package weathercache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"weather_service/internal/models"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cache *redis.Client
}

func NewWeatherCache(cache *redis.Client) *Cache {
	return &Cache{
		cache: cache,
	}
}

func (c Cache) SetWeather(ctx context.Context, city string, weathervalue *models.WeatherResponse) error {
	value, err := json.Marshal(weathervalue)
	if err != nil {
		return err
	}
	if err := c.cache.Set(ctx, city, value, time.Minute*30).Err(); err != nil {
		return err
	}
	return nil
}

func (c Cache) GetWeather(ctx context.Context, city string) (*models.WeatherResponse, error) {
	value, err := c.cache.Get(ctx, city).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, models.ErrNoCacheCity
		}
		return nil, err
	}

	response := models.WeatherResponse{}
	if err := json.Unmarshal([]byte(value), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
