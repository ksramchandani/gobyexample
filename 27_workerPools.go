package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	// Get jobs
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		results <- j * 2 // Send to results
	}

}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Create workers
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 0; j < 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Get from results
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}

}
