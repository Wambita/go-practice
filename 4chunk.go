package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

/*
Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

	If the size is 0 it should print a newline ('\n').

# Expected function

func Chunk(slice []int, size int) {

}

# Usage

Here is a possible program to test your function :

package main

	func main() {
		Chunk([]int{}, 10)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
		Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	}

And its output :

$ go run .
[]

[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
$
*/
func maicmmn() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}
	var result [][]int

	for start := 0; start < len(slice); start += size {
		end := start + size
		if end > len(slice) {
			end = len(slice)
		}
		subSlice := slice[start:end]
		result = append(result, subSlice)
	}
	fmt.Println(result)
}
