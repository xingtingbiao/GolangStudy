package main

import (
	"fmt"
	"os"
)

/*
命令行参数练习
*/
func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	// 第一个参数值是程序的路径
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]
	arg := os.Args[3]

	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}
