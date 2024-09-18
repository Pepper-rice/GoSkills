package main

import (
	"fmt"     // 用于格式化输入输出，全称format
	"reflect" // 用于在运行时通过反射检查类型
)

func main() {
	// 声明 5 个元素的数组，默认会初始化 0 值
	var a [5]int
	fmt.Println(a)

	// 将数组的第 5 个元素赋值为 5
	a[4] = 5
	fmt.Println(a[4])

	// 内置函数 len() 获取数组长度
	fmt.Println(len(a))

	// 声明并初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 反射获取类型
	fmt.Println(reflect.TypeOf(a))

	// Go切片常见API
	s := make([]int, 5, 10) // 创建一个切片，长度为5，容量为10
	src := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s切片：", s)
	fmt.Println("s切片的容量：", cap(s))
	fmt.Println("向s切片追加元素，返回新的切片：", append(s, 1))
	fmt.Println("从数组或切片中获取 子切片：", s[4:6])
	fmt.Println("复制切片：", copy(s, src)) // 如果目标切片比源切片小，copy 会截断源切片。返回复制的元素个数。
	fmt.Println("任意类型转字符串：", fmt.Sprint(s))

}
