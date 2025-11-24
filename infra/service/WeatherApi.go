package service

import (
	"cepGCR/internal/dto/external"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const WeatherApiPath = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"

type WeatherApi struct {
	httpClient *http.Client
}

func NewWeatherApi(client *http.Client) *WeatherApi {
	return &WeatherApi{
		httpClient: client,
	}
}

func (b *WeatherApi) GetWeather(city string) (*external.WeatherApiResponse, error) {
	encodedCity := url.QueryEscape(city)
	formatedPath := fmt.Sprintf(WeatherApiPath, os.Getenv("WEATHER_API_KEY"), encodedCity)
	resp, err := b.httpClient.Get(formatedPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather data: status code %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var data external.WeatherApiResponse
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
