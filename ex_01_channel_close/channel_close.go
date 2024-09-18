package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 模拟工作协程，使用 j, ok := <-jobs 接收数据，根据第二个值 ok， 如果 jobs 已经关闭，该值将是 false

	go func() {
		for {
			j, ok := <-jobs
			if ok {
				fmt.Println("received job", j)
			} else {
				// done 通道用于通知主线程
				done <- true

				// 这个打印应该放在后面，否则可能会打印多次，因为程序运行太快了
				fmt.Println("received all jobs")
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	fmt.Println("send all jobs")

	// 关闭一个通道，意味着不能再向这个通道发送值，该特性可以向通道的接收方传达已经完成的信息
	close(jobs)

	// done 通道会阻塞到直到工作线程执行完毕，
	// 如果没有这个，程序会一直收到 received all jobs
	<-done
}
