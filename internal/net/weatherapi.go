package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherAPI struct {
	apikey  string
	Weather Weather
}

type Weather struct {
	Location string  `json:"localidade"`
	Current  Current `json:"current"`
}

type Current struct {
	Temp_c float64 `json:"temp_c"`
}

func NewWeatherAPI(apikey string) *WeatherAPI {
	return &WeatherAPI{apikey: apikey}
}

func (w *WeatherAPI) getLocation(cep string) (*Weather, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", url.QueryEscape(cep))
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData Weather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func (w *WeatherAPI) GetTemperatureCelsius(cep string) (*Weather, error) {
	weather, err := w.getLocation(cep)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", w.apikey, url.QueryEscape(weather.Location))
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// var weatherData Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
