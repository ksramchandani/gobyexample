package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("thread: in worker")
	time.Sleep(time.Second)
	done <- true
}

func main() {
	fmt.Println("main: start")
	done := make(chan bool, 1)

	go worker(done)

	// Receiving from channel is a blocking call.
	// If this is removed, the program will exit even before the goroutine has started execution
	fmt.Println(<-done)

	fmt.Println("Receiving from channel is a blocking call. So this will print after receiving from channel")
}
