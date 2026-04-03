package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"github.com/elastic/go-elasticsearch/v9"
)

func searchHelper(client *elasticsearch.Client, searchQuery string) {
	resp, err := client.Search(
		client.Search.WithIndex("sample_web_logs"),
		client.Search.WithBody(strings.NewReader(searchQuery)),
	)
	pkg.ProcessResponse(resp, err)
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

func main() {
	client := pkg.GetClient()

	// Search for top 10 requests that look for elasticsearch
	searchQuery := `
	{
		"_source": ["@timestamp", "clientip", "request", "response", "bytes", "geo.srcdest", "url", "agent"],
		"query": {
			"match": {
				"request": {
					"query": "elasticsearch"
				}
			}
		}
	}`
	searchHelper(client, searchQuery)

	// Another search again. Exact search, with keyword search instead. No match will be found
	searchQuery = `
	{
		"_source": ["@timestamp", "clientip", "request", "response", "bytes", "geo.srcdest", "url", "agent"],
		"query": {
			"term": {
				"request.keyword": {
					"value": "elasticsearch"
				}
			}
		}
	}`
	searchHelper(client, searchQuery)

	// Search for download queries
	searchQuery = `
	{
		"_source": ["@timestamp", "clientip", "request", "response", "bytes", "geo.srcdest", "url", "agent"],
		"query": {
			"match": {
				"url": {
					"query": "downloads"
				}
			}
		},
		"size": 10
	}`
	searchHelper(client, searchQuery)

	// Multi-match, with
}
