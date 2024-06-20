package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	word := os.Args[1]

	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			z01.PrintRune(('a' + 'z') - char)
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune(('A' + 'Z') - char)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
