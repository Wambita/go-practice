package main

import (
	"os"

	"github.com/01-edu/z01"
)

// check if word 1 is contained in word 2 in the right order without going back
func containWord(word1, word2 string) bool {
	i := 0 // track 1st word index

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

func main() {
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
