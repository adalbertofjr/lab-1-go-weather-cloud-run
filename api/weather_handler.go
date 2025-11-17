package api

import (
	"encoding/json"
	"net/http"
)

type WeatherHandler struct {
	healthcheck HealthCheck
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

type HealthCheck struct {
	Status string `json:"status"`
}

func (h *WeatherHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.healthcheck = HealthCheck{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.healthcheck)
}
