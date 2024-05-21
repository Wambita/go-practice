package main

import (
	"fmt"
	"os"
)

/*
	Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

	z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

The output will be followed by a newline ('\n').

If the number of arguments is different from 1, the program displays nothing.

usage:
$ go run . "abc"
nop
$ go run . "hello there"
uryyb gurer
$ go run . "HELLO, HELP"
URYYB, URYC
$ go run .
$
*/
func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	result := Rot13(s)
	fmt.Println(result)
}

func Rot13(s string) string {
	srune := []rune(s)
	var result string

	for _, char := range srune {
		// lowercase letters
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+13)%26)

			// uppercase characters
		} else if char >= 'A' && char <= 'Z' {
			result += string('A' + (char-'A'+13)%26)
		} else {
			// append non-letter characters unchanged
			result += string(char)
		}
	}
	return result
}
