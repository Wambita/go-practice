package main

import "fmt"

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func Sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j+1] > a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	var temp []int

	len1 := len(slice1)
	len2 := len(slice2)
	minlen := len1

	if len1 == 0 && len2 != 0 {
		return Sort(slice2)
	}
	if len2 == 0 && len1 != 0 {
		return Sort(slice1)
	}

	if len2 < len1 {
		minlen = len2
	}

	//apppend ist
	if len1 > len2 {
		temp = append(temp, slice1[minlen:]...)
		slice1 = slice1[:minlen]
	} else if len2 > len1 {
		temp = append(temp, slice2[minlen:]...)
		slice2 = slice2[:minlen]
	}

	//sort
	slice1 = Sort(slice1)
	slice2 = Sort(slice2)

	for i := 0; i < minlen; i++ {
		temp = append(temp, slice1[i])
		temp = append(temp, slice2[i])
	}
	return temp
}
