// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"testing"
// )

// func Test_validateArgs(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		args          []string
// 		expectedFile  string
// 		expectedFlag  string
// 		expectedInput string
// 		expectedError string // Expected error message to be printed
// 	}{
// 		{
// 			name:          "Invalid case with too few arguments",
// 			args:          []string{""},
// 			expectedFile:  "", // Assuming default behavior
// 			expectedFlag:  "",
// 			expectedInput: "",
// 			expectedError: "error: insufficient arguments",
// 		},
// 		{
// 			name:          "Invalid case with unrecognized flag",
// 			args:          []string{"-invalid_flag", "user_input"},
// 			expectedFile:  "", // Assuming default behavior
// 			expectedFlag:  "",
// 			expectedInput: "",
// 			expectedError: "error: unrecognized flag: -invalid_flag",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Redirect stdout to capture printed messages
// 			old := os.Stdout
// 			r, w, _ := os.Pipe()
// 			os.Stdout = w

// 			defer func() {
// 				os.Stdout = old
// 			}()

// 			// Call validateArgs with test args
// 			validateArgs(tt.args)

// 			// Capture the printed output
// 			w.Close()
// 			capturedOutput, _ := captureOutput(r)

// 			// Check if the expected error message is present in captured output
// 			if !strings.Contains(capturedOutput, tt.expectedError) {
// 				t.Errorf("validateArgs() did not print expected error message. Expected: '%s', Got: '%s'", tt.expectedError, capturedOutput)
// 			}
// 		})
// 	}
// }

// // Helper function to capture stdout
//
//	func captureOutput(f *os.File) (string, error) {
//		var output strings.Builder
//		_, err := fmt.Fprint(&output, f)
//		return output.String(), err
//	}
package main
