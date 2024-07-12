package main

import (
	"fmt"
	"os"
)

func HasVowel(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

func piglatin(s string) string {
	result := ""
	for i, char := range s {
		if i == 0 && HasVowel(char) {
			result = s + "ay"
			return result
		} else if i != 0 && HasVowel(char) {
			before := s[:i]
			after := s[i:]
			result = after + before + "ay"
			return result
		}
	}
	result = "No vowels"
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	fmt.Println(piglatin(word))
}
