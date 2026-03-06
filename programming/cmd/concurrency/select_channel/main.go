package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Would print "from ch2" first and then stop
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
