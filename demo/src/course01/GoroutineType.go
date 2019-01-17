package main

import (
	"fmt"
	"time"
)

func main() {
	testGo()
	time.Sleep(10 * time.Millisecond)
}

func testGo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		<- ch1
		fmt.Println("1")
		ch2 <- 1
	}()
	go func() {
		<- ch2
		<- ch3
		fmt.Println("2")
	}()
	go func() {
		ch1 <- 1
		ch3 <- 1
		fmt.Println("3")
	}()
}
