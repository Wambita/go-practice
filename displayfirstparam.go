package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// Write a program that displays its first argument, if there is one.
// too easy to be the solution
func Display() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
}

func mainm() {
	if len(os.Args) > 1 {
		firstparam := os.Args[1]
		for _, c := range firstparam {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
