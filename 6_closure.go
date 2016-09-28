package main

import (
	"fmt"
)

func nextInt() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}

func main() {
	i := nextInt()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}
