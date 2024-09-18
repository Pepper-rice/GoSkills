package main

import "fmt"

func main() {
	i := 1

	// Go 中只有 for 循环，如果想实现其他语言中的 while 循环，那么可以给单个循环条件
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带任何条件的死循环，通过 break 或 return 跳出，continue 表示直接进入下一次循环
	for {
		fmt.Println("loop")
		break
	}
}
