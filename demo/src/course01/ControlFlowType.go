package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//testIf()
	//testSwitch()
	//testFor()
	//testFor2()
	//testSelect()
	//testDefer()
	//testDefer2()
	//testDefer3()
	testDefer4()
}

func testIf() {
	var num = 5
	if num += 4; 10 > num {
		var num int
		num += 3
		fmt.Print(num)
	} else {
		num -= 2
		fmt.Print(num)
	}
	fmt.Println(num)
}

func testSwitch() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	i := rand.Intn(4) % 2
	switch v := ia[i ]; interface{}(v).(type) {
	case rune:
		fmt.Println("Case A.")
	case byte:
		fmt.Println("Case B.")
	default:
		fmt.Println("Unknown!")
	}
}

// 关键字range 跟在可迭代的集合前面
func testFor() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if (i == 3 || i == 4) && j == 2 {
				fmt.Printf("%d*%d=%d  ", i, j, i*j)
			} else {
				fmt.Printf("%d*%d=%d ", i, j, i*j)
			}
		}
		fmt.Println()
	}
}

func testFor2() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for key, value := range map1 {
		fmt.Printf("key:%d  value:%s \n", key, value)
	}
}

func testSelect() {
	ch4 := make(chan int, 1)
	for i := 1; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)
			close(ch4)
		default:
			fmt.Println("No Data!")
			ch4 <- 1
		}
	}
}

func testDefer() {
	f := func(i int) int {
		fmt.Printf("%d ", i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}

func testDefer2() {
	for i := 1; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func testDefer3() {
	for i := 1; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func testDefer4() {
	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Printf("%d ", n)
		}(func() int {
			n := fibonacci(i)
			fmt.Printf("%d ", n)
			return n
		}())
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
