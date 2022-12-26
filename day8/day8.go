package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 8

// Anything but a clean solution
// But it does the job

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])

	// Convert entire grid to 2d array consisting of integers
	grid := make([][]int, height)
	for key, line := range lines {
		grid[key] = make([]int, width)

		for i, char := range line {
			grid[key][i] = util.StringToInt(string(char))
		}
	}

	// Every tree only counts once
	visible := make([][]bool, height)
	for key := range lines {
		visible[key] = make([]bool, width)
	}

	visibleCount := (2 * height) + (width*2 - 4)

	// Left to right
	for i := 1; i < len(grid)-1; i++ {
		highestValue := grid[i][0]
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] > highestValue {
				highestValue = grid[i][j]

				if !visible[i][j] {
					visible[i][j] = true
					visibleCount++
				}
			}
		}
	}

	// Right to left
	for i := 1; i < len(grid)-1; i++ {
		highestValue := grid[i][len(grid[i])-1]
		for j := len(grid[i]) - 2; j >= 1; j-- {
			if grid[i][j] > highestValue {
				highestValue = grid[i][j]

				if !visible[i][j] {
					visible[i][j] = true
					visibleCount++
				}
			}
		}
	}

	// Top to bottom
	for j := 1; j < len(grid[0])-1; j++ {
		highestValue := grid[0][j]
		for i := 1; i < len(grid)-1; i++ {
			if grid[i][j] > highestValue {
				highestValue = grid[i][j]

				if !visible[i][j] {
					visible[i][j] = true
					visibleCount++
				}
			}
		}
	}

	// Bottom to top
	for j := len(grid[0]) - 2; j >= 1; j-- {
		highestValue := grid[len(grid)-1][j]
		for i := len(grid) - 1; i >= 1; i-- {
			if grid[i][j] > highestValue {
				highestValue = grid[i][j]

				if !visible[i][j] {
					visible[i][j] = true
					visibleCount++
				}
			}
		}
	}

	return visibleCount
}

func solvePartTwo(input string) int {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	highestScenicScore := 0

	// Convert entire grid to 2d array consisting of integers
	grid := make([][]int, len(lines))
	for key, line := range lines {
		grid[key] = make([]int, len(line))

		for i, char := range line {
			grid[key][i] = util.StringToInt(string(char))
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			current := grid[i][j]
			colLeft, colRight, colUp, colDown := -1, -1, -1, -1

			steps := 1
			for colLeft == -1 || colRight == -1 || colUp == -1 || colDown == -1 {
				// Left
				if colLeft == -1 {
					iL := j - steps

					switch {
					case iL < 0:
						colLeft = steps - 1
					default:
						if grid[i][iL] >= current {
							colLeft = steps
						}
					}
				}

				// Right
				if colRight == -1 {
					iR := j + steps

					switch {
					case iR > len(grid[i])-1:
						colRight = steps - 1
					default:
						if grid[i][iR] >= current {
							colRight = steps
						}
					}
				}

				// Up
				if colUp == -1 {
					iU := i - steps

					switch {
					case iU < 0:
						colUp = steps - 1
					default:
						if grid[iU][j] >= current {
							colUp = steps
						}
					}
				}

				// Down
				if colDown == -1 {
					iD := i + steps

					switch {
					case iD > len(grid)-1:
						colDown = steps - 1
					default:
						if grid[iD][j] >= current {
							colDown = steps
						}
					}
				}

				steps++
			}

			scenicScore := colLeft * colRight * colUp * colDown

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
