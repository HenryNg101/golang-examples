package main

import (
	"log"
	"os"
	"strings"

	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg/pagination"
	"github.com/joho/godotenv"
)

func main() {
	// Read data from the sample data, then delete it
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	client := pkg.GetClient()
	sampleIndexName := os.Getenv("ELASTIC_DEMO_SOURCE")
	pagination.ExportDataFromIndex(client, sampleIndexName, "data.ndjson")
	mapping := pkg.ExportMappingFromIndex(client, sampleIndexName, "mapping.json")

	// Create index
	resp, err := client.Indices.Create(
		"sample_web_logs",
		client.Indices.Create.WithBody(strings.NewReader(string(mapping))),
	)
	pkg.ProcessResponse(resp, err)
	defer resp.Body.Close()

	pkg.BulkInsert(client, "data.ndjson", "sample_web_logs")

	// Delete index. Uncomment to try
	// resp, err = client.Indices.Delete([]string{"sample_web_logs"})
	// pkg.ProcessResponse(resp, err)

	// Clean index. Uncomment to try
	// pkg.CleanUp(client, "sample_web_logs")
}
