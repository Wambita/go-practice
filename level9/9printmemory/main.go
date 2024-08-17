package main

import "github.com/01-edu/z01"

// memory is in hex
func hex(b byte) string {
	hex := "0123456789abcdef"
	return string(hex[b/16]) + string(hex[b%16])
}

func printstr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		// byte 0
		hexstr := hex(b)
		printstr(hexstr)

		// spacing
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	//newline after everything
z01.PrintRune('\n')

//ascii representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			printstr(string(b))
		} else {
			printstr(".")
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
