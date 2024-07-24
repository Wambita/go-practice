package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foot", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abcn", "def"))
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}
	count := 0
	charMap := make(map[rune]int)
	

	//str1 map
	for _, char := range str1 {
		charMap[char]++
	}

	//str2 map
	for _, char := range str2 {
		charMap[char]++
	}
	//count

	for _, v := range charMap {
		if v == 1 {
			count++
		}
	}
	return count
}
