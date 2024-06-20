package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	var result string
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+13)%26)
		} else if char >= 'A' && char <= 'Z' {
			result += string('A' + (char-'A'+13)%26)
		} else {
			result += string(char)
		}
	}
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
