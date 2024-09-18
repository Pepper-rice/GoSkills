package main

import "fmt"

func main() {

	// range 用于遍历，可以遍历数组，切片，map
	nums := []int{2, 3, 4}
	sum := 0

	// 第一个参数表示索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
