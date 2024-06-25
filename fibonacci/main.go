package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	if !isdigit(input) {
		return
	}
	num, _ := strconv.Atoi(input)
	sequence := Seq(num)
	fmt.Println(sequence)
}

// isdigit
func isdigit(num string) bool {
	if len(num) == 0 {
		return false
	}
	for i, char := range num {
		if i == 0 && (char == '-' || char == '+') {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// fibonnacci
func fibonacci(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func Seq(n int) string {
	str := ""
	sep := ", "
	for i := 0; i < n; i++ {
		str += strconv.Itoa(fibonacci(i))
		if i < n && i != n-1 {
			str += sep
		}
	}

	return str
}
