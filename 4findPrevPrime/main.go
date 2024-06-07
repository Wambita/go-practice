package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(1))
	fmt.Println(FindPrevPrime(6))
}

func FindPrevPrime(nb int) string {
	if nb <= 0 {
		return "0"
	}
	for i := nb; i >= 0; i-- {
		if IsPrime(i) {
			return Itoa(i)
		}
	}
	return "0"
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

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
