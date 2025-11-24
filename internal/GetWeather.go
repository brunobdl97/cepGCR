package internal

import (
	infraService "cepGCR/infra/service"
	"encoding/json"
	"net/http"
)

type GetWeatherHandler struct {
	client *http.Client
}

func NewGetWeatherHandler() *GetWeatherHandler {
	return &GetWeatherHandler{
		client: &http.Client{},
	}
}

func (h *GetWeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep, err := NewCep(r.PathValue("cep"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	cepResponse, err := infraService.NewViaCepApi(h.client).GetCep(cep.Get())
	if err != nil {
		http.Error(w, "cannot find zipcode", http.StatusNotFound)
		return
	}

	weatherResponse, err := infraService.NewWeatherApi(h.client).GetWeather(cepResponse.Localidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	weatherConversion := NewWeather(weatherResponse.Current.TempC)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	temp := WResponse{
		TempC: weatherConversion.Celsius,
		TempF: weatherConversion.Fahrenheit,
		TempK: weatherConversion.Kelvin,
	}
	json.NewEncoder(w).Encode(temp)
}

type WResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
