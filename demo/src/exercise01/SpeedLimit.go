package main

import (
	"fmt"
	"time"
)

/**
速率限制实例
*/
func main() {
	speedLimit()
}

func speedLimit() {
	// 连续200间隔限制
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Second)
	for req := range requests {
		<-limiter
		fmt.Printf("request %d %s\n", req, time.Now())
	}

	time.Sleep(time.Second * 2)
	// 新增一个限制, 允许前三个请求可以立即执行
	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i < 4; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Printf("request %d %s\n", req, time.Now())
	}
}
