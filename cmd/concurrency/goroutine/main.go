package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go say("world") // runs concurrently
	say("hello")    // main goroutine

	time.Sleep(time.Second) // allow goroutine to finish
}
