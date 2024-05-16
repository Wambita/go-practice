package main

import "github.com/01-edu/z01"

// program that displays the alphabet in reverse order with even letters in upper case , odd in lower case  followed by a new line
func mainv() {
	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 1 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')

	strn := "abcdefghijklmnopqrstuvwxyz"
	for i := len(strn) - 1; i >= 0; i-- {
		if i%2 == 0 {
			z01.PrintRune(rune(strn[i]) - 32)
		} else {
			z01.PrintRune(rune(strn[i]))
		}
	}
	z01.PrintRune('\n')
}
