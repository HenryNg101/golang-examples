package pkg

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v9"
)

func ExportMappingFromDataStream(client *elasticsearch.Client, streamName string, outputFileName string) []byte {
	f, err := os.Create(filepath.Join(".", outputFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := client.Indices.GetMapping(client.Indices.GetMapping.WithIndex(streamName))
	var r map[string]interface{}
	json.NewDecoder(res.Body).Decode(&r)

	// Since ES data streams can contain more than one source index, we need to get it
	// In this demo, the dataset is simple, it only contains one index, so we just need to get the first one
	var sourceIndex string
	for k := range r {
		sourceIndex = k
		break
	}

	mapping := r[sourceIndex]
	data, _ := json.MarshalIndent(mapping, "", "\t")
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
