package main

import (
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/api"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/cmd/configs"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/infra/web"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	_ = configs
	startServer(configs)
}

func startServer(configs *configs.Conf) {
	webserver := web.NewWebServer(":8000")
	weatherHandler := api.NewWeatherHandler(configs.WeatherAPIKey)
	healthHandler := api.NewHealthCheck()
	webserver.AddHandler("/", weatherHandler.GetWeather)
	webserver.AddHandler("/health", healthHandler.HealthCheck)
	webserver.Start()
}
