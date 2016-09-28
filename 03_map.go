package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 20

	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}

	fmt.Println(m["k1"])
	fmt.Println(len(m))

	fmt.Println("m", m)
	delete(m, "k1")
	fmt.Println("m", m)

	// Safe get a key from a map
	value, ok := m["k1"]
	fmt.Println(ok)
	if ok {
		fmt.Println(value)
	}

	value, ok = m["k2"]
	fmt.Println(ok)
	if ok {
		fmt.Println(value)
	}

	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)

	aVal, ok := n["a"]
	if ok {
		fmt.Println(aVal)
	}

}
