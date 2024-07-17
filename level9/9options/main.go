package main

import (
	"os"
)

// isValidOption checks if a given byte is a valid option character
func isValidOption(option byte) bool {
	return 'a' <= option && option <= 'z'
}

// writeString writes a string to os.Stdout
func writeString(s string) {
	for i := 0; i < len(s); i++ {
		os.Stdout.Write([]byte{s[i]})
	}
}

// writeBinary formats and prints the bitmask as required
func writeBinary(bitmask uint32) {
	result := ""
	// each byte has length of 8
	// Iterate over each byte (4 bytes in total for uint32)
	for i := 0; i < 4; i++ {
		// Extract the current byte from the bitmask and format it as binary
		byteVal := byte((bitmask >> (8 * (3 - i))) & 0xFF) // Get the i-th byte from left to right
		// Format the byte as an 8-bit binary string
		for j := 0; j < 8; j++ {
			if byteVal&(1<<(7-j)) != 0 {
				result += "1"
			} else {
				result += "0"
			}
		}
		result += " " // Separate each byte with a space
	}
	result += "\n" // Add newline at the end of the binary representation
	writeString(result)
}

func main() {
	args := os.Args[1:] // Get command line arguments

	// Check if no arguments or if -h flag is present
	if len(args) == 0 || (len(args) > 0 && args[0] == "-h") {
		writeString("options: abcdefghijklmnopqrstuvwxyz\n") // Print help message and exit
		return
	}

	var bitmask uint32 // Initialize bitmask to store options as bits

	// Process each argument
	for _, arg := range args {
		// Check if argument starts with '-' and has valid options
		if len(arg) > 1 && arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				option := arg[i] // Get each option character
				if !isValidOption(option) {
					writeString("Invalid Option\n") // Print error for invalid option
					return
				}
				// Set the corresponding bit in bitmask for the option character
				bitmask += 1 << (option - 'a') // Set the bit corresponding to option
			}
		} else {
			writeString("Invalid Option\n") // Print error for invalid argument
			return
		}
	}

	writeBinary(bitmask) // Print the bitmask in binary format
}

// functions
// 1.writeString
// 2.validoptions
// 3.writeBinary
// 4.Main
