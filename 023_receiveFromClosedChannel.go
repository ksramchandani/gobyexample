package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	c <- "abc"
	c <- "def"
	close(c)

	fmt.Println(<-c)
	fmt.Println(<-c)

}
