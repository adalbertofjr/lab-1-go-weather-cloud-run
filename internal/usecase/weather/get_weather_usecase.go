package usecase

import (
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/domain/entity"
	domainGateway "github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/domain/gateway"

	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/utility"
)

type WeatherUseCase struct {
	weatherGateway domainGateway.WeatherGateway
}

func NewWeatherUseCase(gateway domainGateway.WeatherGateway) *WeatherUseCase {
	return &WeatherUseCase{weatherGateway: gateway}
}

func (w *WeatherUseCase) GetCurrentWeather(cep string) (*entity.Weather, error) {
	cepFormated, err := utility.CEPFormatter(cep)
	if err != nil {
		return nil, err
	}

	weatherData, err := w.weatherGateway.GetCurrentWeather(cepFormated)
	if err != nil {
		return nil, err
	}

	currentWeather := entity.NewWeather(
		weatherData.Location,
		weatherData.Temp_c)

	return currentWeather, nil
}
