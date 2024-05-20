package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Write a program that takes a string and displays its last word, followed by a newline ('\n').
// A word is a section of string delimited by spaces or by the start/end of the string.
// The output will be followed by a newline ('\n').
// If the number of arguments is different from 1, or if there are no word, the program displays nothing.
func main() {
	if len(os.Args) == 2 {
		var word string
		input := os.Args[1]
		if input == " " {
			return
		}
		sentence := []byte(input)
		for i := len(sentence) - 1; i >= 0; i-- {
			if sentence[i] == ' ' {
				break
			} else {
				word += string(sentence[i])
			}
		}
		words := []rune(word)
		for i := len(words) - 1; i >= 0; i-- {
			z01.PrintRune(words[i])
		}
		z01.PrintRune('\n')

	} else {
		return
	}
}