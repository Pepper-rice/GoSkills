package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// Go 的 range 循环可直接在 channel 上迭代，如果 channel 被关闭并且没有值可以接受时跳出循环
		for j := range jobs {
			fmt.Println("received job", j)
		}
		done <- true

		fmt.Println("received all jobs")
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	fmt.Println("send all jobs")

	close(jobs)

	<-done
}
