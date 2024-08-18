package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	srune := []rune(s)
	word := ""
	var firstword string

	if s == " " || s == "" {
		return "\n"
	}

	for i := 0; i < len(s); i++ {
		if !(string(srune[i]) == " ") {
			word = string(srune[i]) + word
		}
		if string(srune[i]) == " " {
			if len(word) > 0 {
				firstword = word
				break
			}
		}
	}

	return sort(firstword) + "\n"
}

func sort(s string) string {
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}
