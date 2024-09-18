package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// Go 支持指针，使用方法与 C 语言类似
func zeroptr(iptr *int) {
	// 表示解引用
	*iptr = 0

}

func main() {
	i := 1
	fmt.Println("init:", i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	// 打印变量 i 的指针地址
	fmt.Println(&i)
}
