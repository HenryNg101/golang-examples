package main

import (
	"bytes"
	"encoding/json"
	"log"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"google.golang.org/grpc"
)

func (s *server) BulkUpload(stream grpc.ClientStreamingServer[pb.BulkEntry, pb.UploadSummary]) error {
	esClient := pkg.GetClient()
	documentsSent := int32(0)
	successesCount := int32(0)

	for {
		bulkEntry, err := stream.Recv()
		if err != nil {
			break
		}

		resp, err := esClient.Bulk(bytes.NewReader([]byte(bulkEntry.Operations)))
		pkg.ProcessResponse(resp, err)
		defer resp.Body.Close()

		var r map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&r)

		bulkResults := r["items"].([]interface{})

		sentDocs, successes := 0, 0
		for _, result := range bulkResults {
			index := result.(map[string]interface{})["index"]
			statusCode, ok := index.(map[string]interface{})["status"].(float64)
			if !ok {
				log.Fatal("Can't parse status code value")
			}
			sentDocs++
			// documentsSent++
			if statusCode >= 200 && statusCode < 300 {
				successes++
				// successesCount++
			}
		}
		documentsSent += int32(sentDocs)
		successesCount += int32(successes)
		log.Printf("Proccessed %d documents. %d documents were successfully processed", sentDocs, successes)
	}

	return stream.SendAndClose(&pb.UploadSummary{
		SuccessCount:  successesCount,
		DocumentsSent: documentsSent,
	})
}
