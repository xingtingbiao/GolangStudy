package main

import "fmt"

func main() {
	//sum(1, 2, 5, 7)
	// 数组或者切片类型直接作为可变参数需要加...
	// sum([]int{1, 3, 5, 7, 9}...)
	//testIntSeq()
	//fmt.Println(fact(7))
}

// 可变参数的函数
func sum(nubs ...int) {
	total := 0
	for _, n := range nubs {
		total += n
	}
	fmt.Println(total)
}

// 闭包(匿名函数) 闭包就是能够读取其他函数内部变量的函数。
// 例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testIntSeq() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
	fmt.Println(nextInts())
}

// 阶乘的递归练习
func fact(n int) int {
	if n == 2 {
		return 2
	}
	return n * fact(n-1)
}
