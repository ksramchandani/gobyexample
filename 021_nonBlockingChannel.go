package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 1)

	select {
	case msg := <-c1:
		fmt.Println(msg)
	default:
		fmt.Println("no message received on first try")
	}

	c1 <- "abc"
	select {
	case msg := <-c1:
		fmt.Println("received on second try:", msg)
	default:
		fmt.Println("no message received on second try")
	}
}
