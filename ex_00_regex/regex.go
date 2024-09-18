package main

import (
	"fmt"
	"regexp" // 正则表达式操作
)

func main() {
	// 测试字符串是否符合正则表达式(匹配任何以“p”开头，接着一个或多个小写字母，再接着“ch”结尾的字符串)
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 通过 Compile 得到一个优化过的结构体(将表达式p([a-z]+)ch编译并存储在变量compile中)
	compile, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(compile.MatchString("peach"))

	// 查找匹配到的字符串(精确匹配第一个符合表达式的子字符串)
	fmt.Println(compile.FindString("aaa   peach bbb"))

	// 返回完全匹配和局部匹配的字符串：peach 和 ea(返回切片的第一个元素是完全匹配的字符串，后续元素是表达式中的捕获组（子表达式）匹配的内容)
	fmt.Println(compile.FindStringSubmatch("aaa   peach pinch"))

	// 返回全部的匹配项
	fmt.Println(compile.FindAllString("peach punch pinch", -1))
}
