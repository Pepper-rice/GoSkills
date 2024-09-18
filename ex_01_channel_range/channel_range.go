package main

import "fmt"

func main() {
	queue := make(chan string, 3)

	queue <- "one"
	queue <- "two"
	queue <- "three"

	close(queue)

	// 所以，一个非空的通道是可以关闭的，并且，通道中的值仍然可以被接收到
	// range 可以遍历通道取值，并自动处理关闭与零值
	//for q := range queue {
	//	fmt.Println(q)
	//}

	// 如果 channel 中已经没有数据的话将产生一个零值的数据。如果你是用 for 循环而不是迭代的话，需要接收的时候判断通道是否已经关闭，否则是死循环
	i := 0
	for {
		i++
		q := <-queue
		fmt.Println(fmt.Sprintf("%d, %s\n", i, q))
	}
}
