package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World \n")
}

func headers(writer http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v:%v\n", name, h)
		}
	}
}

func main() {

	// 用于注册路由
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 监听
	http.ListenAndServe(":8080", nil)
}
