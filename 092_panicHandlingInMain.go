package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start Main")
	if err := test(); err != nil {
		if err.Error() == "Mimic Panic" {
			// Graceful handling of Panic
			fmt.Println("Gracefully Handling Panic")
		} else {
			fmt.Println("Error Handling")
		}

	}
	fmt.Println("End Main")
}

func mimicError(key string) error {
	return fmt.Errorf("Error %s\n", key)
}

func test() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic error %v\n", r)
			err = fmt.Errorf("%v", r) // Set err to reflect there was a panic in the code
		}
	}()

	fmt.Println("Start Test")

	err = mimicError("1")
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}

	panic("Mimic Panic") // Stop execution of all application code and propagate the panic upstream

	fmt.Println("End Test")
	return err
}
