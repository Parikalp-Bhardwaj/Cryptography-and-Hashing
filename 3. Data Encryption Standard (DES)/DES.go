package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	originalText := "this is secret text$$$"
	fmt.Println(originalText)
	mytext := []byte(originalText)

	key := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}
	iv := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}

	cryptoText, _ := encrypt(key, iv, mytext)
	fmt.Println(string(cryptoText))
	decryptedText, _ := decrypt(key, iv, cryptoText)
	fmt.Println(string(decryptedText))

}

func encrypt(key, iv, mytext []byte) ([]byte, error) {

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println(block, block.BlockSize())
	blockSize := block.BlockSize()
	data := add_padding(mytext, blockSize)
	fmt.Println(data)
	mode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(data))
	mode.CryptBlocks(crypted, []byte(data))
	return crypted, nil

}

func decrypt(key, iv, cryptoText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cryptoText))
	blockMode.CryptBlocks(origData, cryptoText)
	origData = remove_padding(origData)
	return origData, nil
}

func add_padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func remove_padding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
