package main

import (
	"fmt"
)

type geometry interface {
	area() int
}

type rect struct {
	height int
	width  int
}

type circle struct {
	radius int
}

func (r rect) area() int {
	return r.height * r.width
}

func (c circle) area() int {
	return c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println()
}

func main() {
	r := rect{10, 20}
	c := circle{10}

	measure(r)
	measure(c)

}
