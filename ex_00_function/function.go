package main

import "fmt"

// 函数，参数与返回值
func plus(a int, b int) int {
	return a + b

}

// Go 支持多个返回值
func vals() (int, int) {
	return 3, 7
}

// 可变参数函数与 Java 类似
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	res := plus(2, 3)
	fmt.Println(res)

	a, b := vals()
	_, c := vals()

	fmt.Println(a, b)
	fmt.Println(c)

	// 常规调用可变参数
	total1 := sum(1, 2, 3)
	fmt.Println(total1)

	// 将多个值的数组传入可变参数函数
	nums := []int{1, 2, 3}
	total2 := sum(nums...)
	fmt.Println(total2)
}
