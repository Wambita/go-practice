package main

import "fmt"

func SwapBits(octet byte) byte {
	return (octet >> 4) | (octet << 4)
}

func main() {
	var b byte = 0b00100110
	fmt.Printf("%08b", SwapBits(b))
}
