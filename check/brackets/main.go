package main

import (
	"fmt"
	"os"
)

func isok(s string) bool {
	mapp := make(map[rune]rune)
	mapp['}'] = '{'
	mapp[')'] = '('
	mapp[']'] = '['
	stack := []rune{}

	for _, ch := range s {
		// push to stack
		if ch == '[' || ch == '{' || ch == '(' {
			stack = append(stack, ch)
			// pop from stack
		} else if ch == ']' || ch == '}' || ch == ')' {
			if len(stack) > 0 && mapp[ch] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isOkVer2(s string) bool {
	mapp := make(map[rune]rune)
	mapp['}'] = '{'
	mapp[')'] = '('
	mapp[']'] = '['
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack = append(stack, c)
		case ']', '}', ')':
			if len(stack) > 0 && mapp[c] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]

	// print nothing if  no args are provided
	if len(args) == 0 {
		return
	}

	// iterate through each argument prvided
	for _, str := range args {
		if isOkVer2(str) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
		return
	}
}
