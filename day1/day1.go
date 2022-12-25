package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strconv"
	"strings"
)

const day = 1

func solvePartOne(input string) int {
	elfsAsString := strings.Split(input, "\n\n")

	highestCalories := 0
	for _, elfCaloriesString := range elfsAsString {
		elfCaloriesString = strings.TrimPrefix(elfCaloriesString, "\n")
		elfCaloriesString = strings.TrimSuffix(elfCaloriesString, "\n")
		caloriesAsStrings := strings.Split(elfCaloriesString, "\n")

		caloriesCarried := 0
		for _, calories := range caloriesAsStrings {
			caloriesAsInt, err := strconv.Atoi(calories)
			if nil != err {
				panic(err)
			}

			caloriesCarried += caloriesAsInt
		}

		if caloriesCarried > highestCalories {
			highestCalories = caloriesCarried
		}
	}

	return highestCalories
}

func solvePartTwo(input string) int {
	elfsAsString := strings.Split(input, "\n\n")

	highestCalories := 0
	secondHighestCalories := 0
	thirdHighestCalories := 0
	for _, elfCaloriesString := range elfsAsString {
		elfCaloriesString = strings.TrimPrefix(elfCaloriesString, "\n")
		elfCaloriesString = strings.TrimSuffix(elfCaloriesString, "\n")
		caloriesAsStrings := strings.Split(elfCaloriesString, "\n")

		caloriesCarried := 0
		for _, calories := range caloriesAsStrings {
			caloriesAsInt, err := strconv.Atoi(calories)
			if nil != err {
				panic(err)
			}

			caloriesCarried += caloriesAsInt
		}

		// Not clean but functional :)
		// Case 1: Higher than or equal highest calories
		if caloriesCarried >= highestCalories {
			if highestCalories > secondHighestCalories {
				if secondHighestCalories > thirdHighestCalories {
					thirdHighestCalories = secondHighestCalories
				}

				secondHighestCalories = highestCalories
			}

			highestCalories = caloriesCarried

			continue
		}

		// Case 2: Higher than or equal second-highest calories
		if caloriesCarried >= secondHighestCalories {
			if secondHighestCalories > thirdHighestCalories {
				thirdHighestCalories = secondHighestCalories
			}

			secondHighestCalories = caloriesCarried

			continue
		}

		// Case 3: Higher than third-highest calories
		if caloriesCarried > thirdHighestCalories {
			thirdHighestCalories = caloriesCarried
		}
	}

	return highestCalories + secondHighestCalories + thirdHighestCalories
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
