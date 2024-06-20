package main

// Write a function Max that will return the maximum value in a slice of integers. If the slice is empty it will return 0.
import (
	"fmt"
)

func mainnn() {
	a := []int{11, 67, 98, 567, 9876, 6, 4, 8, 653322, 56788999}
	max := Max(a)
	fmt.Println(max)
	maxs2 := max2(a)
	fmt.Println(maxs2)
}

// attempt 1
func Max(a []int) int {
	if len(a) > 1 {
		for i := 0; i < len(a)-1; i++ {
			for j := 1; j < len(a); j++ {
				if a[i] < a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
		return a[0]
	} else {
		return 0
	}
}

// attempt 2
func max2(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		max := 0
		for _, i := range a {
			if max < i {
				max = i
			}
		}
		return max
	}
}
