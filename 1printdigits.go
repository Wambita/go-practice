package main

import "github.com/01-edu/z01"

func maina() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	printDigits(678)
	z01.PrintRune('\n')
}

// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').

// attempt2 with recursion
func printDigits(n int) {
	// base case just one digit
	if n < 10 {
		z01.PrintRune('0' + rune(n))
		return
	}

	// case 2 more than 1 digit
	// recursion to print the digits except the last one
	printDigits(n / 10)
	z01.PrintRune('0' + rune(n%10))
}
