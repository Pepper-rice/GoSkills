package main

import (
	"crypto/md5"    // MD5的哈希算法
	"crypto/sha1"   // SHA1的哈希算法
	"crypto/sha256" // SHA256的哈希算法
	"encoding/hex"  // 提供 编码 和 解码 十六进制字符串的功能。
	"fmt"
)

func main() {
	str := ""

	md5Bytes := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", md5Bytes)
	fmt.Println(len(md5Str))
	fmt.Println(md5Str)

	fmt.Println(hex.EncodeToString(md5Bytes[:]))

	sha1Bytes := sha1.Sum([]byte(str))
	sha1Str := fmt.Sprintf("%x", sha1Bytes)
	fmt.Println(len(sha1Str))
	fmt.Println(sha1Str)

	sha256Bytes := sha256.Sum256([]byte(str))
	sha256Str := fmt.Sprintf("%x", sha256Bytes)
	fmt.Println(len(sha256Str))
	fmt.Println(sha256Str)

}
