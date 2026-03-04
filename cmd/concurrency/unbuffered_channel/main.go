package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // blocks until someone receives
	}()

	value := <-ch // blocks until someone sends
	fmt.Println(value)
}
