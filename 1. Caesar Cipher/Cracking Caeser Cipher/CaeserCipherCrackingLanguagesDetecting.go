package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var ENGLISH_WORDS []string

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}0987654321"

func main() {
	// text := "The most important property of a program is that it is correct. For the Caesar Cipher, Latin letters are rotated, in either direction, by a fixed shift amount. Decoding is the reverse of encoding."
	get_data()

	text := "FU@SWRJUDSK@CDQGCKDVKLQJAB"
	crake_caesar(text)
}
func get_data() {
	file, err := os.Open("./english_words.txt")
	// we can use a dictionary and check whether the given words are
	// are present in a dictionary or not

	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ENGLISH_WORDS = append(ENGLISH_WORDS, scanner.Text())
	}

	file.Close()
	fmt.Println(len(ENGLISH_WORDS))

}

func count_words(text string) int {
	text = strings.ToUpper(text)
	// transform the text into a list of words (split by spaces)
	mystr := strings.Split(text, ` `)

	str := make(map[string]bool)
	for _, w := range mystr {
		str[string(w)] = true
	}
	// counts the number of english words in the text
	var count int
	count = 1

	// consider all the words in the text and check whether the
	// given word in english or not

	for _, w := range mystr {
		if str[w] {
			count += 1
		}
	}

	return count

}

func is_text_english(text string) bool {
	// number of English words in a given text
	matches := count_words(text)
	// you can define your own classification alogrithm in this case
	// the assumption: if 50% of the words in the text are english words
	// then it is an enlgish text
	if (float64(matches)/float64(len(strings.Split(text, ` `))))*100 >= 50 {
		return true
	} else {
		return false
	}
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
		if is_text_english(decode_text) {
			fmt.Printf("With key %v the result is %v \n", key, decode_text)
		}
	}
}
