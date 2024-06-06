package main

import "fmt"

func Atoi(s string) int {
	// variables
	number := 0
	sign := 1

	// loop through each char
	for i, num := range s {
		// check for negative numbers
		if i == 0 && num == '-' {
			sign = -1
			continue
		} else if i == 0 && num == '+' { // positive numbers
			continue
		}

		// check if it ranges between 0-9
		if num < '0' || num > '9' {
			return 0
		} else {
			number = number*10 + int(num-48)
		}
	}
	// correct sign for number
	return number * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
