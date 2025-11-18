package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/net"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/utility"
)

type WeatherHandler struct {
	CEP CEP
}

type CEP struct {
	CEP string `json:"cep"`
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (h *WeatherHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}

	cepValid, err := utility.CEPFormatter(cep)
	if err != nil {
		http.Error(w, "Invalid CEP format", http.StatusBadRequest)
		return
	}

	h.CEP = CEP{CEP: cepValid}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.CEP)
}

func (h *WeatherHandler) GetCepViaCep(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", "04446160")
	body := net.FetchData(url, w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=1e5f9bfac6b34555b7c212718251511&q=S%C3%A3o%20Paulo&aqi=no")
	body := net.FetchData(url, w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
