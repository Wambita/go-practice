package main

import "github.com/01-edu/z01"

// function that returns the last rune in a string
func maino() {
	z01.PrintRune(LastRune(" Hello!"))
	z01.PrintRune(LastRune("Salut! "))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')

	z01.PrintRune('\n')
	z01.PrintRune(lastRune(" Hello!"))
	z01.PrintRune(lastRune("Salut! "))
	z01.PrintRune(lastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	srune := []rune(s)

	for i := len(s) - 1; i >= 0; i-- {
		if srune[i] != ' ' {
			return srune[i]
		}
	}
	return ' '
}

// lastrune easy attempt
func lastRune(s string) rune {
	srune := []rune(s)
	return srune[len(srune)-1]
}
