package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	plaintext := "This is a secret"

	encrpyt := EncryptAES([]byte(key), plaintext)

	// plaintext
	fmt.Println(plaintext)

	// ciphertext
	fmt.Println(encrpyt)

	decypted := DecryptAES([]byte(key), encrpyt)
	fmt.Println("Decrypted", decypted)
}

func EncryptAES(key []byte, plain_text string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	out := make([]byte, len(plain_text))

	c.Encrypt(out, []byte(plain_text))

	return hex.EncodeToString(out)

}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	pt := make([]byte, len(ciphertext))
	cipher.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}
