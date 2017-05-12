package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{1, 10, 5, 3, 2}
	f := []float64{3.1, 2.2, 5.5, 4.3}
	s := []string{"def", "xyz", "abc"}

	sort.Ints(i)
	fmt.Println(i)

	sort.Float64s(f)
	fmt.Println(f)

	sort.Strings(s)
	fmt.Println(s)
}
