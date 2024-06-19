package main

import (
	"os"
)

func Atoi(s string) int {
	number := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
			continue
		} else if i == 0 && char == '+' {
			continue
		}
		if char < '0' || char > '9' {
			return 0
		} else {
			number = number*10 + int(char-'0')
		}
	}
	return number * sign
}

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isnegative := false

	if nb < 0 {
		isnegative = true
		nb = -nb
	}

	var number []rune

	for nb > 0 {
		digit := nb % 10
		number = append([]rune{rune(digit + '0')}, number...)
		nb = nb / 10
	}
	if isnegative {
		number = append([]rune{'-'}, number...)
	}
	return string(number)
}

// is digit
func isDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, char := range s {

		if i == 0 && (char == '-' || char == '+') {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	value1 := args[0]
	operator := args[1]
	value2 := args[2]
	var result int

	if !(isDigit(value1)) || !(isDigit(value2)) {
		return
	}

	value1Int := Atoi(value1)
	value2Int := Atoi(value2)

	if (value1Int >= 9223372036854775807) || (value1Int <= -9223372036854775808) {
		return
	}
	if (value2Int >= 9223372036854775807) || (value2Int <= -9223372036854775808) {
		return
	}

	// operators

	switch operator {
	case "+":
		result = value1Int + value2Int
	case "-":
		result = value1Int - value2Int
	case "*":
		result = value1Int * value2Int
	case "/":
		if value2Int == 0 {
			message := "No division by 0\n"
			_, err := os.Stdout.WriteString(message)
			if err != nil {
				return
			}
			return
		}
		result = value1Int / value2Int
	case "%":
		if value2Int == 0 {
			message := "No modulo by 0"
			_, err := os.Stdout.Write([]byte(message))
			if err != nil {
				return
			}
			return
		}
		result = value1Int % value2Int
	default:
		return
	}

	resultStr := Itoa(result)
	resultbyte := []byte(resultStr)
	_, err := os.Stdout.Write(resultbyte)
	if err != nil {
		return
	}
}
