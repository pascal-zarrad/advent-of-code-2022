package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 3

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	compartmentPairs := strings.Split(input, "\n")

	prioritySum := 0

	for _, compartmentPair := range compartmentPairs {

		half := len(compartmentPair) / 2

		firstCompartment := compartmentPair[:half]
		secondCompartment := compartmentPair[half:]

		for _, char := range firstCompartment {
			if strings.ContainsRune(secondCompartment, char) {
				// Uppercase char
				if char < 97 {
					prioritySum += int(char) - 64 + 26
					break
				}

				// Lowercase char
				prioritySum += int(char) - 96
				break
			}
		}
	}

	return prioritySum
}

func solvePartTwo(input string) int {
	input = strings.Trim(input, "\n")
	compartmentPairs := strings.Split(input, "\n")

	prioritySum := 0

	for i := 0; i < len(compartmentPairs); i += 3 {
		// Do not check if index exists, we just assume input is right
		bagA := compartmentPairs[i]
		bagB := compartmentPairs[i+1]
		bagC := compartmentPairs[i+2]

		for _, char := range bagA {
			if strings.ContainsRune(bagB, char) && strings.ContainsRune(bagC, char) {
				// Uppercase char
				if char < 97 {
					prioritySum += int(char) - 64 + 26
					break
				}

				// Lowercase char
				prioritySum += int(char) - 96
				break
			}
		}
	}

	return prioritySum
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
