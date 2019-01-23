package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
一些集合函数的练习, 类型于java或者python中的高阶函数(map, reduce, filter)
*/
func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

	fmt.Println(Reduce(strs, func(s string, s2 string) string {
		buffer := bytes.Buffer{}
		buffer.WriteString(s)
		buffer.WriteString(s2)
		return buffer.String()
	}))
}

// 返回某个元素的下标
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 包含某个元素
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 集合中任意一个满足即可
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 集合中所有元素满足条件
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return false
		}
	}
	return true
}

// 过滤函数, 只保留满足函数的元素
func Filter(vs []string, f func(string) bool) []string {
	s := make([]string, len(vs))
	for _, v := range vs {
		if f(v) {
			s = append(s, v)
		}
	}
	return s
}

// map 操作
func Map(vs []string, f func(string) string) []string {
	for i, v := range vs {
		vs[i] = f(v)
	}
	return vs
}

// reduce操作
func Reduce(vs []string, f func(string, string) string) string {
	var s string
	for i := 0; i < len(vs); i++ {
		s = f(s, vs[i])
	}
	return s
}
