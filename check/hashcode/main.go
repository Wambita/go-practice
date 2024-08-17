package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	hash := ""
	length := len(dec)
	var hashchar rune
	for _, char := range dec {
		hashchar = ((char + rune(length)) % 127)
		if !(hashchar >= ' ' && hashchar <= '~') {
			hashchar = (char + rune(33))
		}
		hash += string(hashchar)
	}
	return hash
}
