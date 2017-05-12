package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type People []Person

// All you have to do is implement the following methods
// func (receiver) Len() int {}
// func (receiver) Less(i, j int) bool {}
// func (receiver) Swap(i, j int) {}

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].age < p[j].age {
		return true
	}
	return false
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p := People{
		Person{"abc", 10},
		Person{"xyz", 5},
		Person{"tuv", 7},
	}

	fmt.Printf("Original %v\n", p)
	sort.Sort(p)
	fmt.Printf("Sorted %v\n", p)

}
