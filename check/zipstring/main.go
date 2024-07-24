package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func alph(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
func ZipString(s string) string {

	if len(s) == 0 {
		return ""
	}

	currentChar := s[0]
	count := 1
	result := ""

	for i := 1; i < len(s); i++ {
		if s[i] == currentChar && alph(s[i]) {
			count++
		} else {
			result += strconv.Itoa(count) + string(currentChar)
			currentChar = s[i]
			count = 1
		}
	}

	if !alph(currentChar) {
		result += string(currentChar)
	} else {
		result += strconv.Itoa(count) + string(currentChar)
	}
	return result
}
