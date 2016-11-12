package main

import (
	"fmt"
	"time"
)

func main() {
	// Send to burstyLimiter first
	burstyLimiter := make(chan bool, 5)

	go func() {
		for i := 0; i < 5; i++ {
			burstyLimiter <- true
		}
	}()

	// Create ticker and send to burstyLimiter on every tick
	t := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-t.C
			burstyLimiter <- true
		}
	}()

	// Also create a ticker that will send to burstyLimiter every 5 seconds
	t2 := time.NewTicker(5 * time.Second)
	go func() {
		for {
			<-t2.C

			for i := 0; i < 5; i++ {
				burstyLimiter <- true
			}
		}
	}()

	// Now create a channel to which you will send and recieve data
	requests := make(chan int, 25)

	// Send Data to channel
	for i := 0; i < 25; i++ {
		requests <- i
	}
	close(requests)

	// Recieve data from channel
	for elem := range requests {
		<-burstyLimiter
		fmt.Println(elem)
	}
}
