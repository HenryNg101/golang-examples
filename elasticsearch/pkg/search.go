package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"
)

func SearchHelper(client *elasticsearch.Client, searchQuery string) {
	resp, err := client.Search(
		client.Search.WithIndex("sample_web_logs"),
		client.Search.WithBody(strings.NewReader(searchQuery)),
	)
	ProcessResponse(resp, err)
	defer resp.Body.Close()

	var r map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&r)

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	doc, err := json.MarshalIndent(hits, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(doc))
	fmt.Println()
}
