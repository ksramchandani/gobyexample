package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func addWorker(opsPtr *uint64) {
	atomic.AddUint64(opsPtr, 1)

	runtime.Gosched()
}

func main() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go addWorker(&ops)
	}

	time.Sleep(1 * time.Microsecond)
	val := atomic.LoadUint64(&ops)

	fmt.Println(val)

}
