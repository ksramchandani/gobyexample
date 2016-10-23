package main

import (
	"fmt"
	"time"
)

func f(c chan string) {
	// Send to channel
	c <- "abcdef"
}

func main() {
	// Make a channel of string
	c := make(chan string)

	// Kick off goroutine that will send to channel
	go f(c)

	// Put main thread to sleep so that thread will finish execution
	time.Sleep(1 * time.Second)

	// Receive from channel
	fmt.Println(<-c)

}
