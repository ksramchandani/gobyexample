package main

import (
	"fmt"
)

func noPointerIncrement(i int) {
	i += 1
}

func pointerIncrement(iPtr *int) {
	*iPtr += 1
}

func main() {
	x := 0
	fmt.Println("Initial: ", x)

	noPointerIncrement(x)
	fmt.Println("After noPointerIncrement: ", x)

	pointerIncrement(&x)
	fmt.Println("After pointerIncrement: ", x)

}
