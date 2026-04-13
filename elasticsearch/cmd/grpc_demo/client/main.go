package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewElasticSearchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.SearchByRequest(ctx, &pb.SearchRequest{
		Query:       "elasticsearch",
		SearchType:  "match",
		SearchField: "request",
		Limit:       10,
		OutputFields: []string{
			"@timestamp",
			"clientip",
			"request",
			"response",
			"bytes",
			"geo.srcdest",
			"url", "agent",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max score is: ", *resp.MaxScore)
	for _, document := range resp.Documents {
		fmt.Println(document)
	}
	fmt.Printf("\n\n")

	// Server streaming RPC call
	streamingCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.StreamSearch(streamingCtx, &pb.SearchRequest{
		Query:       "elasticsearch",
		SearchType:  "match",
		SearchField: "request",
		Limit:       10,
		OutputFields: []string{
			"@timestamp",
			"clientip",
			"request",
			"response",
			"bytes",
			"geo.srcdest",
			"url", "agent",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		doc, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(doc.Json)
	}
}
