package main

import (
	"fmt"
	"os"
)

// function to check for vald charatcters
func isValid(s string) bool {
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c == '-') {
			return false
		}
	}
	return true
}

// fuction to check if -h option is present in the string
func isHactive(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if rune(s[i]) == '-' && rune(s[i+1]) == 'h' {
			return true
		}
	}
	return false
}

// function to convert  []string to string
func toStr(arr []string) string {
	str := ""
	for _, s := range arr {
		str += s
	}
	return str
}

func main() {
	args := os.Args[1:]

	// options
	options := "options : abcdefghijklmnopqrstuvwxyz"

	// no args
	if len(args) == 0 {
		fmt.Println(options)
		return
	}

	// args to string conversion
	str := toStr(args)

	// check if  args are valid
	if !isValid(str) {
		fmt.Println("Invalid options")
		return
	}

	// check if  h is active
	if isHactive(str) {
		fmt.Println(options)
		return
	}

	// array of size 32 to rep each letter's active state
	bits := make([]int, 32)

	// map to track which chars have been processed
	seen := make(map[rune]bool)

	// loop through string to set corresponding bits for each option
	for _, char := range str {
		if !seen[char] && (char >= 'a' && char <= 'z') {
			// pos of char relative to a
			pos := int(char-'a') + 1
			// set corresponding bit n the array
			bits[pos-1] = 1
			// mark char as seen/active
			seen[char] = true
		}
	}

	//prepare final binary output string
	res := ""
	count := 0

	//loop throught the bits in reverse order to create binary output
	for i:=len(bits)-1; i >=0; i-- {
		//add a space btwn every 8bits except the last one
		if count == 8 &&  i != 0{
			res += "  "
			count = 0
		}
		//t is  append bit to binary output string based on whether it is set or not
		if bits[i] == 1 {
			res += "1"
		} else {
			res += "0"
		}
		count++
	}

	// print the final binary output string
	fmt.Println(res)
}
