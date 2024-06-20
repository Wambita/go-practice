package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

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
