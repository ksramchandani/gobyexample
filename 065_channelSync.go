package main

import (
	"fmt"
	"time"
)

func boring(msg string, c chan string) {
	// Send on channel indefinitely
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s : %d\n", msg, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Create channel
	stringChannel := make(chan string)

	// Start goroutine
	go boring("boring", stringChannel)

	// Receive 5 messages from the channel
	for i := 0; i < 5; i++ {
		fmt.Printf("In main: Received from channel : %s", <-stringChannel)
	}
}
