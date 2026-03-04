package main

import (
	"context"
	"fmt"
	"time"
)

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker stopped")
// 			return
// 		default:
// 			fmt.Println("working...")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go worker(ctx)

// 	time.Sleep(3 * time.Second)
// 	defer cancel()

// 	time.Sleep(time.Second)
// }

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker stopped:", ctx.Err())
// 			return
// 		default:
// 			fmt.Println("working...")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel() // always good practice

// 	go worker(ctx)

// 	time.Sleep(4 * time.Second)
// }

func main() {
	// Parent context: starts as an empty background context.
	parentCtx := context.Background()

	// Child context 1: derived from parent with a timeout.
	childCtx1, cancel1 := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel1() // Release resources

	// Child context 2: derived from parent with a value.
	childCtx2 := context.WithValue(parentCtx, "userID", 123)

	go func(ctx context.Context) {
		// select {
		if _, ok := <-ctx.Done(); !ok {
			// This will be triggered when parentCtx or childCtx1 is canceled/times out.
			fmt.Printf("Child 1 received cancellation signal. Error: %v\n", ctx.Err())
		}
	}(childCtx1)

	// Accessing a value from the hierarchy
	userID := childCtx2.Value("userID")
	fmt.Printf("User ID from childCtx2: %v\n", userID)

	// Wait for the timeout to trigger cancellation of childCtx1
	time.Sleep(3 * time.Second)
}
