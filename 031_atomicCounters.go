package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i uint64 = 1

	// Add in a atomic fashion
	atomic.AddUint64(&i, 1)

	// Fetch in a atomic fashion
	val := atomic.LoadUint64(&i)
	fmt.Println(val)
}
