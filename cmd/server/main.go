package main

import (
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/cmd/configs"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/infra/api"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/infra/web"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	startServer(configs)
}

func startServer(configs *configs.Conf) {
	webserver := web.NewWebServer(configs.WebServerPort)
	weatherHandler := api.NewWeatherHandler(configs.WeatherAPIKey)
	healthHandler := api.NewHealthCheck()
	webserver.AddHandler("/", weatherHandler.GetWeather)
	webserver.AddHandler("/health", healthHandler.HealthCheck)
	webserver.Start()
}
