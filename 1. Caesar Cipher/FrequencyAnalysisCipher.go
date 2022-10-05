package main

// In cryptography, frequency analysis is the study of the frequency of letters
// or groups of letters in a ciphertext. The method is used as an aid to breaking
// substitution ciphers
import (
	"fmt"
	"sort"
	"strings"
)

const TEXT = " ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}0987654321"

func main() {
	// text := "The most important property of a program is that it is correct. For the Caesar Cipher, Latin letters are rotated, in either direction, by a fixed shift amount. Decoding is the reverse of encoding."
	// frequency_analysis(text)
	text := "FU@SWRJUDSK@CDQGCKDVKLQJAB"
	caesar_crack(text)
}

func frequency_analysis(text string) map[string]int {
	text = strings.ToUpper(text)
	letter_frequency := make(map[string]int)
	isLetter := make(map[string]bool)
	for _, letter := range TEXT {
		letter_frequency[string(letter)] = 0
		isLetter[string(letter)] = true
	}

	for _, letter := range text {
		if isLetter[string(letter)] {
			letter_frequency[string(letter)]++
		}
	}

	return letter_frequency
}

func caesar_crack(text string) {

	freq := frequency_analysis(text)
	keys := make([]string, 0, len(freq))

	for k := range freq {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})
	// fmt.Println(freq)
	var m int
	for _, k := range keys {
		m = freq[k]
		break
	}
	// # now how we can find out the key
	// #formula is
	// # key = value of cipher text most frequent letter - value of E
	// # Most freq letter is LETTERS.find(freq[0][0])
	// # value of E is LETTERS.find('E')
	fmt.Println(m - strings.Index(TEXT, "E"))
}
