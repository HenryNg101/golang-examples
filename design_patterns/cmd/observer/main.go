package main

import (
	"fmt"
	"time"
)

type Subject struct {
	subscribers []chan string
}

func (s *Subject) Subscribe() chan string {
	ch := make(chan string)
	s.subscribers = append(s.subscribers, ch)
	return ch
}
func (s *Subject) Notify(msg string) {
	for _, ch := range s.subscribers {
		go func(c chan string) {
			c <- msg
		}(ch)
	}
}

func main() {
	subject := &Subject{}

	sub1 := subject.Subscribe()
	sub2 := subject.Subscribe()

	fmt.Println(sub1)

	go func() {
		for msg := range sub1 {
			fmt.Println("Observer 1:", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("Observer 2:", msg)
		}
	}()

	subject.Notify("event happened")
	time.Sleep(1 * time.Second)
}
