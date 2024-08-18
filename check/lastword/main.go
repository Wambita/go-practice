package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	srune := []rune(s)
	word := ""
	lastword := ""

	if s == " " {
		return "\n"
	}

	for i := len(s) - 1; i >= 0; i-- {
		if srune[i] == ' ' {
			if len(word) > 0 {
				lastword = word
			} else {
				word = string(srune[i]) + word
			}
		}
	}
	return lastword + "\n"
}
