package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/**
原子计数
*/
func main() {
	atomicCount()
}

func atomicCount() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for true {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("ops:", atomic.LoadUint64(&ops))
}
