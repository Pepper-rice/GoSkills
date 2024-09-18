package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// 创建 AES 加密器，密钥长度必须是 16、24 或 32 字节，分别对应 128、192 和 256 位
	key := []byte("0123456789abcdef")
	iv := []byte("0123456789abcdef")
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		return
	}

	// 准备需要加密的数据
	plaintext := []byte("Hello, world!")

	// 对数据进行填充，使其满足加密块大小，加密块大小为 16 字节
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	// 创建加密块链，使用 CBC 加密模式，iv 的长度需要和 block.BlockSize() 一致
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密数据
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	cipherHex := hex.EncodeToString(ciphertext)

	fmt.Printf("plaintext: %s\n", plaintext)
	fmt.Printf("ciphertext: %x\n", ciphertext)
	fmt.Printf("ciphertext: %s\n", cipherHex)

	// 解密

	cipherbytes, _ := hex.DecodeString(cipherHex)
	// 创建解密器
	block, _ = aes.NewCipher(key)

	// 创建解密块链
	unmode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	plainbytes := make([]byte, len(cipherbytes))
	unmode.CryptBlocks(plainbytes, cipherbytes)

	plainbytes = PKCS7UnPadding(plainbytes)

	fmt.Printf("decode plaintext: %s\n", plainbytes)
}

// PKCS7Padding 对数据进行填充，满足加密块大小
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7UnPadding 去除填充的数据
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
