package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)

	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, len(s))

	s = append(s, "d")
	fmt.Println(s, len(s))

	s = append(s, "e")
	fmt.Println(s, len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c, len(c))

	fmt.Println("3:5", s[3:5])
	fmt.Println(":3", s[:3])
	fmt.Println("2:", s[2:])

	s2 := []string{"x", "y", "z"}
	fmt.Println("s2", s2)

	twoD := make([][]int, 3)
	for row := 0; row < 3; row++ {
		innerLen := 4
		twoD[row] = make([]int, innerLen)
		for column := 0; column < innerLen; column++ {
			twoD[row][column] = row + column
		}
	}
	fmt.Println(twoD)
}
