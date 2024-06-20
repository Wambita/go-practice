package main

import (
	"fmt"
	"os"
)

func mainn() {
	count := 0
	if len(os.Args) > 1 {
		args := os.Args[1:]

		for i := 0; i < len(args); i++ {
			count++
		}
	} else {
		return
	}
	fmt.Println(count)
}
