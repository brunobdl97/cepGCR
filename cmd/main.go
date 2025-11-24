package main

import (
	"cepGCR/internal"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	handler := internal.NewGetWeatherHandler()

	mux.HandleFunc("/cep/{cep}", handler.Handle)

	http.ListenAndServe(":8080", mux)
}
