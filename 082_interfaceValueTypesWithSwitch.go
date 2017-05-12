package main

import (
	"fmt"
)

type Element interface{}
type List []Element

func main() {
	l := make(List, 3)
	l[0] = 1
	l[1] = 2.4
	l[2] = "hello"

	for _, elem := range l {
		switch elem.(type) {
		case int:
			fmt.Println("int", elem)
		case float64:
			fmt.Println("float64", elem)
		case string:
			fmt.Println("string", elem)
		}
	}
}
