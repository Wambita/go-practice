package main

import (
	"fmt"
)

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}

	isNegative := false
	if nb < 0 {
		isNegative = true
		nb = -nb
	}

	var number []rune
	for nb > 0 {
		digit := nb % 10
		number = append([]rune{rune(digit + '0')}, number...)
		nb /= 10
	}

	if isNegative {
		number = append([]rune{'-'}, number...)
	}

	return string(number)
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
