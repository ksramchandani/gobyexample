package main

import "fmt"

func sendToJobs(j chan int) {
	// Send to channel
	for i := 0; i < 3; i++ {
		j <- i
	}

	// Close channel
	close(j)
}

func receiveFromJobs(j chan int, done chan bool) {
	// Receive from channel
	for {
		r, ok := <-j
		if ok {
			fmt.Println("Received", r)
		} else {
			fmt.Println("Done receiving")
			break
		}
	}

	// Unblock main execution
	done <- true

}

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	// Kick off go routine that will send to jobs
	// Once sending is complete, the channel is closed
	go sendToJobs(jobs)

	// Kick off go routine that will receive from jobs
	// Once receive is complete, it will unblock the main execution
	go receiveFromJobs(jobs, done)

	// Block main execution until go routines have finished
	<-done

}
