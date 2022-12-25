package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
)

const day = 6

func solvePartOne(input string) int {
	for key := range input {
		if key+3 > len(input) {
			break
		}

		sequence := input[key : key+4]

		unique := true
		for k, v := range sequence {
			runeCount := 0

			for _, c := range sequence[k:] {
				if c == v {
					runeCount++
				}
			}

			if runeCount != 1 {
				unique = false
				break
			}
		}

		if unique == true {
			return key + 4
		}
	}

	return 0
}

func solvePartTwo(input string) int {
	for key := range input {
		if key+14 > len(input) {
			break
		}

		sequence := input[key : key+14]

		unique := true
		for k, v := range sequence {
			runeCount := 0

			for _, c := range sequence[k:] {
				if c == v {
					runeCount++
				}
			}

			if runeCount != 1 {
				unique = false
				break
			}
		}

		if unique == true {
			return key + 14
		}
	}

	return 0
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
