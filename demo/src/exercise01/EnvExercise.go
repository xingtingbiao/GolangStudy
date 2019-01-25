package main

import (
	"fmt"
	"os"
)

/*
系统环境变量的练习
*/
func main() {
	testEnv()
}

// 查询系统环境变量
func testEnv() {
	_ = os.Setenv("FOO", "1")
	fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("BAR"))

	fmt.Println()
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
