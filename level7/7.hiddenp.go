package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*
Instructions

Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

If s1 is an empty string, it is considered hidden in any string.

If the number of arguments is different from 2, the program displays nothing.
Usage

$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
$ go run .
$
*/

func mainmm() {
	if len(os.Args) != 3 {
		return
	}

	word1 := os.Args[1]
	word2 := os.Args[2]
	start := 0
	var result string
	srune := []rune(word2)

	for _, char := range word1 {
		for i := start; i < len(srune); i++ {
			if char == srune[i] {
				result += string(srune[i])
				start = start + 1
				break
			}
		}
	}
	if word1 == result {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		for _, i := range result {
			z01.PrintRune(i)
		}
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
