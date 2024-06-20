package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]

	if input == " " || input == "" || len(input) == 0 {
		return
	}

	srune := []rune(input)
	var word string
	var lastWord string

	for i := len(srune) - 1; i >= 0; i-- {
		if srune[i] == ' ' {
			if len(word) > 0 {
				lastWord = word
				break
			}
		} else {
			word = string(srune[i]) + word
		}
	}
	for _, char := range lastWord {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
