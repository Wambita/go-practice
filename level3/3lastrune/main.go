package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune(LastRune("Heloo  "))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			return rune(s[i])
		}
		continue
	}
	return ' '
}
