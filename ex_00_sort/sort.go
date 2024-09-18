package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort 包实现内建以及自定义数据类型的排序
	// 注意：它的排序是原地排序的
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Sorted String:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Sorted Int", ints)

	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted?", sorted)
}
