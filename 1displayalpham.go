package main

import (
	"github.com/01-edu/z01"
)

// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').
func mainb() {
	// case 1 : given a string that is sorted in lowercase
	stringd := "abcdefghijklmnopqrstuvwxyz"
	z01.PrintRune('1')
	for i, c := range stringd {
		i = i + 1
		if (i % 2) == 0 {
			c = c - ('a' - 'A')
			z01.PrintRune(c)
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

	// case 2 : given a string that is sorted in uppercase
	z01.PrintRune('2')
	stringd2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, c := range stringd2 {
		i = i + 1
		if (i % 2) == 0 {
			z01.PrintRune(c)
		} else {
			c = c + ('a' - 'A')
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

	// case 3 : given a string that is sorted in both uppercase (creating a function outside)
	z01.PrintRune('3')
	for i, c := range stringd {
		i++
		if i%2 == 0 {
			c = toUpper(c)
			z01.PrintRune(c)
		} else {
			c = toLower(c)
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

	// case 4 : given a string that is sorted in both uppercase (creating a function outside)
	z01.PrintRune('4')
	for i, c := range stringd2 {
		i++
		if i%2 == 0 {
			c = toUpper(c)
			z01.PrintRune(c)
		} else {
			c = toLower(c)
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')

	////case 5 : Converting without use of a string, use runes instead (lowercase)
	z01.PrintRune('5')
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')

	////case 6 : Converting without use of a string, use runes instead (upperrcase)
	z01.PrintRune('6')
	for i := 'A'; i <= 'Z'; i++ {
		if i%2 == 1 {
			z01.PrintRune(i + 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')

	////case 7 : Converting without use of a string, use runes instead (combined cases)
	z01.PrintRune('7')
	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')

	////case 8 : Converting using string, mixed cases but inorder alphabetically
	z01.PrintRune('8')
	strx := "ABCDefghiJKlmnOPQRStuvwXYZ"
	for _, l := range strx {
		if (l >= 'A' && l <= 'Z') && l%2 == 1 {
			l = l + 32
			z01.PrintRune(l)
		} else if (l >= 'a' && l <= 'z') && l%2 == 1 {
			z01.PrintRune(l)
		} else if (l >= 'A' && l <= 'Z') && (l%2) == 0 {
			z01.PrintRune(l)
		} else {
			l = l - 32
			z01.PrintRune(l)
		}
	}
	z01.PrintRune('\n')

	////case 9 : Converting using string, mixed cases not in order alphabetically
	//changing the to the same case and append to a new slice of runes
	z01.PrintRune('9')
	strb := "QWErtyuioPASDfghjkLMNbvcxz"
	strc := make([]rune, len(strb))
	for _, i := range strb {
		if i >= 'A' && i <= 'Z' {
			i = i + 32
			strc = append(strc, i)
		} else {
			strc = append(strc, i)
		}
	}

	// sorting
	for i := 0; i < len(strc); i++ {
		for j := 0; j < len(strc)-1; j++ {
			if strc[j] > strc[i] {
				strc[j], strc[i] = strc[i], strc[j]
			}
		}
	}

	// even to caps odd to low then print
	for _, l := range strc {
		if (l >= 'A' && l <= 'Z') && l%2 == 1 {
			l = l + 32
			z01.PrintRune(l)
		} else if (l >= 'a' && l <= 'z') && l%2 == 1 {
			z01.PrintRune(l)
		} else if (l >= 'A' && l <= 'Z') && (l%2) == 0 {
			z01.PrintRune(l)
		} else {
			l = l - 32
			z01.PrintRune(l)
		}
	}
}

// main ends here
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r = r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r = r + ('a' - 'A')
	}
	return r
}
