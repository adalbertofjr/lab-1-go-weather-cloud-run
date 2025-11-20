package domain

type Weather struct {
	Location string  `json:"localidade"`
	Temp_c   float64 `json:"temp_c"`
	Temp_f   float64 `json:"temp_f"`
	Temp_k   float64 `json:"temp_k"`
}

func NewWeather(location string, tempC float64) Weather {
	weather := Weather{
		Location: location,
		Temp_c:   tempC,
	}
	weather.calcFahrenheit()
	weather.calcKelvin()

	return weather
}

func (w *Weather) calcFahrenheit() (*Weather, error) {
	w.Temp_f = (w.Temp_c * 1.8) + 32
	return w, nil
}

func (w *Weather) calcKelvin() (*Weather, error) {
	w.Temp_k = w.Temp_c + 273.
	return w, nil
}
