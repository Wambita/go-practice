package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string.
// This rewrite must respect the order in which these characters appear in the second string.
// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.
// If the number of arguments is different from 2, the program displays nothing.

func containWord(word1, word2 string) bool {
	i := 0 // track index first word

	for j := 0; j < len(word2) && i < len(word1); j++ {
		if word1[i] == word2[j] {
			i++
		}
	}
	return i == len(word1)
}

func wdmatch(word1, word2 string) string {
	if containWord(word1, word2) {
		return word1
	}
	return ""
}

func mainppppp() {
	if len(os.Args) != 3 {
		return
	}

	word1 := os.Args[1]
	word2 := os.Args[2]

	result := wdmatch(word1, word2)
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
