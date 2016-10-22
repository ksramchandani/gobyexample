package main

import (
	"fmt"
)

func main() {
	rows := 3
	columns := 4

	twoD := make([][]int, rows)
	for i := 0; i < rows; i++ {
		twoD[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
