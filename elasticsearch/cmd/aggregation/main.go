package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"github.com/elastic/go-elasticsearch/v9"
)

func aggregationHelper(client *elasticsearch.Client, searchQuery string) {
	resp, err := client.Search(
		client.Search.WithIndex("sample_web_logs"),
		client.Search.WithBody(strings.NewReader(searchQuery)),
	)
	pkg.ProcessResponse(resp, err)
	defer resp.Body.Close()

	var r map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&r)

	hits := r["aggregations"].(map[string]interface{})

	doc, err := json.MarshalIndent(hits, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(doc))
	fmt.Println()
}

func main() {
	client := pkg.GetClient()

	// Bucket aggregation. List top 10 countries that were the destination the most
	aggQuery := `
	{
		"_source": false,
		"aggs" : {
			"destination_countries" : { "terms" : { "field" : "geo.dest" } }
		}
	}`
	fmt.Println("Top 10 destination countries that receive requests the most: ")
	aggregationHelper(client, aggQuery)

	// Metric aggregation. Counting the amount of memory has travelled through the internet
	aggQuery = `
	{
		"_source": false,
		"aggs" : {
			"total_memory_travelled" : { "sum" : { "field" : "memory" } }
		}
	}`
	fmt.Println("Total memory travelled through internet: ")
	aggregationHelper(client, aggQuery)

	// Check system status, by seeing how many response statuses there were
	aggQuery = `
	{
		"_source": false,
		"aggs" : {
			"statuses" : { "terms" : { "field" : "response.keyword" } }
		}
	}`
	fmt.Println("Reponses codes count: ")
	aggregationHelper(client, aggQuery)

	// Range aggregation, classify memory travel in requests
	aggQuery = `
	{
		"_source": false,
		"aggs" : {
			"memory_usage" : { 
				"histogram" : { 
					"field" : "memory",
					"interval": 10000
				}
			}
		}
	}`
	fmt.Println("Reponses codes count: ")
	aggregationHelper(client, aggQuery)
}
