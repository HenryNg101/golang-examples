package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
	})

	defer writer.Close()

	for i := 1; i <= 20; i++ {
		orderID := fmt.Sprintf("order-%d", i)

		msg := kafka.Message{
			Key:   []byte(strconv.Itoa(i % 3)), // simulate user ID
			Value: []byte(orderID),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			panic(err)
		}

		fmt.Println("Produced:", orderID)
		time.Sleep(500 * time.Millisecond)
	}
}
