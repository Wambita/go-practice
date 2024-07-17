package main

import (
	"os"
)

// checking for overflows
// if an overflow is encountered , it goes to the other side of the range, no overflow in division
func overflowAdd(value1 int, value2 int) bool {
	c := value1 + value2
	return value1 > 0 && value2 > 0 && c < 0
}

func overflowSub(value1 int, value2 int) bool {
	c := value1 - value2
	return value1 < 0 && value2 < 0 && c > 0
}

func overflowMul(value1 int, value2 int) bool {
	c := value1 * value2
	if value1 != 0 {
		return c/value1 != value2
	} else if value2 != 0 {
		return c/value2 != value1
	}
	return false
}

func PrintString(s string) {
	_, _ = os.Stdout.WriteString(s)
}

// Atoi
func Atoi(s string) (int, bool) {
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
			return 0, false
		} else {
			if overflowMul(number, 10) {
				return 0, false
			}
			if overflowAdd(number*10, int(char-'0')) {
				return 0, false
			}
			number = number*10 + int(char-'0')
		}
	}

	return number * sign, true
}

// Itoa
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

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	operator := args[1]
	var result int

	value1Int, ok1 := Atoi(args[0])
	value2Int, ok2 := Atoi(args[2])
	if !ok1 || !ok2 {
		return
	}

	// operators

	switch operator {
	case "+":
		if overflowAdd(value1Int, value2Int) {
			return
		}
		result = value1Int + value2Int

	case "-":
		if overflowSub(value1Int, value2Int) {
			return
		}
		result = value1Int - value2Int
	case "*":
		if overflowMul(value1Int, value2Int) {
			return
		}
		result = value1Int * value2Int
	case "/":
		if value2Int == 0 {
			PrintString("No division by 0\n")
			return
		}
		result = value1Int / value2Int
	case "%":
		if value2Int == 0 {
			PrintString("No modulo by 0\n")
			return
		}
		result = value1Int % value2Int
	default:
		return
	}

	resultStr := Itoa(result)
	PrintString(string(resultStr))
	PrintString("\n")
}

// functions
// Atoi
// Itoa
// overflowAdd
// overflowSub
// overflowMul
// main
// print string
