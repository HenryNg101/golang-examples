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
	})

	defer writer.Close()

	for i := 1; i <= 20; i++ {
		orderID := fmt.Sprintf("order-%d", i)
		msg := kafka.Message{
			Key:   []byte(strconv.Itoa(i % 3)), // simulate user ID
			Value: []byte(orderID),
			Topic: "orders",
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			panic(err)
		}

		fmt.Println("Produced:", orderID)

		paymentID := fmt.Sprintf("payment-%d", i)
		msg = kafka.Message{
			Key:   []byte(strconv.Itoa(i % 3)), // simulate user ID
			Value: []byte(paymentID),
			Topic: "payments",
		}
		err = writer.WriteMessages(context.Background(), msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("Produced:", paymentID)
		time.Sleep(500 * time.Millisecond)
	}
}
