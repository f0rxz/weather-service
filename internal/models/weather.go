package models

import "errors"

var (
	ErrNoLocation = errors.New("Error there is no location such that")
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Localtime string `json:"localtime"`
}

type Current struct {
	TempC       float64   `json:"temp_c"`       // Temperature in Celsius
	IsDay       int       `json:"is_day"`       // 1 = day, 0 = night
	Condition   Condition `json:"condition"`    // Weather condition
	WindKPH     float64   `json:"wind_kph"`     // Wind speed in km/h
	FeelslikeC  float64   `json:"feelslike_c"`  // Feels-like temperature in Celsius
	LastUpdated string    `json:"last_updated"` //Last time updated
}

type Condition struct {
	Text string `json:"text"` // Weather condition text ("Partly cloudy")
}
