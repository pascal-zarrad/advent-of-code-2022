package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 4

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	assignmentPairs := strings.Split(input, "\n")

	assignmentCollisionCount := 0

	for _, assignmentPair := range assignmentPairs {
		assignments := strings.Split(assignmentPair, ",")
		assignmentA := strings.Split(assignments[0], "-")
		assignmentB := strings.Split(assignments[1], "-")

		aLowBound := util.StringToInt(assignmentA[0])
		aHighBound := util.StringToInt(assignmentA[1])

		bLowBound := util.StringToInt(assignmentB[0])
		bHighBound := util.StringToInt(assignmentB[1])

		switch {
		case aLowBound >= bLowBound && aHighBound <= bHighBound:
			assignmentCollisionCount++
		case bLowBound >= aLowBound && bHighBound <= aHighBound:
			assignmentCollisionCount++
		}
	}

	return assignmentCollisionCount
}

func solvePartTwo(input string) int {
	input = strings.Trim(input, "\n")
	assignmentPairs := strings.Split(input, "\n")

	assignmentCollisionCount := 0

	for _, assignmentPair := range assignmentPairs {
		assignments := strings.Split(assignmentPair, ",")
		assignmentA := strings.Split(assignments[0], "-")
		assignmentB := strings.Split(assignments[1], "-")

		aLowBound := util.StringToInt(assignmentA[0])
		aHighBound := util.StringToInt(assignmentA[1])

		bLowBound := util.StringToInt(assignmentB[0])
		bHighBound := util.StringToInt(assignmentB[1])

		switch {
		case aLowBound >= bLowBound && aLowBound <= bHighBound:
			assignmentCollisionCount++
		case aHighBound <= bHighBound && aHighBound >= bHighBound:
			assignmentCollisionCount++
		case bLowBound >= aLowBound && bLowBound <= aHighBound:
			assignmentCollisionCount++
		case bHighBound <= aHighBound && bHighBound >= aHighBound:
			assignmentCollisionCount++
		}
	}

	return assignmentCollisionCount
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
