package main

import (
	"fmt"
	"os"
)

func main() {
	/* Set and Get Environment Variables */
	if err := os.Setenv("FOO", "1"); err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("FOO"))

	if err := os.Setenv("BAR", "2"); err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("BAR"))

	/* Get all Environment Variables */
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
