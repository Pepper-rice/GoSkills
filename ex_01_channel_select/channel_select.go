package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个无缓冲的字符串通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 启动第一个协程，延迟1秒后向c1通道发送消息"one"
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	// 启动第二个协程，延迟3秒后向c2通道发送消息"two"
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	// 选择器 select 可以让你同时等待多个通道的操作，将协程、通道、选择器结合 循环两次，因为我们有两个通道需要接收数据
	for i := 0; i < 2; i++ {
		select {
		// 当c1通道有数据可读时，执行此case
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		// 当c2通道有数据可读时，执行此case
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 选择器 select 用途之一就是实现超时，下面等待 c1 通道结果，只等待 2 秒
	select {
	case res := <-c1:
		fmt.Println("received", res)
	// 如果在2秒内c1没有数据，则执行超时处理
	// time.After返回一个通道，该通道在指定时间后会发送当前时间
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 1")
	}
}
