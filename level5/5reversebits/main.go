package main

import "fmt"

func ReverseBits(oct byte) byte {
	var reversed byte = 0
	for i := 0; i < 8; i++ {
		bits := (oct >> i) & 1
		reversed = reversed | (bits << (7 - i))
	}
	return reversed
}

func main() {
	var b byte = 0b00100110
	fmt.Printf("%08b", ReverseBits(b))
}
