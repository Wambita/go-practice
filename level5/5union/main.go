package main

import (
	"os"

	"github.com/01-edu/z01"
)

// check if char is present in string to prevent duplicates
func Contain(result string, char rune) bool {
	for _, c := range result {
		if char == c {
			return true
		}
	}
	return false
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func union(word1, word2 string) string {
	sentence := word1 + word2
	var result string
	for _, char := range sentence {
		if !Contain(result, char) {
			result += string(char)
		} else {
			continue
		}
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	word1 := os.Args[1]
	word2 := os.Args[2]
	result := union(word1, word2)
	Print(result)
}
