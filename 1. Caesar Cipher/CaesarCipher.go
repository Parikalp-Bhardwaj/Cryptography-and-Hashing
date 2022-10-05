package main

import (
	"fmt"
	"math"
	"strings"
)

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}0987654321"
const PRIVATE_KEY = 3

func main() {
	fmt.Printf("length is  %v \n", len(TEXT))
	inp := "cryptography and hashing!321"
	encrypt := encryption(inp)
	fmt.Println(encrypt)
	decrypt := decryption(encrypt)
	fmt.Println(strings.ToLower(decrypt))
}

func encryption(input string) string {
	msg := strings.ToUpper(input)
	cipher_text := ""
	// x = input
	// n = PRIVATE_KEY
	// En(x) = (x+n) % len(Text)
	for _, i := range msg {
		idx := strings.Index(TEXT, string(i))
		// encryption formula is
		f := (idx + PRIVATE_KEY) % len(TEXT)
		cipher_text += string(TEXT[f])
	}
	return cipher_text
}

func decryption(encypt string) string {
	plain_text := ""
	// x = encypt
	// n = PRIVATE_KEY
	// Dn(x) = (x-n) % len(TEXT)
	for _, i := range encypt {
		idx := strings.Index(TEXT, string(i))
		// decryption formula is
		f := (idx - PRIVATE_KEY) % len(TEXT)
		if f < 0 {
			// if f is negative value
			// f = f + len(TEXT)
			f = len(TEXT) - int(math.Abs(float64(f)))
		}
		plain_text += string(TEXT[f])
	}
	return plain_text
}
