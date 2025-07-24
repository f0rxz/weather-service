package weatherusecase

type WeatherUseCase struct{}

func NewWeatherUseCase() *WeatherUseCase {
	return &WeatherUseCase{}
}

func (uc *WeatherUseCase) GetWeather(city string) (string, error) {
	// Заглушка для бизнес-логики
	return "Sunny", nil
}
