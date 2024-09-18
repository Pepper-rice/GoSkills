package main

import (
	"fmt"
	"time"
)

// 利用通道的阻塞特性，可以实现协程通知功能
// done 通道将被用于通知其他协程这个函数的工作已经完成
func worker(done chan bool) {
	fmt.Println("working")

	time.Sleep(time.Second)

	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)

	// 程序将阻塞在这里，直到收到 worker 使用通道发送的通知
	go worker(done)

	// 如果这行代码移除，程序可能在 worker 运行前就结束了
	<-done
}
