package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	ctx := context.Background()

	// --- Client 1: Subscriber to channel 1 ---
	subClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	subscriberClient := subClient.Subscribe(ctx, "channel1")

	// Listener goroutine (like redis-cli hanging)
	go func() {
		ch := subscriberClient.Channel()
		for msg := range ch {
			fmt.Printf("[SUBSCRIBER] %s\n", msg.Payload)
		}
	}()

	//
	//
	// --- Client 2: Normal commands ---
	publisherClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	// --- Simulate activity ---
	// Send messages to channel, with delays
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("hello %d", i)
		publisherClient.Publish(ctx, "channel1", msg)
		time.Sleep(1 * time.Second)
	}

	// Run normal commands
	publisherClient.Set(ctx, "x", 42, 0)
	val, _ := publisherClient.Get(ctx, "x").Result()
	fmt.Println("[CMD] x =", val)

	// Unsubscribe from subscriber client
	time.Sleep(2 * time.Second)
	fmt.Println("[SUBSCRIBER] Unsubscribing...")
	subscriberClient.Unsubscribe(ctx, "channel1")

	time.Sleep(2 * time.Second)
}
