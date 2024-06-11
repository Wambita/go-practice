package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	word := os.Args[1]
	LetterToBeReplaced := os.Args[2]
	LetterToBeReplacedWith := os.Args[3]
	var finalString string

	if len(LetterToBeReplaced) > 1 || len(LetterToBeReplacedWith) > 1 {
		j := 0
		for i := 0; i < len(word) && j < len(LetterToBeReplaced); i++ {
			if word[i] == LetterToBeReplaced[i] {
				finalString += string(LetterToBeReplacedWith[i])
				j++
			} else {
				finalString += string(word[i])
			}
		}
	} else {
		// use this for simplicity
		for _, char := range word {
			if string(char) == LetterToBeReplaced {
				finalString += LetterToBeReplacedWith
			} else {
				finalString += string(char)
			}
		}
	}
	PrintString(finalString)
}

func PrintString(word string) {
	for _, char := range word {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
