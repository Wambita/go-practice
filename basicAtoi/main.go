package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	res := 0;
	for _,char := range s {
		dig := int(char - '0');
		res = res*10 + dig
	}
	return res
}

//signs  taken into account
