package api

import (
	"encoding/json"
	"net/http"

	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/infra/api/dto"
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
		http.Error(w, err.MSG, err.Code)
		return
	}

	weatherDTO := dto.NewWeatherDTO(
		weatherCurrent.Location,
		weatherCurrent.Temp_c,
		weatherCurrent.Temp_f,
		weatherCurrent.Temp_k,
	)

	weatherCurrentJSON, jsonErr := json.Marshal(weatherDTO)
	if jsonErr != nil {
		http.Error(w, "Error marshalling location data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(weatherCurrentJSON))
}
