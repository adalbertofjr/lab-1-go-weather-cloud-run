package dto

type WeatherDTO struct {
	Location string  `json:"localidade"`
	Temp_c   float64 `json:"temp_c"`
	Temp_f   float64 `json:"temp_f"`
	Temp_k   float64 `json:"temp_k"`
}

func NewWeatherDTO(location string, tempC, tempF, tempK float64) *WeatherDTO {
	return &WeatherDTO{
		Location: location,
		Temp_c:   tempC,
		Temp_f:   tempF,
		Temp_k:   tempK,
	}
}
