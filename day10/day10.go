package main

import (
	"fmt"
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 10

func computeNewCycleSum(oldSum int, cycle int, currentX int) int {
	if cycle > 220 {
		return oldSum
	}

	if cycle != 20 && (cycle-20)%40 != 0 {
		return oldSum
	}

	return oldSum + (cycle * currentX)
}

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	commands := strings.Split(input, "\n")

	sum := 0
	cycleCount := 0
	x := 1

	for _, cmd := range commands {
		if cmd == "noop" {
			cycleCount++
			sum = computeNewCycleSum(sum, cycleCount, x)

			continue
		}

		if !strings.HasPrefix(cmd, "addx") {
			panic(fmt.Sprintf("Unexpected instruction received: %s", cmd))
		}

		add := util.StringToInt(cmd[5:])

		// Fist cycle count increase
		cycleCount++
		sum = computeNewCycleSum(sum, cycleCount, x)

		// Second cycle
		cycleCount++
		sum = computeNewCycleSum(sum, cycleCount, x)
		x += add
	}

	return sum
}

func solvePartTwo(input string) int {
	return 0
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
