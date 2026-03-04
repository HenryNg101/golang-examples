package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== WithCancel Demo ===")
	withCancelDemo()

	time.Sleep(1 * time.Second)

	fmt.Println("\n=== WithTimeout Demo ===")
	withTimeoutDemo()

	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Parent → Child Cancellation ===")
	parentChildDemo()

	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Context Value Demo ===")
	contextValueDemo()
}

///////////////////////////////////////////////////////////
//  WithCancel — manual cancellation
///////////////////////////////////////////////////////////

func withCancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "Worker-A")

	time.Sleep(2 * time.Second)
	fmt.Println("Main: canceling context")
	cancel()

	time.Sleep(1 * time.Second)
}

///////////////////////////////////////////////////////////
//  WithTimeout — automatic cancellation
///////////////////////////////////////////////////////////

func withTimeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // always good practice

	go worker(ctx, "Worker-B")

	time.Sleep(4 * time.Second)
}

///////////////////////////////////////////////////////////
// Parent → Child propagation
///////////////////////////////////////////////////////////

func parentChildDemo() {
	parentCtx, cancelParent := context.WithCancel(context.Background())
	childCtx, _ := context.WithTimeout(parentCtx, 5*time.Second)

	go worker(childCtx, "Child-Worker")

	time.Sleep(2 * time.Second)
	fmt.Println("Parent: canceling parent context")
	cancelParent()

	time.Sleep(1 * time.Second)
}

///////////////////////////////////////////////////////////
// Context Values
///////////////////////////////////////////////////////////

type contextKey string

func contextValueDemo() {
	ctx := context.WithValue(context.Background(), contextKey("requestID"), "ABC-123")

	go valueWorker(ctx)

	time.Sleep(1 * time.Second)
}

///////////////////////////////////////////////////////////
// Worker function respecting context
///////////////////////////////////////////////////////////

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stopped: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s working...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

///////////////////////////////////////////////////////////
// Worker reading context value
///////////////////////////////////////////////////////////

func valueWorker(ctx context.Context) {
	id := ctx.Value(contextKey("requestID"))
	fmt.Println("Worker received requestID:", id)
}
