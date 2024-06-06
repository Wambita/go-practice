package main

import "fmt"

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

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
