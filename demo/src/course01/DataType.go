package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	//test01()
	// test02()
	// test03()
	// testSlice()
	//testSlice2()
	//testMap()
	//testChan()
	//testChan2()
	//generator := func(company string, department string, sn uint32) string {
	//	return appendSn(company + "-" + department + "-", sn)
	//}
	//fmt.Println(generateId(generator, "RD"))
}

func test01() {
	//var num1 int = 1
	//var num2 int = 2
	//var num3 int = 3
	
	//num1 := 1
	//num2 := 2
	//num3 := 3
	
	//var num1, num2, num3 int = 1, 2, 3 

	var (
		num1 int
		num2 int
		num3 int
	)
	num1, num2, num3 = 1, 2, 3
	fmt.Println(num1, num2, num3)
}

func test02() {
	var num uint64 = 65535
	size := unsafe.Sizeof(num)
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节\n", num, size)
}

func test03() {
	var num = 010 //八进制前缀0
	var num2 = 0x45
	fmt.Printf("10进制: %d  8进制:%o\n", num, num)
	fmt.Printf("10进制: %d  8进制:%o", num2, num2)
}

func testSlice() {
	array := [...]int{1, 2, 3, 4, 5}
	list := array[2:5]
	fmt.Println(cap(list))
}

// 类似java的List
func testSlice2() {
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := num[4:6:8]
	fmt.Println(len(slice), cap(slice))
	ints := slice[:cap(slice)]
	ints = append(ints, 11, 12, 13)
	fmt.Println(len(ints), cap(ints))
	ints2 := []int{0, 0, 0}
	copy(ints, ints2)
	fmt.Println(len(ints))
}

func testMap() {
	//字典类型的字面量如下: map[K]T  K:表示为键的类型, T:表示为值的类型  map[int]string
	mp := map[int]string{1: "Tom", 2: "Jack"}
	fmt.Println(mp)
	fmt.Println(mp[2])
	s, ok := mp[4]
	fmt.Println(s, ok)
	mp[3] = "xingtb"
	fmt.Println(mp[3])
}

// Channel 通道类型, 是Go语言中一种非常独特的数据结构。
// 它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。
// 相比之下，我们之前介绍的那些数据类型都不是并发安全的。这一点需要特别注意。
func testChan() {
	ch1 := make(chan string, 1)
	go func() {
		ch1 <- "已到达!"
	}()
	value := "数据" + <-ch1
	fmt.Println(value)
}

type Sender chan<- int
type Receiver <-chan int
func testChan2() {
	var myChan = make(chan int, 0)
	var num = 6
	go func() {
		var sender Sender = myChan
		sender <- num
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChan
		fmt.Println("Received!", <-receiver)
	}()
	time.Sleep(time.Second)
}

// 员工ID生成器
type EmployeeIdGenerator func(company string, department string, sn uint32) string
var company = "Gophers"
var sn uint32

func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	if nil == generator {
		return "", false
	}
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}


