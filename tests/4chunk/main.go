package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
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
		subslice := slice[start:end]
		result = append(result, subslice)

	}
	fmt.Println(result)
}
