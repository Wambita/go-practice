package main

import "github.com/01-edu/z01"

// Program that displays "Hello World!" followed by a new line
func hello(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func maind() {
	str := "Hello World!"
	hello(str)
}
