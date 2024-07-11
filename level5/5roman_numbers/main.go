package main

import (
	"fmt"
	"os"
	"strconv"
)

func Roman(num int) (string, string) {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calcSymbols := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	roman := ""
	calculation := ""

	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			num -= value[i]
			roman += symbol[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += calcSymbols[i]
		}
	}

	return roman, calculation
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error cannot convert digit to roman")
		return
	}
	if num <= 0 || num >= 4000 {
		fmt.Println("Error cannot convert digit to roman")
		return
	}
	roman, calc := Roman(num)
	fmt.Println(calc)
	fmt.Println(roman)
}
