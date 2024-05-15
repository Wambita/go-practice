package main

import (
	"fmt"
)

func maincc() {
	stringl := StrLen("htis ososos jhwhdawid")
	fmt.Println(stringl)
}

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	// z01.PrintRune('0' + count)
	return count
}
