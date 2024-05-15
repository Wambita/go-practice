package main

import (
	"fmt"
	"os"
)

// Write a program that displays its first argument, if there is one.
func Display() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
}

func mainc() {
	Display()
}
