package main

import "fmt"

func main() {
	// 默认情况下，通道是无缓冲的，只有对应的接收 <- chan 准备好了时候，才允许发送
	// 而有缓冲通道，允许在没有对应接收者的情况下，缓存一定数量的值
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "out"
	fmt.Println(<-messages)
}
