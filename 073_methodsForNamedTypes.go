package main

import (
	"fmt"
)

type sliceIntegers []int

type peopleAge map[string]int

func (s sliceIntegers) sum() {
	var total int
	for _, num := range s {
		total = total + num
	}
	fmt.Println("sum of integers passed =", total)
}

func (p peopleAge) olderThanEight() {
	fmt.Println("People olderThanEight")
	for k, v := range p {
		if v > 8 {
			fmt.Println(k, v)
		}
	}
}

func main() {
	s := sliceIntegers{1, 2, 3, 4}
	s.sum()

	p := peopleAge{
		"abc": 10,
		"def": 20,
		"ghi": 5,
		"xyz": 7,
	}

	p.olderThanEight()
}
