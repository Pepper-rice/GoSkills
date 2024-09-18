package main

import (
	"fmt"
	"reflect"
)

// intSeq 函数返回一个在其函数体内定义的匿名函数。返回的函数使用闭包的方式隐藏变量 i，
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 返回值是一个函数。这个函数的值包含了自己的值 i, 每次调用都会更新 i 的值
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(reflect.TypeOf(nextInt))
}
