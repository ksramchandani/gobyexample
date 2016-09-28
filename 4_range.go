package main

import (
	"fmt"
)

func main() {
	i := []int{10, 20, 30}
	for k, v := range i {
		fmt.Println(k, v)
	}

	for _, v := range i {
		fmt.Println(v)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	s := "golang"
	for _, v := range s {
		fmt.Printf("%s ", string(v))
	}

}
