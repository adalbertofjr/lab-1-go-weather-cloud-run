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
	webserver.AddHandler("/health", weatherHandler.HealthCheck)
	webserver.Start()
}
