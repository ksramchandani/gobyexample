package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	// Fire off go routine to Tick every second
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	// Stop ticker after waiting 5 seconds
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
