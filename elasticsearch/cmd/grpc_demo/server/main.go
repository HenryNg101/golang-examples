package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"
	"github.com/HenryNg101/golang-examples/elasticsearch/pkg"
	"github.com/elastic/go-elasticsearch/v9"

	"google.golang.org/grpc"
)

func searchHelper(client *elasticsearch.Client, searchQuery string) ([]string, float64) {
	resp, err := client.Search(
		client.Search.WithIndex("sample_web_logs"),
		client.Search.WithBody(strings.NewReader(searchQuery)),
	)
	pkg.ProcessResponse(resp, err)
	defer resp.Body.Close()

	var r map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&r)

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	max_score := r["hits"].(map[string]interface{})["max_score"]

	// Converting result, extracting hit documents, and the max score value
	convertedMaxScore, ok := max_score.(float64)
	if !ok {
		log.Fatal("Error parsing the max Score")
	}
	documents := make([]string, len(hits))
	for id, hit := range hits {
		convertedDocument, err := json.Marshal(hit)
		if err != nil {
			log.Fatal(err)
		}
		documents[id] = string(convertedDocument)
	}

	return documents, convertedMaxScore
}

type server struct {
	pb.UnimplementedElasticSearchServiceServer
}

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

func (s *server) StreamSearch(req *pb.SearchRequest, stream pb.ElasticSearchService_StreamSearchServer) error {
	log.Printf("Received query: %s", req.Query)

	formattedOutputFields, err := json.Marshal(req.OutputFields)
	if err != nil {
		log.Fatal(err)
	}

	esClient := pkg.GetClient()

	// reuse your query building logic
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

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new gRPC server, then register with
	grpcServer := grpc.NewServer()
	pb.RegisterElasticSearchServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
