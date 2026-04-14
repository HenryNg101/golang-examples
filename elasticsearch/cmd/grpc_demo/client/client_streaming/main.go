package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/HenryNg101/golang-examples/elasticsearch/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func bulkInsertRpc(rpcStream grpc.ClientStreamingClient[pb.BulkEntry, pb.UploadSummary], sourceFile string, idxName string) {
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
			rpcStream.Send(&pb.BulkEntry{Operations: buf.String()})
			buf.Reset()
			count = 0
		}
	}

	// send remaining
	if buf.Len() > 0 {
		rpcStream.Send(&pb.BulkEntry{Operations: buf.String()})
	}

	resp, _ := rpcStream.CloseAndRecv()
	fmt.Printf("Uploaded %d documents\n", resp.DocumentsSent)
	fmt.Printf("%d documents were successfully uploaded\n", resp.SuccessCount)
}

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
	rpcStream, err := client.BulkUpload(ctx)
	bulkInsertRpc(rpcStream, "data.ndjson", "sample_web_logs")
}
