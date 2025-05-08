package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 使用协程和通道，实现同样的互斥效果
// Go 共享内存思想是，通过通信使每个数据仅被单个协程所拥有，即通过通道实现共享内存
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 计数器，用于记录读写操作的次数
	var readOps uint64
	var writeOps uint64

	// 创建读写操作通道
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		// 创建一个map作为状态存储
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				// 将对应key的值通过响应通道发送回去
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动100个读操作协程
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 创建一个读操作，随机选择一个key
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				// 发送读请求到reads通道
				reads <- read
				// 等待读操作完成
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 启动100个写操作协程
	for w := 0; w < 100; w++ {
		go func() {
			for {
				// 创建一个写操作，随机选择key和value
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				// 发送写请求到writes通道
				writes <- write
				// 等待写操作完成
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	// 原子操作读取最终的读操作次数
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	// 原子操作读取最终的写操作次数
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
