package main

// In cryptography, a one-time pad is a system in which a randomly generated private key
// is used only once to encrypt a message that is then decrypted by the receiver using a
// matching one-time pad and key.

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ().,"

func main() {
	message := "Rust has different behaviour than other languages. In a language where variables are always references (like Java or Python)"
	fmt.Println("Original Text is ", message)
	key := RandomGenerate(message)
	cipher := encrypt(message, key)
	fmt.Println("Encrpyt text ", cipher)
	plain := decrypt(cipher, key)
	fmt.Println("Decrpyt text ", plain)

}

func encrypt(plain_text string, key []int) string {
	plain_text = strings.ToUpper(plain_text)
	cipher_text := ""
	for idx, char := range plain_text {
		// the value with which we shift the given letter
		// xi = character that we want to encrypt
		indx_char := strings.Index(TEXT, string(char))
		// otp = now every key will assign to char
		opt := key[idx]
		// formula is Ei(xi) = (xi + opt) % 26
		// we have i-th letter of the key for encrypting the i-th letter
		cipher_text += string(TEXT[(indx_char+opt)%len(TEXT)])
	}
	return cipher_text
}

func decrypt(cipher_text string, key []int) string {
	cipher_text = strings.ToUpper(cipher_text)
	plain_text := ""
	for idx, char := range cipher_text {
		// the value with which we shift the given letter
		// xi = character that we want to encrypt
		indx_char := strings.Index(TEXT, string(char))
		// otp = now every key will assign to char
		opt := key[idx]
		// formula is Ei(xi) = (xi - opt) % 26
		// we have i-th letter of the key for encrypting the i-th letter
		dx := (indx_char - opt) % len(TEXT)
		if dx < 0 {
			dx = len(TEXT) - int(math.Abs(float64(dx)))
		}
		plain_text += string(TEXT[dx])
	}
	return plain_text
}

// First we have to generate random number
func RandomGenerate(mytext string) []int {
	random := make([]int, 0, len(mytext))
	max := len(TEXT) - 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(mytext); i++ {
		random = append(random, rand.Intn(max-0)+0)
	}

	return random
}
