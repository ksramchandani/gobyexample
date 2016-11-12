package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	i := 5
	s := "Hello world"

	a = i
	fmt.Println(a, reflect.TypeOf(a))

	a = s
	fmt.Println(a, reflect.TypeOf(a))

}
