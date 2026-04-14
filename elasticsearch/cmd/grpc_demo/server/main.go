package main

import (
	"encoding/json"
	"log"
	"net"
	"strings"

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
