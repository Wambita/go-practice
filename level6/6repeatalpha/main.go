package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]

	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			count := 0
			for i := 'a'; i <= 'z'; i++ {
				count++
				if char == i {
					for j := 0; j < count; j++ {
						z01.PrintRune(char)
					}
					break
				}
			}
		} else if char >= 'A' && char <= 'Z' {
			count := 0
			for i := 'A'; i <= 'Z'; i++ {
				count++
				if char == i {
					for j := 0; j < count; j++ {
						z01.PrintRune(char)
					}
					break
				}
			}
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
