package main

import (
	"fmt"
	"os" // 提供操作系统功能接口，包括文件和目录操作。
)

func main() {
	// 获取当前工作目录
	pwd, err := os.Getwd()
	fmt.Println(pwd)

	//err = Mkdir("subdir", 0755)
	//if err != nil {
	//	panic(err)
	//}

	// 当前目录创建一个子目录
	//err = os.Mkdir("subdir", 0755)
	//if err != nil {
	//	panic(err)
	//}

	// 创建一个子目录树，类似于 Unix 的 mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	if err != nil {
		panic(err)
	}

	//  删除一个文件或者空的目录，如果非空会返回一个错误
	defer func() {
		err := os.Remove("subdir")
		if err != nil {
			panic(err)
		}
	}()

	// 删除一个目录，等同于 Unix 的 rm -rf，不检测目录是否为空， 慎用
	//defer os.RemoveAll("subdir")
}

func Mkdir(dir string, perm os.FileMode) error {
	if !DirExists(dir) {
		return os.Mkdir(dir, perm)
	}

	return nil
}

// DirExists 检测目录是否存在，返回 bool 或 panic
func DirExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}
