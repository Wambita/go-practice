package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	word1 := os.Args[1]
	word2 := os.Args[2]

	result := Inter(word1, word2)
	PrintString(result)
}

// check if char already exists in the result
func Contain(result string, char rune) bool {
	for _, c := range result {
		if char == c {
			return true
		}
	}
	return false
}

// func inter
func Inter(word1, word2 string) string {
	// stores the matches found
	result := ""
	for _, char := range word1 {
		if Contain(result, char) {
			continue
		} else {
			for _, char2 := range word2 {
				if char == char2 {
					result += string(char2)
					break
				}
			}
		}
	}
	return result
}

// print string as runes
func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
