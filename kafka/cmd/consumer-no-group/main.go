package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "orders",
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Printf("Consumer received: %s (partition=%d, offset=%d)\n",
			string(msg.Value), msg.Partition, msg.Offset)
	}
}
