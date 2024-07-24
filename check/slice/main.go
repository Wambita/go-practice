package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	var slice = []string{}
	//length of strings is one
	//negative num
	if len(nbrs) == 1 {
		if nbrs[0] < 1 {
			nbrs[0] = -nbrs[0]
			pos := len(a) - nbrs[0]
			a = a[pos:]
			return a
		} else {
			a = a[nbrs[0]:]
			return a
		}

	}

	if len(nbrs) > 1 {
		start := nbrs[0]
		finish := nbrs[len(nbrs)-1]

		if start == 0 || finish == 0 {
			return nil
		}

		if start < 0 {
			starts := -start
			start = len(a) - starts
		}
		if finish < 0 {
			finishes := -finish
			finish = len(a) - finishes
		}
		a = a[start:finish]
		return a
	}

	return slice
}
