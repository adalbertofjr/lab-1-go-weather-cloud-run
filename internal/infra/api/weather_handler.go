package api

import (
	"encoding/json"
	"net/http"

	usecase "github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/usecase/weather"
)

type WeatherHandler struct {
	usecase *usecase.WeatherUseCase
}

func NewWeatherHandler(useCase *usecase.WeatherUseCase) *WeatherHandler {
	return &WeatherHandler{usecase: useCase}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	weatherCurrent, err := h.usecase.GetCurrentWeather(cep)
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
