package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main()  {

	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	s1 := args[0]
	s2 := args[1]

	result := ""
	for _, r1 := range s1 {
		if Contains(result, r1) {
			continue
		}
		for _, r2 := range s2 {
			if r1 == r2 {
				result += string(r1)
				break
			}
		}
	}

	printString(result)
	printString("\n")


	
}

func Contains(s string, r rune) bool {
	for _, r2 := range s {
		if r == r2 {
			return true
		}
	}
	return false
}

func printString( s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

