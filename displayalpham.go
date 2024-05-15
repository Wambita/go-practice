package main

import "github.com/01-edu/z01"

// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').
func main() {
	stringd := "abcdefghijklmnopqrstuvwxyz"
	stringd2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, c := range stringd {
		i = i + 1
		if (i % 2) == 0 {
			c = c - ('a' - 'A')
			z01.PrintRune(c)
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

	for i, c := range stringd2 {
		i = i + 1
		if (i % 2) == 0 {
			z01.PrintRune(c)
		} else {
			c = c + ('a' - 'A')
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
