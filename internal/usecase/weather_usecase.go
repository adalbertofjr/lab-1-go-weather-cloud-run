package usecase

import (
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/domain"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/dto"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/internal/net"
	"github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/utility"
)

type WeatherUseCase struct {
	apikey string
}

func NewWeatherUseCaseInput(apikey string) *WeatherUseCase {
	return &WeatherUseCase{apikey: apikey}
}

func (w *WeatherUseCase) GetCurrentWeather(cep string) (*dto.WeatherDTO, error) {
	cepFormated, err := utility.CEPFormatter(cep)
	if err != nil {
		return nil, err
	}

	weatherAPI := net.NewWeatherAPI(w.apikey)
	weatherAPIData, err := weatherAPI.GetTemperatureCelsius(cepFormated)
	if err != nil {
		return nil, err
	}

	currentWeather := domain.NewWeather(weatherAPIData.Location, weatherAPIData.Current.Temp_c)

	weatherDTO := dto.NewWeatherDTO(currentWeather.Location, currentWeather.Temp_c, currentWeather.Temp_f, currentWeather.Temp_k)

	return weatherDTO, nil
}
