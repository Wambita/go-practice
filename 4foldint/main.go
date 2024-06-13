package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	Mul := func(acc int, cur int) int {
		return acc * cur
	}
	Add := func(acc int, cur int) int {
		return acc + cur
	}
	Sub := func(acc int, cur int) int {
		return acc - cur
	}
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	a = append(a, n)

	if len(a) == 0 {
		return
	}
	result := a[0]
	for i := 0; i < len(a)-1; i++ {
		result = f(abs(result), abs(a[i+1]))
	}

	result2 := Itoa(result)
	PrintString(result2)
}

// Printstring
func PrintString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

// Itoa
func Itoa(nb int) string {
	// if number is zero return 0
	if nb == 0 {
		return "0"
	}

	// variables
	isneg := false
	var number []rune

	// negative numbers
	if nb < 0 {
		isneg = true
		nb = -nb // absolute value
	}

	// atoi conversion
	for nb > 0 {
		result := nb % 10 // get last number adding last number
		number = append([]rune{rune(result + '0')}, number...)
		nb = nb / 10
	}
	// append - sign if number is negative
	if isneg {
		number = append([]rune{'-'}, number...)
	}
	return string(number)
}

// absolute values
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
