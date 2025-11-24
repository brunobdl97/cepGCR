package main

import (
	"cepGCR/internal"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mux := http.NewServeMux()

	handler := internal.NewGetWeatherHandler()

	mux.HandleFunc("/cep/{cep}", handler.Handle)

	http.ListenAndServe(":8080", mux)
}
