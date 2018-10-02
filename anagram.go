package main

import (
	"fmt"
	"unicode"
)

type runeCounter map[rune]int

func anagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Count := runeCounter{}
	str2Count := runeCounter{}

	for _, v := range str1 {
		str1Count[unicode.ToLower(v)]++
	}

	for _, v := range str2 {
		str2Count[unicode.ToLower(v)]++
	}

	for k, v := range str1Count {
		if v != str2Count[k] {
			return false
		}
	}

	return true
}

func main() {
	r := anagram("iceman", "cinema")
	fmt.Printf("anagram() result: %t\n", r)
}
