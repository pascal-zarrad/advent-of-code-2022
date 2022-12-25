package util

import (
	"fmt"
	"os"
	"strconv"
)

// Some helper functions to speed up challenge progress

// ReadInput reads the challenge file from filesystem
func ReadInput(part int) string {
	content, err := os.ReadFile(fmt.Sprintf("day%d/day%d.txt", part, part))

	if nil != err {
		panic(fmt.Sprintf("Missing day%d.txt input file! Error: %v", part, err))
	}

	return string(content)
}

// PrintResult prints the result preformatted to the console.
func PrintResult[T any](part int, result T) {
	fmt.Println(fmt.Sprintf("Part %d result is: %v", part, result))
}

// StringToInt converts the passed string to an int.
// If the string cannot be converted, a panic will be thrown.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if nil != err {
		panic(err)
	}

	return i
}
