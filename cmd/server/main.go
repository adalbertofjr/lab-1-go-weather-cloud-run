package main

import (
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/api"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/infra/web"
)

func main() {
	startServer()
}

func startServer() {
	webserver := web.NewWebServer(":8000")
	weatherHandler := api.NewWeatherHandler()
	healthHandler := api.NewHealthCheck()
	webserver.AddHandler("/", weatherHandler.GetWeather)
	webserver.AddHandler("/health", healthHandler.HealthCheck)
	webserver.Start()
}
