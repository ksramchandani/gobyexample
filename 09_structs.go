package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	p1 := person{"abc", 1}
	fmt.Println("p1", p1)

	p2 := person{name: "def", age: 10}
	fmt.Println("p2", p2)

	p3 := &person{name: "def", age: 10}
	fmt.Println("p3", p3)
	fmt.Println("*p3", *p3)

	sp1 := &p1
	fmt.Println("sp1", sp1)
	fmt.Println("sp1.name", sp1.name)

}
