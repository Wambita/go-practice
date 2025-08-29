package main

import (
	"fmt"
)

func allUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func longestUniqueWord(s string) string {
	words := Split(s, " ")
	longest := ""

	for _, word := range words {
		if allUnique(word) && len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func Split(s, delim string) []string {
	var res []string
	word := ""

	i := 0
	dlen := len(delim)
	for i < len(s) {
		if i+dlen <= len(s) && s[i:i+dlen] == delim {
			if word != "" {
				res = append(res, word)
				word = ""
			}
			i += dlen
		} else {
			word += string(s[i])
			i++
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}

func main() {
	text := "hello world kite book sun dream karinbo"
	fmt.Println("Longest word with unique characters:", longestUniqueWord(text))
}
