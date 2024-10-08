package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
} 

func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	// case 1 :nbrs== 0
	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]
	end := length

	// normal
	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// negatve indices
	if start < 0 {
		start = length + start
	}

	if end < 0 {
		end = end + length
	}
	if end == 0 && len(nbrs) > 1 {
		return nil
	}

	// within bounds
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}

	// start greater than length or end
	if start >= length || start >= end {
		return []string{}
	}

	return a[start:end]
}
