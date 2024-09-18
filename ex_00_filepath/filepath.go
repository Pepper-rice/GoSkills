package main

import (
	"fmt"
	"path/filepath" // 提供了一些函数来操作 文件路径
)

func main() {
	// 可移植的构建路径，应该使用 join 而不是直接拼接
	p := filepath.Join("dir1", "dir2", "filename.jpg")
	fmt.Println("p:", p)

	// Join 会移除多余的分隔符，规范化路径
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir2", "filename"))

	// Dir 和 Base 可以用来分割路径中的目录和文件
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// 扩展名
	fmt.Println("extend(p):", filepath.Ext(p))

	// 计算相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
