package main

import "github.com/01-edu/z01"

// Write a function that returns the first rune of a string.
func main() {
	z01.PrintRune(FirstRune(" Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune(firstRune("   fana"))
	z01.PrintRune('\n')

	z01.PrintRune(firstRune(" Hello!"))
	z01.PrintRune(firstRune("Salut!"))
	z01.PrintRune(firstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	for _, char := range s {
		if char != ' ' {
			return char
		}
	}
	// if string contains only space , return a default rune
	return ' '
}

// attempt 2 easy
func firstRune(s string) rune {
	return rune(s[0])
}
