package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	limiter := time.Tick(1 * time.Second)

	// Send to requests first
	for i := 10; i < 15; i++ {
		requests <- i
	}
	close(requests)

	// Now receive from requests with blocking on rate limiting
	for elem := range requests {
		<-limiter
		fmt.Println(elem)
	}

}
