package main

import (
	"errors"
	"fmt"
)

func divideTen(arg float64) (float64, error) {
	if arg == 0 {
		return -1, errors.New("Will cause zero division")
	}
	return 10.0 / arg, nil
}

func main() {
	result, err := divideTen(2.0)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(result)

	result, err = divideTen(0)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(result)

}
