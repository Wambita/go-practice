package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	ok := ispowof2(num)
	fmt.Println(ok)
}

func ispowof2(n int) bool {
	return (n&(n-1) == 0) && n > 0
}
