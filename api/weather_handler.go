package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherHandler struct {
	Weather Weather
	APIKey  string
}

type Weather struct {
	CEP      string  `json:"cep"`
	Location string  `json:"localidade"`
	Current  Current `json:"current"`
}

type Current struct {
	Temperature float64 `json:"temp_c"`
}

func NewWeatherHandler(apikey string) *WeatherHandler {
	return &WeatherHandler{APIKey: apikey}
}

func (h *WeatherHandler) GetLocation(cep string) (Weather, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return Weather{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, err
	}

	var weatherData Weather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return Weather{}, err
	}

	return weatherData, nil
}

func (h *WeatherHandler) GetTemperatureCelsius(cep string) (Weather, error) {
	weather, err := h.GetLocation(cep)
	if err != nil {
		return Weather{}, err
	}

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", h.APIKey, url.QueryEscape(weather.Location))
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return Weather{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Weather{}, err
	}

	// var weatherData Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return Weather{}, err
	}

	return weather, nil
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	weather, err := h.GetTemperatureCelsius(cep)
	if err != nil {
		http.Error(w, "Error fetching location", http.StatusInternalServerError)
		return
	}

	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		http.Error(w, "Error marshalling location data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(weatherJSON)
}
