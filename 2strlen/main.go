package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
	StrLen2("Hello World!")
}

func StrLen(s string) int {
	return len(s)
}

func StrLen2(s string) int {
	count := 0
	for range s {
		count++
	}
	PrintDigits(count)
	z01.PrintRune('\n')
	return count
}

func PrintDigits(n int) {
	if n < 10 {
		z01.PrintRune('0' + rune(n))
		return
	}
	PrintDigits(n / 10)
	z01.PrintRune('0' + rune(n%10))
}
