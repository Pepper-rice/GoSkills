package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器和定时器类似，使用一个通道来发送数据
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Second * 10)

	// 停止打点器
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
