package main

import "fmt"

// const 用于声明一个常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// const 可以出现在任何 var 出现的地方
	const n = 5000000

	// const 可以执行运算
	const d = 3e20 / n
	fmt.Println(d)
}
