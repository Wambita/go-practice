package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Atoi
func Atoi(word string) int {
	number := 0
	sign := 1

	for index, char := range word {
		if index == 0 && char == '-' {
			sign = -1
			continue
		} else if index == 0 && char == '+' {
			continue
		}

		if char < '0' && char > '9' {
			return 0
		} else {
			number = (number * 10) + int(char-'0')
		}
	}
	return number * sign
}

// Itoa
func Itoa(number int) string {
	if number == 0 {
		return "0"
	}
	isNegative := false

	if number < 0 {
		isNegative = true
		number = -number
	}

	var answer []rune
	for number > 0 {
		digit := number % 10
		answer = append([]rune{rune(digit + '0')}, answer...)
		number = number / 10
	}

	if isNegative {
		answer = append([]rune{'-'}, answer...)
	}
	return string(answer)
}

// PrintString
func PrintString(word string) {
	for _, char := range word {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// tabmult
func main() {
	if len(os.Args) != 2 {
		return
	}
	number := os.Args[1]
	if Atoi(number) >= 2147483647 || Atoi(number) < -2147483647 {
		errorString := "overflow error"
		PrintString(errorString)
		return
	}

	for i := 1; i <= 9; i++ {
		AnswerString := Itoa(i) + " " + "x" + " " + number + " " + "=" + " " + Itoa(i*Atoi(number))
		PrintString(AnswerString)
	}
}
