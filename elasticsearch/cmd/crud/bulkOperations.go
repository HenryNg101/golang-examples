package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"github.com/elastic/go-elasticsearch/v9"
)

func bulkInsert(client *elasticsearch.Client, sourceFile string, idxName string) {
	f, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f) // Scanner, scan data file line by line
	const batchSize = 1000         // Batch size, to optimize bulk operations instead of having to call API multiple times
	var buf bytes.Buffer           // Buffer to contain current batch info
	count := 0

	for scanner.Scan() {
		doc := scanner.Bytes()

		meta := fmt.Sprintf(`{"index":{"_index":"%s"}}`, idxName)
		buf.WriteString(meta + "\n")
		buf.Write(doc)
		buf.WriteString("\n")

		count++

		if count >= batchSize {
			resp, err := client.Bulk(bytes.NewReader(buf.Bytes()))
			pkg.ProcessResponse(resp, err)
			defer resp.Body.Close()

			buf.Reset()
			count = 0
		}
	}

	// send remaining
	if buf.Len() > 0 {
		resp, err := client.Bulk(bytes.NewReader(buf.Bytes()))
		pkg.ProcessResponse(resp, err)
		defer resp.Body.Close()
	}
}

// Bulk delete
func cleanUp(client *elasticsearch.Client) {

}
