package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:29092"},
		GroupID:  "order-processors",
		Topic:    "orders",
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	defer reader.Close()

	for {
		// Manual commit, instead of auto commit using ReadMessage()
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			panic(err)
		}
		// Intentionally disabled commit, to see how reprocess is done
		// Normally, you process, then call CommitMessages() like below
		// err = reader.CommitMessages(context.Background(), msg)
		// if err != nil {
		// 	panic(err)
		// }

		fmt.Printf("Consumer received: %s (partition=%d, offset=%d)\n",
			string(msg.Value), msg.Partition, msg.Offset)
	}
}
