package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)

		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 3)

		c2 <- "two"
	}()

	// 选择器 select 可以让你同时等待多个通道的操作，将协程、通道、选择器结合
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 选择器 select 用途之一就是实现超时，下面等待 c1 通道结果，只等待 2 秒
	select {
	case res := <-c1:
		fmt.Println("received", res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 1")
	}
}
