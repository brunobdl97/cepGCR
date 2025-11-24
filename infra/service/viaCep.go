package service

import (
	"cepGCR/internal/dto/external"
	"encoding/json"
	"fmt"
	"net/http"
)

const ViaCepPath = "http://viacep.com.br/ws/%s/json/"

type ViaCepApi struct {
	httpClient *http.Client
}

func NewViaCepApi(client *http.Client) *ViaCepApi {
	return &ViaCepApi{
		httpClient: client,
	}
}

func (b *ViaCepApi) GetCep(cep string) (*external.ViaCepResponse, error) {
	formatedPath := fmt.Sprintf(ViaCepPath, cep)
	resp, err := b.httpClient.Get(formatedPath)
	if err != nil {
		return nil, err
	}

	var data external.ViaCepResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
