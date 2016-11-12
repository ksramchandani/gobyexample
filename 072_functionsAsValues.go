// Write a program that will be given a slice as an input and will return the odd and even number slices

package main

import (
	"fmt"
)

type determineType func(int) bool

func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func filter(s []int, f determineType) []int {
	var result []int
	for _, num := range s {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)
	odd := filter(s, isOdd)
	fmt.Println("odd", odd)
	even := filter(s, isEven)
	fmt.Println("even", even)
}
