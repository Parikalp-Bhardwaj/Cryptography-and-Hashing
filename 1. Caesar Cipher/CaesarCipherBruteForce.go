package main

import (
	"fmt"
	"math"
	"strings"
)

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}0987654321"

func main() {

	crake_caesar("FU@SWRJUDSK@CDQGCKDVKLQJ$ AB")
}

func crake_caesar(str string) {
	// we try all possible key value
	for key := 0; key < len(TEXT); key++ {
		decode_text := ""
		for _, char := range str {
			idx := strings.Index(TEXT, string(char))
			//  formula how we can decrypt the text 
            //  Dn(x) = (x - n) % len(ALPHABET)
			f := (idx - key) % len(TEXT)
			if f < 0 {
				// if f is negative value
				// f = f + len(TEXT)
				f = len(TEXT) - int(math.Abs(float64(f)))
			}
			decode_text += string(TEXT[f])

		}
		fmt.Printf("With key %v the result is %v \n", key, decode_text)
	}
}
