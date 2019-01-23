package main

import (
	"fmt"
	"sort"
)

/**
排序练习
*/
func main() {
	//testSortApi()
	udfSort()
}

// 基本的排序api
func testSortApi() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{6, 4, 7, 5}
	fmt.Println(sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))
}

// 自定义排序: 通过字符串的长度进行排序
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func udfSort() {
	fruits := ByLength{"peach", "banana", "kiwi"}
	sort.Sort(fruits)
	fmt.Println(fruits)
}
