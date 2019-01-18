package main

import (
	"fmt"
)

func main() {
	//testGo()
	// testSelect2()
	//time.Sleep(10 * time.Millisecond)
	testSelect3()
}

func testGo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		<-ch1
		fmt.Println("1")
		ch2 <- 1
	}()
	go func() {
		<-ch2
		<-ch3
		fmt.Println("2")
	}()
	go func() {
		ch1 <- 1
		ch3 <- 1
		fmt.Println("3")
	}()
}

func testSelect2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func(p chan int) {
		for i := 0; i < 20; i++ {
			ch1 <- 1
		}
	}(ch1)

	go func(p chan int) {
		for i := 0; i < 20; i++ {
			ch2 <- 2
		}
	}(ch2)

	go func(p chan int) {
		for i := 0; i < 20; i++ {
			ch3 <- 3
		}
	}(ch3)

	go func(p ...chan int) {
		for i := 0; i < 20; i++ {
			ch4 <- 4
		}
	}(ch4)

	go func(p ...chan int) {
		for i := 0; i < 20; i++ {
			select {
			case e, ok := <-ch1:
				fmt.Printf("%d  %s\n", e, ok)
			case e, ok := <-ch2:
				fmt.Printf("%d  %s\n", e, ok)
			case e, ok := <-ch3:
				fmt.Printf("%d  %s\n", e, ok)
			case e, ok := <-ch4:
				fmt.Printf("%d  %s\n", e, ok)
				//default:
				//	fmt.Println("No Data!")
			}
		}
	}(ch1, ch2, ch3, ch4)
}

func testSelect3() {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)
	ch3 := make(chan int, 20)
	ch4 := make(chan int, 20)

	for i := 0; i < 20; i++ {
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
		ch4 <- 4
	}

	for i := 0; i < 20; i++ {
		select {
		case e, ok := <-ch1:
			fmt.Printf("step%d: %d  %s\n", i, e, ok)
		case e, ok := <-ch2:
			fmt.Printf("step%d: %d  %s\n", i, e, ok)
		case e, ok := <-ch3:
			fmt.Printf("step%d: %d  %s\n", i, e, ok)
		case e, ok := <-ch4:
			fmt.Printf("step%d: %d  %s\n", i, e, ok)
		default:
			fmt.Println("No Data!")
		}
	}
}
