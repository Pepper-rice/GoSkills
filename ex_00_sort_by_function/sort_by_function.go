package main

import (
	"fmt"
	"sort"
)

// 为了使用 Go 自定义函数排序，需要一个对应的类型，创建一个 []string 的别名，byLength
type byLength []string

// Len 实现 sort.Interface 接口的 Len 方法
func (s byLength) Len() int {
	return len(s)
}

// Swap 实现 sort.Interface 接口的 Swap 方法
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less 实现 sort.Interface 接口的 Less 方法 定义排序规则：按字符串长度升序排序
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// 创建一个字符串切片
	fruits := []string{"peach", "banana", "kiwi"}
	// 强制转换为 byLength 类型并使用 sort.Sort 进行排序
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
