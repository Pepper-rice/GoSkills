package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器表示未来某时刻触发的事件
	timer1 := time.NewTimer(time.Second * 2)

	// timer1.C 会一直阻塞，直到定时器通道 C 明确发送了实效的值
	fail := <-timer1.C
	fmt.Println(fail)

	// 如果只是单纯的等待，time.Sleep 就够了，定时器的作用是可以在触发之前取消它
	// 使用 Stop() 方法
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

}
