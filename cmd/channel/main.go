package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 10
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Channel closed")
	}
	close(ch)
	ch <- 10
	v, ok = <-ch
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Channel closed")
	}
}
