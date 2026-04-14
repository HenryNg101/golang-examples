package pkg

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
)

func BulkInsert(client *elasticsearch.Client, sourceFile string, idxName string) {
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
			ProcessResponse(resp, err)
			defer resp.Body.Close()

			buf.Reset()
			count = 0
		}
	}

	// send remaining
	if buf.Len() > 0 {
		resp, err := client.Bulk(bytes.NewReader(buf.Bytes()))
		ProcessResponse(resp, err)
		defer resp.Body.Close()
	}
}

// Bulk delete
func CleanUp(client *elasticsearch.Client, idxName string) {
	query := `{"query": {"match_all": {}}}`
	resp, err := client.DeleteByQuery(
		[]string{idxName},
		bytes.NewReader([]byte(query)),
		client.DeleteByQuery.WithContext(context.Background()),
		client.DeleteByQuery.WithRefresh(true), // Forces a refresh after the operation to make changes visible immediately
	)
	ProcessResponse(resp, err)
	defer resp.Body.Close()
}
