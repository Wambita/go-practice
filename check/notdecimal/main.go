package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	res := ""

	if !isNum(dec) {
		return dec + "\n"
	}

	for i := 0; i < len(dec); i++ {
		if string(dec[i]) == "." {
			res += ""
		} else {
			res += string(dec[i])
		}
	}

	return res + "\n"
}

func isNum(s string) bool {
	for _, c := range s {
		if !(c >= '0' && c <= '9')&& c!='.' && c!='-' {
			return false
		}
	}
	return true
}
