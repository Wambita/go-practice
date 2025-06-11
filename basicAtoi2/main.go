package main

import (
	"fmt"
	
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	res := 0 
	for _,char := range s {
		if  !(char >= '0' && char <= '9' ){
			return 0;
		}
		dig := int(char - '0')
		res = res*10 + dig
	}
	return res
}