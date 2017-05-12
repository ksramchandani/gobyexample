package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)

	r := []rune(s)
	r[0] = 'c'
	newString := string(r)
	fmt.Println(newString)
}
