package main

import (
	"fmt"
)

func main() {

	// 示例1：基本的 switch 用法
	// switch 语句用于基于不同条件执行不同代码块
	// 这里我们将变量 i 的值与 case 进行匹配
	i := 2
	fmt.Println("示例1：基本的 switch 用法")
	switch i {
	case 1:
		fmt.Println("结果：i 等于 1")
	case 2:
		fmt.Println("结果：i 等于 2")
	case 3:
		fmt.Println("结果：i 等于 3")
	default:
		fmt.Println("结果：i 不在 1 到 3 的范围内")
	}

	fmt.Println("\n示例2：类型开关")
	// 示例2：类型开关（type switch）
	// 类型开关用于比较接口变量的类型，而不是它的值
	// 这个示例定义了一个函数，打印出传入变量的类型
	whatAmI := func(i any) {
		switch t := i.(type) { // 使用.(type) 进行类型断言，判断传入变量的实际类型
		case bool:
			fmt.Println("输入是布尔型：", t)
		case int:
			fmt.Println("输入是整数型：", t)
		case string:
			fmt.Println("输入是字符串型：", t)
		default:
			fmt.Println("输入是其他类型，不确定类型是：", t)
		}
	}

	// 调用 whatAmI 函数，传入不同类型的参数
	whatAmI(true)  // 输出布尔类型
	whatAmI(1)     // 输出整数类型
	whatAmI("abc") // 输出字符串类型
	whatAmI(3.14)  // 输出未知类型（浮点型，未定义在 switch 的 case 中）

}
