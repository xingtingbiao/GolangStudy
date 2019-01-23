package main

import (
	"fmt"
	"time"
)

/**
工作池实例, 类似java的线程池
*/
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w < 4; w++ {
		go worker1(w, jobs, results)
	}

	for j := 1; j < 6; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 1; i < 6; i++ {
		<-results
	}
}

func worker1(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, j)
		result <- j * 2
	}
}
