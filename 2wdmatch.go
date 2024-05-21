package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string.
// This rewrite must respect the order in which these characters appear in the second string.
// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.
// If the number of arguments is different from 2, the program displays nothing.

func mainp() {
	if len(os.Args) > 3 || len(os.Args) == 1 {
		return
	} else {
		// declare variables
		word1 := os.Args[1]        // fisrt word
		word2 := os.Args[2]        // 2nd word
		word2rune := []rune(word2) // word2 changed to rune for comparison
		var answer string          // store matched chars for comparison
		start := 0                 // store position of char being searched to prevent duplicates

		for _, char := range word1 {
			for i := start; i < len(word2); i++ {
				if char == word2rune[i] {
					answer += string(word2rune[i])
					start = start + 1
					break
				}
			}
		}
		if word1 == answer {
			for _, i := range answer {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
		}
	}
}
