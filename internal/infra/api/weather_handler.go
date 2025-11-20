package api

import (
	"encoding/json"
	"net/http"

	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/usecase"
)

type WeatherHandler struct {
	apikey string
}

// type Weather struct {
// 	Location string  `json:"localidade"`
// 	Current  Current `json:"current"`
// }

// type Current struct {
// 	Temperature float64 `json:"temp_c"`
// }

func NewWeatherHandler(apikey string) *WeatherHandler {
	return &WeatherHandler{apikey: apikey}
}

// func (h *WeatherHandler) GetLocation(cep string) (Weather, error) {
// 	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", url.QueryEscape(cep))
// 	client := http.Client{}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return Weather{}, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	var weatherData Weather
// 	err = json.Unmarshal(body, &weatherData)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	return weatherData, nil
// }

// func (h *WeatherHandler) GetTemperatureCelsius(cep string) (Weather, error) {
// 	weather, err := h.GetLocation(cep)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", h.APIKey, url.QueryEscape(weather.Location))
// 	client := http.Client{}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return Weather{}, fmt.Errorf("failed to get location data: status code %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	// var weatherData Weather
// 	err = json.Unmarshal(body, &weather)
// 	if err != nil {
// 		return Weather{}, err
// 	}

// 	return weather, nil
// }

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	weatherUseCase := usecase.NewWeatherUseCaseInput(h.apikey)
	weatherCurrent, err := weatherUseCase.GetCurrentWeather(cep)
	if err != nil {
		http.Error(w, "Error fetching location", http.StatusInternalServerError)
		return
	}

	weatherCurrentJSON, err := json.Marshal(weatherCurrent)
	if err != nil {
		http.Error(w, "Error marshalling location data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(weatherCurrentJSON))
}
