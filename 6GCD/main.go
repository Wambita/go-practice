package main

import (
	"fmt"
	"os"
)

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isneg := false
	if nb < 0 {
		isneg = true
		nb = -nb
	}
	var num []rune
	for nb > 0 {
		result := nb % 10
		num = append([]rune{rune(result + '0')}, num...)
		nb = nb / 10
	}
	if isneg {
		num = append([]rune{'-'}, num...)
	}
	return string(num)
}

func Atoi(s string) int {
	num := 0
	sign := 1

	if s == "0" {
		return 0
	}
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		} else if i == 0 && c == '+' {
			continue
		}
		num = num*10 + int(c-48)
	}
	return num * sign
}

func GCD(num1, num2 int) string {
	if num1 < 0 {
		return "0"
	}
	if num2 < 0 {
		return "0"
	}
	gcd := 0
	if num1 > num2 {
		for d := 1; d <= num1; d++ {
			if (num1%d) == 0 && (num2%d) == 0 {
				gcd = d
			}
		}
	} else if num2 > num1 {
		for d := 1; d <= num2; d++ {
			if (num2%d) == 0 && (num1%d) == 0 {
				gcd = d
			}
		}
	}
	num := Itoa(gcd)
	return num
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	num1 := Atoi(os.Args[1])
	num2 := Atoi(os.Args[2])
	num := GCD(num1, num2)
	fmt.Println(num)
}
