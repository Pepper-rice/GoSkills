package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WithCancel()

	//WithDeadline()

	WithValue()
}

func WithCancel() {
	// WithCancel 创建一个可以取消的 context，可以在 goroutine 中使用该 context 来控制 goroutine 的生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine canceled")
				return
			default:
				fmt.Println("goroutine running")
			}

			time.Sleep(time.Second * 1)
		}
	}()

	// 模拟操作，等待两秒
	time.Sleep(time.Second * 3)

	// 通知子协程退出，调用 cancel 函数时，goroutine会收到取消通知并结束执行
	cancel()

	// 延迟退出主协程，为了看到通知后的信息，否则可能来不及打印
	time.Sleep(time.Millisecond * 100)
}

func WithDeadline() {
	// 这个例子与上面的例子类似，只是使用 WithDeadline 创建了一个带有截止时间的 context，不用手动的通知 cancel()
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine deadline canceled")
				return
			default:
				fmt.Println("goroutine running")
			}

			time.Sleep(time.Second * 1)
		}
	}()

	// 模拟一些操作，比截止时间更长，为了看到通知后的信息，否则可能来不及打印
	time.Sleep(5 * time.Second)
}

func WithValue() {
	// WithValue 创建一个带有 key-value 的 context，可以通过 Value 函数获取 value
	key := "name"
	ctx := context.WithValue(context.Background(), "name", "gopher")

	go func() {
		if v := ctx.Value(key); v != nil {
			fmt.Println("Value:", v)
		}
	}()

	// 模拟一些操作，比截止时间更长，为了看到通知后的信息，否则可能来不及打印
	time.Sleep(1 * time.Second)
}
