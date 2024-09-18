package main

import (
	"flag" // 解析命令行参数
	"fmt"
	"os"
	"strings"
)

func main() {
	// 假设程序以 exec a b c d 的方式运行
	args := os.Args

	// 第一个参数是该程序的路径
	fmt.Println(args)
	// [1:] 是全部的参数，即["a", "b", "c", "d"]
	fmt.Println(args[1:])

	// 命令行参数声明，分别是参数名称，默认值，说明。 -h 可以显示
	// 命令行参数仅支持字符串、整型、布尔型
	// exec -word=test -num=100 -bool=true
	wordParam := flag.String("word", "foo", "a string")
	intParam := flag.Int("num", 42, "an int")
	boolParam := flag.Bool("bool", false, "a bool")

	flag.Parse()

	fmt.Println("word:", *wordParam)
	fmt.Println("num:", *intParam)
	fmt.Println("bool:", *boolParam)

	// 获取所有的环境变量
	environs := os.Environ()
	for _, env := range environs {
		pair := strings.SplitN(env, "=", 2)
		fmt.Println(pair)
	}

	// 获取单个的环境变量
	fmt.Println(os.Getenv("JAVA_HOME"))
}
