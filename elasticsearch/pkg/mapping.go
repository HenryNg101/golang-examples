package pkg

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v9"
)

func ExportMappingFromIndex(client *elasticsearch.Client, idxName string, outputFileName string) []byte {
	f, err := os.Create(filepath.Join(".", outputFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := client.Indices.GetMapping(client.Indices.GetMapping.WithIndex(idxName))
	var r map[string]interface{}
	json.NewDecoder(res.Body).Decode(&r)

	mapping := r[idxName]
	data, _ := json.MarshalIndent(mapping, "", "\t")
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
