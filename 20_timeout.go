package main

import (
	"fmt"
	"time"
)

func send(c chan string, msg string) {
	time.Sleep(2 * time.Second)
	c <- msg
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go send(c1, "c1")
	go send(c2, "c2")

	select {
	case msg := <-c1:
		fmt.Println("Received on channel 1:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout on channel 1")
	}

	select {
	case msg := <-c2:
		fmt.Println("Received on channel 2:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout on channel 2")
	}

}
