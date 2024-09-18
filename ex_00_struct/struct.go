package main

import (
	"encoding/json" // 用于JSON编码和解码
	"fmt"           // 用于格式化输出
)

// 定义一个基本的结构体 `Person`，用于表示一个人，包含姓名和年龄
type Person struct {
	Name string // 首字母大写，表示这是一个导出的字段
	Age  int    // 首字母大写，表示这是一个导出的字段
}

// 定义一个结构体 `Student`，表示学生，内嵌 `Person` 结构体，继承了`Person`的字段
type Student struct {
	Person     // 内嵌结构体，继承 `Person` 的字段
	ID     int // 学生特有的学号，首字母大写
}

func main() {
	// 1. 直接创建一个 `Person` 实例
	p1 := Person{"Bob", 20}
	fmt.Println("1. 直接创建Person实例 p1：", p1)

	// 2. 使用命名字段创建 `Person` 实例，可以更清晰地看到字段赋值
	p2 := Person{Name: "Alice", Age: 25}
	fmt.Println("2. 使用命名字段创建Person实例 p2：", p2)

	// 3. 使用点操作符 `.` 访问结构体字段
	fmt.Println("3. 访问p1的Name字段：", p1.Name)

	// 4. 使用结构体指针访问字段，Go会自动解引用
	sp := &p1
	fmt.Println("4. 通过指针访问p1的Name字段：", sp.Name)

	// 5. 通过结构体指针修改字段的值（修改p1的Age）
	sp.Age = 30
	fmt.Println("5.1 通过指针修改p1的Age字段为30：", sp.Age)
	fmt.Println("5.2 修改后，直接访问p1的Age字段：", p1.Age)

	// 6. 创建一个 `Student` 实例，并初始化其字段
	s2 := Student{}
	s2.Name = "Charlie"
	s2.Age = 18
	s2.ID = 101
	fmt.Println("6. 创建Student实例 s2，并初始化字段：", s2)

	// 7. 将 `Student` 结构体实例序列化为JSON格式
	s2json, e := json.Marshal(s2)
	if e != nil {
		fmt.Println("7. JSON序列化出错：", e)
	} else {
		fmt.Println("7. Student结构体 s2 转换为 JSON：", string(s2json))
	}
}
