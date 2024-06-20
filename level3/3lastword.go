package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Write a program that takes a string and displays its last word, followed by a newline ('\n').
// A word is a section of string delimited by spaces or by the start/end of the string.
// The output will be followed by a newline ('\n').
// If the number of arguments is different from 1, or if there are no word, the program displays nothing.
func mainpp() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	if input == " " {
		return
	}

	sentence := []rune(input)
	var word string
	var lastWord string

	// iterate over input string in reverse to find last word
	for i := len(sentence) - 1; i >= 0; i-- {
		// if space is encountered and word isn't empty == found last word
		if sentence[i] == ' ' {
			if len(word) > 0 {
				lastWord = word
				break
			}
		} else {
			word = string(sentence[i]) + word // to print correctly
			// if you use word += sentence[i] // have to print in reverse
		}
	}
	for _, char := range lastWord {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
