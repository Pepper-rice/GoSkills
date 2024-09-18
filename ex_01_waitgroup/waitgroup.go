package main

import (
	"fmt"
	"sync"
	"time"
)

// 我们改写 worker_pool 中的工作协程
// WaitGroup 必须通过指针进行传递
func worker(id int, wg *sync.WaitGroup) {

	// 当函数 return 时，通知 WaitGroup 工作已经完成
	defer wg.Done()

	fmt.Println("worker starting", id)

	time.Sleep(time.Second)

	fmt.Println("worker done", id)
}

func main() {
	// WaitGruop 用户等待该函数启动的所有协程
	var wg sync.WaitGroup

	// 启动几个协程，并递增 WaitGroup 的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞，直到 WaitGroup 计数器恢复为 0
	wg.Wait()
}
