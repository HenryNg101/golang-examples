package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
)

// Implemented RPC method
func (s *server) SearchByRequest(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Printf("Received query: %s", req.Query)

	formattedOutputFields, err := json.Marshal(req.OutputFields)
	if err != nil {
		log.Fatal(err)
	}

	esClient := pkg.GetClient()
	searchQuery := fmt.Sprintf(
		`
		{
			"_source": %s,
			"size": %d,
			"query": {
				"%s": {
					"%s": {
						"query": "%s"
					}
				}
			}
		}`,
		string(formattedOutputFields),
		req.Limit,
		req.SearchType,
		req.SearchField,
		req.Query,
	)

	documents, maxScore := searchHelper(esClient, searchQuery)

	// Send the response back to client
	time.Sleep(2 * time.Second)
	return &pb.SearchResponse{
		MaxScore:  &maxScore,
		Documents: documents,
	}, nil
}
