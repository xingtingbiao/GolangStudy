package main

import (
	"errors"
	"fmt"
)

func main() {
	// 注意这里defer...recover()写在了可能发生异常的前面, 类似提前设置捕获手段
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()
	fmt.Println("Enter main")
	outerFunc()
	fmt.Println("Quit main")
}

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}
