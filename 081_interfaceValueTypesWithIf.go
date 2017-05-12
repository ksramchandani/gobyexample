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

	fmt.Println(l)

	for _, elem := range l {
		if value, ok := elem.(int); ok {
			fmt.Printf("elem is of type int and its value is %d. ok = %t\n", value, ok)
		} else if value, ok := elem.(float64); ok {
			fmt.Printf("elem is of type float64 and its value is %f. ok = %t\n", value, ok)
		} else if value, ok := elem.(string); ok {
			fmt.Printf("elem is of type string and its value is %q. ok = %t\n", value, ok)
		}
	}
}
