package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
)

func (s *server) StreamSearch(req *pb.SearchRequest, stream pb.ElasticSearchService_StreamSearchServer) error {
	log.Printf("Received query: %s", req.Query)

	formattedOutputFields, err := json.Marshal(req.OutputFields)
	if err != nil {
		log.Fatal(err)
	}

	// Get the client, build query, and search
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

	documents, _ := searchHelper(esClient, searchQuery)

	for _, doc := range documents {
		err := stream.Send(&pb.Document{
			Json: doc,
		})
		if err != nil {
			return err
		}

		time.Sleep(500 * time.Millisecond) // Artificial delay, just like in real life, where there can be delays between each stream send
	}

	return nil
}
