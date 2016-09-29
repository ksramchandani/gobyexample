package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)

	c <- "string1"
	c <- "string2"

	fmt.Println(<-c)
	fmt.Println(<-c)

}
