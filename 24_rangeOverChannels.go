package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	c <- "abc"
	c <- "def"
	close(c)

	for elem := range c {
		fmt.Println(elem)
	}
}
