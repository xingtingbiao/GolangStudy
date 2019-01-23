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
	//testLock()
	testLock2()
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

// 互斥的另一种实现方式, 通过chan的阻塞来实现
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func testLock2() {
	var readOps uint64 = 0
	var writeOps uint64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	state := make(map[int]int)
	go func(state map[int]int) {
		for true {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}(state)

	for r := 0; r < 100; r++ {
		go func() {
			for true {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for true {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))
	fmt.Println("state:", state)
}
