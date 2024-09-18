package main

import "fmt"

func main() {
	// 通道是连接多个协程的管道，可以从一个协程将值发送到通道，在另一个协程中接收
	message := make(chan string)

	// 无缓存的通道，也叫同步通道。默认发送和接收操作是"阻塞"的，直到发送方和接收方都就绪
	go func() {
		message <- "ping"
	}()

	msg := <-message

	fmt.Println(msg)
}
