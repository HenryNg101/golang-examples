package main

import (
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
		Transport: &http.Transport{
			// Intentionally disable any attempt of using HTTP/2.0 for simplicity
			// Got network hanging issue before, and I checked that default config use this option as "true"
			ForceAttemptHTTP2: false,
		},
		Username: os.Getenv("ELASTIC_USER"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Test connection
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}
