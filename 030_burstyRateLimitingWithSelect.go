package main

import (
	"fmt"
	"time"
)

func worker(t1 *time.Ticker, t2 *time.Ticker, burstyLimiter chan bool) {
	for {
		select {
		case <-t1.C:
			burstyLimiter <- true
		case <-t2.C:
			for i := 0; i < 5; i++ {
				burstyLimiter <- true
			}
		}
	}
}

func main() {
	// Create a channel to which you will send and recieve data
	requests := make(chan int, 25)

	// Send Data to channel
	for i := 0; i < 25; i++ {
		requests <- i
	}
	close(requests)

	// Send to burstyLimiter first
	burstyLimiter := make(chan bool, 5)

	// Create ticker and send to burstyLimiter on every tick
	t1 := time.NewTicker(1 * time.Second)

	// Also create a ticker that will send to burstyLimiter every 5 seconds
	t2 := time.NewTicker(5 * time.Second)

	go worker(t1, t2, burstyLimiter)

	// Recieve data from channel
	for elem := range requests {
		<-burstyLimiter
		fmt.Println(elem)
	}
}
