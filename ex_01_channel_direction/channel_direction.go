package main

import "fmt"

// 使用通道作为函数参数，可以指定通道是否为只读或只写，该函数定义了一个只能发送数据的通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 该函数定义了两个通道，其中 pings 只用于接收数据，pongs 只用于发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")

	pong(pings, pongs)

	fmt.Println(<-pongs)
}
