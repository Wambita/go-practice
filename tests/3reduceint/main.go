package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

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
	dif := func(acc int, cur int) int {
		return acc - cur
	}
	as := []int{600, 500, 3444, 909}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
	ReduceInt(as, dif)
}

func ReduceInt(a []int, f func(int, int) int) {
	// len of slice
	if len(a) == 0 {
		return
	}
	result := a[0]
	for i := 0; i < len(a)-1; i++ {
		result = f(result, a[i+1])
	}
	result2 := strconv.Itoa(result)
	for _, num := range result2 {
		z01.PrintRune(num)
	}
	z01.PrintRune('\n')
}
