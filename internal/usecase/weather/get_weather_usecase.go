package usecase

import (
	domainGateway "github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/domain/gateway"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/infra/api/dto"

	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/utility"
)

type WeatherUseCase struct {
	weatherGateway domainGateway.WeatherGateway
}

func NewWeatherUseCase(gateway domainGateway.WeatherGateway) *WeatherUseCase {
	return &WeatherUseCase{weatherGateway: gateway}
}

func (w *WeatherUseCase) GetCurrentWeather(cep string) (*dto.WeatherDTO, error) {
	cepFormated, err := utility.CEPFormatter(cep)
	if err != nil {
		return nil, err
	}

	weatherData, err := w.weatherGateway.GetCurrentWeather(cepFormated)
	if err != nil {
		return nil, err
	}

	weatherDTO := dto.NewWeatherDTO(
		weatherData.Location,
		weatherData.Temp_c,
		weatherData.Temp_f,
		weatherData.Temp_k)

	return weatherDTO, nil
}
