package main

import (
	"log"
	"reflect"
)

var testCases = []struct {
	args [][]int
	want []int
}{
	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[]int{3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{4, 5, 6, 7, 8, 9}, {1, 2, 3}},
		[]int{9, 8, 7, 6, 3, 5, 2, 4, 1},
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}},
		[]int{9, 8, 7, 3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{1, 2, 3}, {4, 5}},
		[]int{3, 2, 5, 1, 4},
	},
	{
		[][]int{{}, {4, 5, 6}},
		[]int{6, 5, 4},
	},
	{
		[][]int{{1, 2, 3}, {}},
		[]int{3, 2, 1},
	},
	{
		[][]int{{}, {}},

		[]int{},
	},
	{
		[][]int{{1, 2, 4}, {10, 20, 30, 40, 50}},

		[]int{50, 40, 4, 30, 2, 20, 1, 10},
	},
}

func main() {
	for _, tc := range testCases {
		got := RevConcatAlternate(tc.args[0], tc.args[1])
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%#v, %#v) == %#v instead of %#v\n",
				"RevConcatAlternate",
				tc.args[0],
				tc.args[1],
				got,
				tc.want,
			)
		}
	}
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	var temp []int

	len1 := len(slice1)
	len2 := len(slice2)
	minlen := len1

	if len1 == 0 && len2 == 0 {
		return []int{}
	}
	if len1 == 0 && len2 != 0 {
		for i := len2 - 1; i >= 0; i-- {
			temp = append(temp, slice2[i])
		}
		return temp
	}
	if len2 == 0 && len1 != 0 {
		for i := len1 - 1; i >= 0; i-- {
			temp = append(temp, slice1[i])
		}
		return temp
	}

	if len2 < len1 {
		minlen = len2
	}

	// apppend list
	if len1 > len2 {
		for i := len1 - 1; i >= minlen; i-- {
			temp = append(temp, slice1[i])
		}
		slice1 = slice1[:minlen]
	} else if len2 > len1 {
		for i := len2 - 1; i >= minlen; i-- {
			temp = append(temp, slice2[i])
		}
	}

	// sort
	slice1 = (slice1)
	slice2 = (slice2)

	for i := minlen - 1; i >= 0; i-- {
		temp = append(temp, slice1[i], slice2[i])
	}
	return temp
}
