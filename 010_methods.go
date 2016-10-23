package main

import (
	"fmt"
)

type rect struct {
	height int
	width  int
}

func (rPtr *rect) double() {
	rPtr.height = rPtr.height * 2
	rPtr.width = rPtr.width * 2
}

func (rPtr *rect) area() int {
	return rPtr.height * rPtr.width
}

func main() {

	r := rect{10, 20}
	fmt.Println(r)

	rPtr := &r
	rPtr.double()
	fmt.Println(r)

	fmt.Println(rPtr.area())

}
