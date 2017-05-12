// Write a Program using Interfaces, that will return the Max of a integer, float or struct (containing age) slice

// First start with getting max of a int slice and a float slice. Then adding a struct slice is easy
// as you only have to implement the relevant methods

package main

import (
	"fmt"
	"reflect"
)

type MaxData interface {
	Len() int
	Get(index int) interface{}
	Bigger(index1, index2 int) int
}

type Person struct {
	name string
	age  int
}

type IntegerSlice []int
type Float64Slice []float64
type People []Person

func (i IntegerSlice) Len() int {
	return len(i)
}

func (f Float64Slice) Len() int {
	return len(f)
}

func (p People) Len() int {
	return len(p)
}

func (i IntegerSlice) Get(index int) interface{} {
	return i[index]
}

func (f Float64Slice) Get(index int) interface{} {
	return f[index]
}

func (p People) Get(index int) interface{} {
	return p[index]
}

func (i IntegerSlice) Bigger(index1, index2 int) int {
	if i[index1] > i[index2] {
		return index1
	}
	return index2
}

func (f Float64Slice) Bigger(index1, index2 int) int {
	if f[index1] > f[index2] {
		return index1
	}
	return index2
}

func (p People) Bigger(index1, index2 int) int {
	if p[index1].age > p[index2].age {
		return index1
	}
	return index2
}

func Max(data MaxData) interface{} {
	fmt.Println(reflect.TypeOf(data))

	// Set max_value to the first element
	maxIndex := 0
	for iCount := 0; iCount < data.Len(); iCount++ {
		maxIndex = data.Bigger(maxIndex, iCount)
	}
	return data.Get(maxIndex)
}

func main() {
	i := IntegerSlice{1, 10, 2, 4, 8}
	f := Float64Slice{3.4, 5.2, 8.9, 2.3}
	p := People{
		Person{"abc", 15},
		Person{"def", 20},
		Person{"xyz", 10},
	}

	fmt.Println(Max(i))
	fmt.Println(Max(f))
	fmt.Println(Max(p))

}
