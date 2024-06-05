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

func nmain() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	word := (os.Args[1])
	var result string
	for _, c := range word {
		result += string(mirror(c))
	}
	for _, i := range result {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func mirror(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return 'A' + 'Z' - r
	} else if r >= 'a' && r <= 'z' {
		return 'a' + 'z' - r
	}
	return r
}

func mirror2(r rune) rune {
	m := make(map[rune]rune)

	r2 := 'z'
	for r := 'a'; r <= 'z' && r2 >= 'a'; r++ {
		m[r] = r2
		r2--
	}

	r2 = 'Z'
	for r := 'A'; r <= 'Z' && r2 >= 'A'; r++ {
		m[r] = r2
		r2--
	}

	out, ok := m[r]
	if !ok {
		return r
	}
	return out
}
