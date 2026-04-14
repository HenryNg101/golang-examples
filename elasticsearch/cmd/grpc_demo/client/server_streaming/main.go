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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.StreamSearch(ctx, &pb.SearchRequest{
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
