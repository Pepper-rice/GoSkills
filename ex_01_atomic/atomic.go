package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 计数器
	var ops uint64

	var wg sync.WaitGroup

	// 启动 50 个协程，每个协程增加 1000 次
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 可以安全的访问 ops
	fmt.Println("ops:", ops)

	// 更安全的获取 ops
	fmt.Println("atomic ops:", atomic.LoadUint64(&ops))
}
