package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 返回随机整数，0 <= n < 100
	fmt.Println(rand.Intn(1000))

	fmt.Println(rand.Intn(1000) + 100)

	// 默认情况下，种子是确定的。使用 NewSource 初始化种子
	// 相同的种子，会生成相同的随机数
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	fmt.Println(random.Intn(1000))
	fmt.Println(random.Intn(1000) + 100)

}
