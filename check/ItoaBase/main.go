package main

import (
	"fmt"
	"math"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	fmt.Println(ItoaBase(255, 16))  // "FF"
	fmt.Println(ItoaBase(-255, 16)) // "-FF"
	fmt.Println(ItoaBase(10, 2))    // "1010"
	fmt.Println(ItoaBase(42, 4))    // "222"
	fmt.Println(ItoaBase(0, 10))    // "0"

	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MaxInt, base)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MinInt, base)
	}
}

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return "wrong base"
	}

	result := ""

	if value < 0 {
		if value == math.MinInt {
			return  "-"  + ItoaBase(-(value/base), base) + string("0123456789ABCDEF"[-(value%base)])
		}
		return "-" + ItoaBase(-value, base)
	}

	s := "0123456789ABCDEF"

	for value > 0 {
		rem := value % base
		result = string(s[rem]) + result
		value /= base
	}

	if result == "" {
		return "0"
	}
	return result
}
