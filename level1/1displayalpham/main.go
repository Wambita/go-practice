package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i + ('A' - 'a'))
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
