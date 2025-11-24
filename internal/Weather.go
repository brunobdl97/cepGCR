package internal

type Weather struct {
	Fahrenheit float64
	Celsius    float64
	Kelvin     float64
}

func NewWeather(celsius float64) *Weather {
	weather := &Weather{
		Celsius: celsius,
	}

	return &Weather{
		Celsius:    celsius,
		Fahrenheit: weather.ToFahrenheit(),
		Kelvin:     weather.ToKelvin(),
	}
}

func (w *Weather) ToFahrenheit() float64 {
	return (w.Celsius * 9 / 5) + 32
}

func (w *Weather) ToKelvin() float64 {
	return w.Celsius + 273.15
}
