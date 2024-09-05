package main

import (
	"fmt"
	"os"
)

func split(s, sep string) []string {
	res := []string{}
	start := 0

	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if start < i {
				res = append(res, s[start:i])
			}
			start = i + len(sep)
		}
	}
	if start < len(s) {
		res = append(res, s[start:])
	}
	return res
}

func join(s []string) string {
	res := ""
	for i, c := range s {
		res += c
		if i < len(s)-1 {
			res += " "
		}
	}
	return res
}

func rostring(s []string) []string {
	if len(s) == 0 {
		return s
	}
	res := append(s[1:], s[0])
	return res
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	words := os.Args[1]
	w1 := split(words, "  ")
	fmt.Println(w1)
	rotated := rostring(w1)
	fmt.Println(rotated)
	rotstring := join(rotated)
	fmt.Println(rotstring)
}
