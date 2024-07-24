package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
		fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	temp := []int{}
	if len(slice1) == 0 && len(slice2) != 0 {
		return Sort(slice2)
	}
	if len(slice2) == 0 && len(slice1) != 0 {
		return Sort(slice1)
	}
	if len(slice1) == len(slice2) {
		return equal(slice1, slice2)
	}
	if len(slice1) < len(slice2) {
		return greater(slice1, slice2)
	}
	if len(slice1) > len(slice2) {
		return less(slice1, slice2)
	}

	return temp
}

func Sort(s1 []int) []int {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1)-i-1; j++ {
			if s1[j+1] > s1[j] {
				s1[j], s1[j+1] = s1[j+1], s1[j]
			}
		}
	}
	return s1
}

func equal(s1, s2 []int) []int {
	sort1 := Sort(s1)
	sort2 := Sort(s2)
	temp := []int{}

	for i := 0; i < len(s1); i++ {
		temp = append(temp, sort1[i])
		temp = append(temp, sort2[i])
	}
	return temp
}

func greater(s1, s2 []int) []int {
	temp := []int{}
	for i := len(s2) - 1; i > (len(s2)-len(s1))-1; i-- {
		temp = append(temp, s2[i])
	}
	// fmt.Println(temp)
	so := s2[:len(s1)]
	sort1 := Sort(s1)
	sort2 := Sort(so)
	for i := 0; i < len(s1); i++ {
		temp = append(temp, sort1[i])
		temp = append(temp, sort2[i])
	}
	return temp
}

func less(s1, s2 []int) []int {
	temp := []int{}
	for i := len(s1) - 1; i >= (len(s1)-len(s2))-1; i-- {
		temp = append(temp, s1[i])
	}

	so := s1[:len(s2)]
	sort1 := Sort(so)
	sort2 := Sort(s2)
	for i := 0; i < len(s2); i++ {
		temp = append(temp, sort1[i])
		temp = append(temp, sort2[i])
	}
	return temp
}
