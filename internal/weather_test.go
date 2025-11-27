package internal

import "testing"

func TestNewWeather(t *testing.T) {
	weather := NewWeather(25)

	if weather.Celsius != 25 {
		t.Fatalf("expected celsius 25, got %.2f", weather.Celsius)
	}

	if weather.Fahrenheit != 77 {
		t.Fatalf("expected fahrenheit 77, got %.2f", weather.Fahrenheit)
	}

	expectedKelvin := 298.15
	if weather.Kelvin != expectedKelvin {
		t.Fatalf("expected kelvin %.2f, got %.2f", expectedKelvin, weather.Kelvin)
	}
}
