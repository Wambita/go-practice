package main

import (
	"github.com/01-edu/z01"
)

func mainvv() {
	StrLen("htis")
}

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	PrintDigits(count)
	z01.PrintRune('\n')
	return count
}

// func to print digits
func PrintDigits(count int) {
	// Case 1 if num is less that 10 , print and return
	if count < 10 {
		z01.PrintRune('0' + rune(count))
		return
	}

	// Recursive call to print digits of the num except the last digit
	PrintDigits(count / 10)
	// Print the last digit
	z01.PrintRune('0' + rune(count%10))
}
