# Fibonacci Sequence Generator

This is a simple Go program that generates and prints the Fibonacci sequence up to a specified number.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [License](#license)

## Introduction

This program takes a single integer input from the command line and prints the Fibonacci sequence up to that number. The input must be a valid integer, and the program handles both positive and negative signs.

## Installation

1. Make sure you have [Go](https://golang.org/dl/) installed on your system.
2. Clone this repository or copy the source code into a local directory.

```bash
git clone https://github.com/Wambita/go-practice/fibonacci.git
cd fibonacci
```

## Usage

To run the program, use the `go run` command followed by the filename and the desired number as an argument:

```bash
go run main.go <number>
```

Replace `<number>` with the integer up to which you want the Fibonacci sequence.

### Example

```bash
go run main.go 10
```

This command will output:

```
0, 1, 1, 2, 3, 5, 8, 13, 21, 34
```

## Code Explanation

- **Main Function**:
  - Checks if exactly one argument is provided.
  - Validates if the argument is a digit.
  - Converts the argument to an integer.
  - Generates the Fibonacci sequence up to the given number and prints it.

- **`isdigit` Function**:
  - Verifies if the input string is a valid integer, allowing an optional leading `+` or `-`.

- **`fibonacci` Function**:
  - Recursively calculates the Fibonacci number for a given index.

- **`Seq` Function**:
  - Constructs and returns the Fibonacci sequence as a string.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
