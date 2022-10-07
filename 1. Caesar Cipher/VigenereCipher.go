package main

import (
	"fmt"
	"math"
	"strings"
)

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	text := "CRYPTOGRAHY AND HASHING"
	encrypt := vigenere_encrypt(text, "GOLANG")
	decrypt := vigenere_decrypt(encrypt, "GOLANG")
	fmt.Println(encrypt)
	fmt.Println(decrypt)
}

func vigenere_encrypt(plain_text, key string) string {
	plain_text = strings.ToUpper(plain_text)
	key = strings.ToUpper(key)
	// it will represent the index of the letters of the key
	key_idx := 0
	cipher_text := ""
	// we have to consider all the characters in the plain_text
	for _, s := range plain_text {
		// xi = character that we want to encrypt
		// ki = key for encrypting the
		// formula is Ei(xi) = (xi + ki) % 26
		// we have i-th letter of the key for encrypting the i-th letter
		txt := strings.Index(TEXT, string(s))
		k := strings.Index(TEXT, string(key[key_idx]))
		idx := (txt + k) % len(TEXT)
		cipher_text += string(TEXT[idx])
		key_idx += 1
		if key_idx == len(key) {
			key_idx = 0
		}
	}
	return cipher_text

}

func vigenere_decrypt(cipher_text, key string) string {
	cipher_text = strings.ToUpper(cipher_text)
	key = strings.ToUpper(key)
	plain_text := ""
	key_idx := 0
	for _, s := range cipher_text {
		// xi = character that we want to decrypt
		// ki = key for encrypting the
		// formula is Di(xi) = (xi - ki) % 26
		// we have i-th letter of the key for encrypting the i-th letter
		txt := strings.Index(TEXT, string(s))
		k := strings.Index(TEXT, string(key[key_idx]))
		idx := (txt - k) % len(TEXT)
		if idx < 0 {
			idx = len(TEXT) - int(math.Abs(float64(idx)))
		}
		plain_text += string(TEXT[idx])
		key_idx += 1
		if key_idx == len(key) {
			key_idx = 0
		}
	}
	return plain_text
}
