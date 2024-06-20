package main

import "github.com/01-edu/z01"

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{100, 200, 300, 400, 500, 111, 78}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)

	ReduceInt2(as, mul)
	ReduceInt2(as, sum)
	ReduceInt2(as, div)
}

// attempt 1
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	result := a[0]

	for i := 0; i < len(a)-1; i++ {
		result = f(result, a[i+1])
	}
	PrintDigits(result)
	z01.PrintRune('\n')
}

// attempt2
func ReduceInt2(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	result := a[0]

	for _, value := range a[1:] {
		result = f(result, value)
	}

	PrintDigits(result)
	z01.PrintRune('\n')
}

func PrintDigits(n int) {
	if n < 10 {
		z01.PrintRune('0' + rune(n))
		return
	}

	PrintDigits(n / 10)
	z01.PrintRune('0' + rune(n%10))
}
