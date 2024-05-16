package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program that displays its last argument if there is one
func main() {
	if len(os.Args) > 1 {
		lastparam := os.Args[len(os.Args)-1]
		for _, c := range lastparam {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
