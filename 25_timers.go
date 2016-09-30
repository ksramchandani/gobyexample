package main

import (
	"fmt"
	"time"
)

func stopTimer(timer2 *time.Timer, done chan bool) {
	timer2.Stop()
	fmt.Println("Timer2 stopped")
	done <- true
}

func main() {
	done := make(chan bool, 1)

	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")

	timer2 := time.NewTimer(10 * time.Second)
	stopTimer(timer2, done)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
		done <- true
	}()

	<-done
}
