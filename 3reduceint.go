package main

import "fmt"

// The function should have as parameters a slice of integers a []int and a function f func(int, int) int.
// For each element of the slice, it should apply the function f func(int, int) int, save the result and then print it.
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	result := a[0]
	for i := 0; i < len(a)-1; i++ {
		result = f(result, a[i+1])
	}
	fmt.Println(result)
}

func mainx() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
