package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*
Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

The case of the letter remains unchanged, for example :

'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

The final result will be followed by a newline ('\n').

If the number of arguments is different from 1, the program prints a new line.
Usage

$ go run . "abc"
zyx$
$
$ go run . "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
$
$ go run .
$
*/

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := (os.Args[1])
	count := 25 // count from a - m 
	countb := -1 // count from m - z
	for _, c := range word {
		if (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M') {
			z01.PrintRune(c + rune(count))
			count = count - 2
		} else if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			z01.PrintRune(c + rune(countb))
			countb = countb - 2
		
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
