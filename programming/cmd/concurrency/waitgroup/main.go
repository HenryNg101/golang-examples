package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			fmt.Println("Worker", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
