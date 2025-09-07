package weatherservice

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"weather_service/internal/models"
)

type Service interface {
	GetWeather(ctx context.Context, data string) (*models.WeatherResponse, error)
}

type service struct {
	apiKey    string
	transport http.RoundTripper
}

func NewService(apiKey string, customTransport http.RoundTripper) Service {
	if customTransport == nil {
		customTransport = http.DefaultTransport
	}

	return &service{
		apiKey:    apiKey,
		transport: customTransport,
	}
}

func (s service) GetWeather(ctx context.Context, data string) (*models.WeatherResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Add("key", s.apiKey)
	values.Add("q", data)
	req.URL.RawQuery = values.Encode()

	client := http.Client{Transport: s.transport}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusBadRequest {
		return nil, models.ErrNoLocation
	}

	result := &models.WeatherResponse{}
	if err = json.NewDecoder(res.Body).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
