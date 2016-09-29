package main

import (
	"fmt"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s : %d\n", s, i)
	}
}

func main() {
	f("main")

	go f("thread1")

	go func(msg string) {
		fmt.Printf("thread2 : %s\n", msg)
	}("abc")

	var input string
	fmt.Printf("Enter text : ")
	fmt.Scanln(&input)
	fmt.Println(input)
}
