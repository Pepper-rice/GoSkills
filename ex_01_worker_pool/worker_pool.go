package main

import (
	"fmt"
	"time"
)

// 定义 worker 程序，模拟执行 1 秒
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)

		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)

		results <- j * 2
	}

}

func main() {
	const jobNum = 10

	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)

	// 模拟启动三个 worker， 刚开始是阻塞的，因为没收到任何任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送五个任务，然后关闭通道
	for j := 1; j <= jobNum; j++ {
		jobs <- j
	}
	close(jobs)

	// 获取五个任务的执行结果，确保所有的 worker 协程都已经完成
	// 这是通过额外的 results 通道实现的，另一个方法是使用 WaitGroup
	for a := 1; a <= jobNum; a++ {
		fmt.Println("result:", <-results)
	}
}
