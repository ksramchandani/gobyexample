package main

import (
	"fmt"
	"time"
)

func send(c chan string, msg string) {
	time.Sleep(time.Second)
	c <- msg
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go send(c1, "c1")
	go send(c2, "c2")

	// Select is a blocking call and it will block until one of the cases becomes true
	// Here select will block until we receive on either of the 2 channels
	// Let's say we receive on c1. So select will be unblocked and it print "c1"
	// But we still have to wait to receive on c2. Hence, the for loop

	for i := 0; i < 2; i++ {
		fmt.Println("Will block on select now")
		select {
		case msg := <-c1:
			fmt.Println("Received from channel 1", msg)

		case msg := <-c2:
			fmt.Println("Received from channel 2", msg)
		}
	}
}
