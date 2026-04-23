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
	dataStreamSource := os.Getenv("ELASTIC_DATA_STREAM_SOURCE")
	esDataFile, esMappingFile, targetIndex := "es_data.ndjson", "es_mapping.json", "sample_web_logs"

	pagination.ExportDataFromDataStream(client, dataStreamSource, esDataFile)
	mapping := pkg.ExportMappingFromDataStream(client, dataStreamSource, esMappingFile)

	// Create index
	resp, err := client.Indices.Create(
		targetIndex,
		client.Indices.Create.WithBody(strings.NewReader(string(mapping))),
	)
	pkg.ProcessResponse(resp, err)
	defer resp.Body.Close()

	pkg.BulkInsert(client, esDataFile, targetIndex)

	// Delete index. Uncomment to try
	// resp, err = client.Indices.Delete([]string{targetIndex})
	// pkg.ProcessResponse(resp, err)

	// Clean index. Uncomment to try
	// pkg.CleanUp(client, targetIndex)
}
