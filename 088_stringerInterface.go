package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "Name: " + h.name + ", Age: " + strconv.Itoa(h.age) + ", Phone: " + h.phone
}

func main() {
	h := Human{"abc", 18, "111-111-1111"}
	fmt.Println(h)

}
