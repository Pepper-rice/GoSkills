package main

import (
	"fmt"
	"net/http"
)

// hello 处理函数用于响应 /hello 路径的请求
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World \n")
}

// headers 处理函数用于响应 /headers 路径的请求
func headers(writer http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		// 对于每个字段名称下的所有值进行遍历
		for _, h := range headers {
			// 将头部字段名和值写入响应
			fmt.Fprintf(writer, "%v:%v\n", name, h)
		}
	}
}

func main() {
	// 注册 HTTP 路由处理函数
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	// 启动 HTTP  监听 8080 端口 服务器第二个参数为 nil 表示使用默认的 ServeMux
	http.ListenAndServe(":8080", nil)
}
