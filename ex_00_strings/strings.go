package main

import (
	"fmt"
	"reflect"
	"strconv" // 类型转换，字符串转数值
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello，中国"
	strs := []string{"a", "b", "c"}
	strInt := "123456"
	ints := 123456

	// 字符串比较
	if str == "Hello，中国" {
		fmt.Println("str == \"Hello，中国\" 是否相等?", "equal")
	} else {
		fmt.Println("str == \"Hello，中国\" 是否相等？", "not equal")
	}

	// 字符串操作示例
	isContains := strings.Contains("test", "es") // 检查"es"是否是"test"的子串
	splits := strings.Split("a,b, c", ",")       // 根据逗号分割字符串
	join := strings.Join(strs, "-")              // 用"-"连接字符串切片
	trimSpace := strings.TrimSpace("    		abc ") // 去除字符串两端的空白字符

	// 字符长度
	len := len(str)
	lenByte := utf8.RuneCountInString(str)

	// 字符串与数值转换
	atoi, _ := strconv.Atoi(strInt)                 // 将字符串转为整数
	parseInt, _ := strconv.ParseInt(strInt, 10, 64) // 将字符串转为64位整数
	formatInt := strconv.FormatInt(int64(ints), 10) // 将整数转为字符串

	// 输出结果和类型信息，帮助理解
	fmt.Println("strs的类型是：", reflect.TypeOf(strs))
	fmt.Println("是否包含子串：", isContains)
	fmt.Println("按逗号分割：", splits)
	fmt.Println("连接字符串：", join)
	fmt.Println("字符串的字节长度：", len)
	fmt.Println("字符串的字符长度：", lenByte)
	fmt.Println("去除空白字符：", trimSpace)
	fmt.Println("atoi转换结果类型：", reflect.TypeOf(atoi))
	fmt.Println("atoi转换结果：", atoi)
	fmt.Println("parseInt转换结果类型：", reflect.TypeOf(parseInt))
	fmt.Println("parseInt转换结果：", parseInt)
	fmt.Println("formatInt转换结果类型：", reflect.TypeOf(formatInt))
	fmt.Println("formatInt转换结果：", formatInt)

	// 使用反引号定义多行字符串，避免转义问题
	strBytes := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Println("strBytes的类型是：", reflect.TypeOf(strBytes))
	fmt.Println("strBytes的内容是：", strBytes)
}
