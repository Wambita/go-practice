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

	for _, char := range word {
		for _, c := range LetterToBeReplaced {
			for _, v := range LetterToBeReplacedWith {
				if char == c {
					finalString += string(v)
				} else {
					finalString += string(char)
				}
			}
		}
		PrintString(finalString)
	}
}

func PrintString(word string) {
	for _, char := range word {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
