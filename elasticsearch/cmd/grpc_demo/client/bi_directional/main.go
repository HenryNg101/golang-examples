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

	rpcClient := pb.NewElasticSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := rpcClient.LiveSearch(ctx)

	go func() {
		for {
			doc, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(doc.Json)
		}
	}()

	// send messages
	for queryId, query := range []string{"elasticsearch", "beats"} {
		err := stream.Send(&pb.SearchRequest{
			Query:       query,
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
			log.Println("send error:", err)
		}
		fmt.Println("Query", queryId+1)
		time.Sleep(6 * time.Second) // Sleep, so that the above goroutine is triggered, and can take stream response from server
		fmt.Printf("\n\n")
	}
	defer conn.Close() // Put it down here, so when the above sleep is triggered, this isn't called and leads to errors like "rpc error: code = Canceled desc = context canceled"
}
