package main

import (
	"encoding/base64" // 提供 Base64 编码和解码的功能
	"fmt"
)

func main() {
	data := "123456abc"
	base64Data := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(base64Data)

	url := "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu"
	base64Url := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(base64Url)

	// 会将 +、/ 替换为 -、_
	urlSafe := "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu+"
	base64UrlSafe := base64.URLEncoding.EncodeToString([]byte(urlSafe))
	fmt.Println(base64UrlSafe)

	// 使用 urlEncode 加密的，必须用 urlDecode
	urlSafeDecode, _ := base64.URLEncoding.DecodeString(base64UrlSafe)
	fmt.Println(string(urlSafeDecode))

}
