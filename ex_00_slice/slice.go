package main

import (
	"fmt"
)

func main() {
	// 切片，实际可以理解为动态数组，需要使用内置函数 make() 来创建，默认值依然为 0 值
	i := make([]int, 3)
	fmt.Println(i)

	// 切片和数组是不同的类型，但是打印的结果是类似的
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	// len() 获取切片长度
	fmt.Println("len:", len(s))

	// append() 返回一个新的切片
	ss := append(s, "d", "f")
	fmt.Println(ss)

	// 切片的复制
	c := make([]string, len(ss))
	copy(c, s)
	fmt.Println(c)

	// 切片支持 [low:high] 进行切片操作，类似 Python
	fmt.Println(ss[2:5])

	// 切片也可以组成多维，且长度可以不一致
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
