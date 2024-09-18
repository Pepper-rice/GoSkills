package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "file.data"

	// 文件写入
	data := []byte("All the data \nI wish to write to a file")
	err := os.WriteFile(filename, data, 0777)
	if err != nil {
		panic(err)
	}

	// 文件读取
	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))

	// 文件追加
	for i := 0; i < 5; i++ {
		err = WriteFileAppend(filename, []byte("\n"), 0777)
		err = WriteFileAppend(filename, data, 0777)
		if err != nil {
			panic(err)
		}
	}

	// 文件删除
	//os.Remove(filename)
}

func WriteFileAppend(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}

	return err
}
