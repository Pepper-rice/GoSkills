package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// 协程是轻量级线程
	go f("goroutine")

	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Println(msg)
		}
	}("going")

	// 主程序要停留一段时间否则看不到
	time.Sleep(time.Second * 1)

	fmt.Println("done")
}
