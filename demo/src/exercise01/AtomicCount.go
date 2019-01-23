package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//atomicCount()
	testLock()
}

// 原子计数
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

// 互斥体实例, 类似于java中的ReentrantLock类
func testLock() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0
	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for true {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for true {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
