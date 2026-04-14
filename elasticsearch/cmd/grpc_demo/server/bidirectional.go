package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) LiveSearch(stream grpc.BidiStreamingServer[pb.SearchRequest, pb.Document]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("client closed stream")
				return nil
			}

			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Canceled {
				log.Println("client canceled connection")
				return nil
			}

			log.Println("unexpected error:", err)
			return err
		}
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
	}
}
