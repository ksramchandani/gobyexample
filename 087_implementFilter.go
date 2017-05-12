package main

import (
	"fmt"
)

type testInt func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func filter(s []int, f testInt) []int {
	var result []int
	for _, value := range s {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice = ", s)
	odd := filter(s, isOdd)
	fmt.Println("odd slice = ", odd)
	even := filter(s, isEven)
	fmt.Println("even slice = ", even)
}
