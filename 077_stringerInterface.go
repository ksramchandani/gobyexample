package main

import (
	"fmt"
)

// Human doc
type Human struct {
	name  string
	age   int
	phone string
}

// For String interface, all you have to do is implement the following method
// func (r Receiver) String() string {
// }

func (h Human) String() string {
	return fmt.Sprintf("name of human = %s with age %d and phone %s\n", h.name, h.age, h.phone)
}

func main() {
	h := Human{"abc", 10, "xxx-yyy-zzzz"}
	fmt.Println(h)
}
