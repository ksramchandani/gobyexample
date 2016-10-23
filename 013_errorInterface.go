package main

import (
	"fmt"
)

// First define struct for that specific error handling
type argError struct {
	denominator float64
	reason      string
}

// Now let's define an Error method that operates on that struct
func (e *argError) Error() string {
	return fmt.Sprintf("%v - %s\n", e.denominator, e.reason)
}

// Let's try to use it in a function
func divideTen(den float64) (float64, error) {
	if den == 0 {
		return -1, &argError{den, "Zero Division Error"}
	}
	return 10 / den, nil
}

func main() {
	result, err := divideTen(2)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println(result)
	}

	result, err = divideTen(0)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println(result)
	}

}
