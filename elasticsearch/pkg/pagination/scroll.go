package pagination

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

// Make use of scroll search to export all
func ExportDataFromIndex(client *elasticsearch.Client, idxName string, outputFileName string) {
	f, err := os.Create(filepath.Join(".", outputFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scrollID := ""
	for {
		var res *esapi.Response
		var err error

		// First scroll, so just do normal search with scroll
		if scrollID == "" {
			res, err = client.Search(
				client.Search.WithIndex(idxName),
				client.Search.WithSize(1000),
				client.Search.WithScroll(time.Minute),
				client.Search.WithBody(strings.NewReader(`
					{
						"query":{
							"match_all":{}
						}
					}`,
				)),
			)
		} else {
			res, err = client.Scroll(
				client.Scroll.WithScrollID(scrollID),
				client.Scroll.WithScroll(time.Minute),
			)
		}

		if err != nil {
			log.Fatal(err)
		}

		var r map[string]interface{}
		json.NewDecoder(res.Body).Decode(&r)
		res.Body.Close()

		hits := r["hits"].(map[string]interface{})["hits"].([]interface{}) // All the records that hits
		if len(hits) == 0 {
			break
		}

		scrollID = r["_scroll_id"].(string) // Get the scroll ID for next scroll search request

		for _, h := range hits {
			source := h.(map[string]interface{})["_source"]
			data, _ := json.Marshal(source) // Encode the document within the _source field
			// fmt.Print(string(data))         // write to file instead

			_, err := f.Write(data)
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString("\n")
		}
	}
}
