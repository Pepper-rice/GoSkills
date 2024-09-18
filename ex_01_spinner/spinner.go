package main

import (
	"fmt"
	"time"
)

func main() {
	// 在一个单独的 goroutine 显示一个小动画
	go spinner(100 * time.Millisecond)

	const n = 45

	fibN := fib(n)

	// 主函数返回程序退出，所有 goroutine 都会中止
	fmt.Printf("\nFib(%d) = %d", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
