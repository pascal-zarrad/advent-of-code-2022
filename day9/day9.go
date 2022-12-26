package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 9

const (
	MoveUp    = "U"
	MoveDown  = "D"
	MoveLeft  = "L"
	MoveRight = "R"
)

type Vector2D struct {
	X int
	Y int
}

func simulateTailMovement(head Vector2D, tail Vector2D) Vector2D {
	headTailDiffX := head.X - tail.X
	headTailDiffY := head.Y - tail.Y

	// Simple cases - horizontal & vertical
	switch {
	case headTailDiffX == 0 && headTailDiffY == 2: // Head is two above
		tail.Y++
	case headTailDiffX == 0 && headTailDiffY == -2: // Head is two down
		tail.Y--
	case headTailDiffX == 2 && headTailDiffY == 0: // Head is two right
		tail.X++
	case headTailDiffX == -2 && headTailDiffY == 0: // Head is two left
		tail.X--

	// Complex cases - diagonal move with up or down moved 2
	case headTailDiffX == 1 && headTailDiffY == 2: // Head is one right and two up
		tail.X++
		tail.Y++
	case headTailDiffX == -1 && headTailDiffY == 2: // Head is left and two up
		tail.X--
		tail.Y++
	case headTailDiffX == 1 && headTailDiffY == -2: // Head is one right and two down
		tail.X++
		tail.Y--
	case headTailDiffX == -1 && headTailDiffY == -2: // Head is one left and two down
		tail.X--
		tail.Y--

	// Complex cases - diagonal move with left or right moved 2
	case headTailDiffX == 2 && headTailDiffY == 1: // Head is two right and one up
		tail.X++
		tail.Y++
	case headTailDiffX == -2 && headTailDiffY == 1: // Head is two left and one up
		tail.X--
		tail.Y++
	case headTailDiffX == 2 && headTailDiffY == -1: // Head is two right and one down
		tail.X++
		tail.Y--
	case headTailDiffX == -2 && headTailDiffY == -1: // Head is two left and one down
		tail.X--
		tail.Y--

	// Complex cases - diagonal move with all moved 2
	case headTailDiffX == 2 && headTailDiffY == 2: // Head is up and right 2
		tail.X++
		tail.Y++
	case headTailDiffX == -2 && headTailDiffY == -2: // Head is up 2 and left 2
		tail.X--
		tail.Y--
	case headTailDiffX == -2 && headTailDiffY == 2: // Head is down 2 and right 2
		tail.X--
		tail.Y++
	case headTailDiffX == 2 && headTailDiffY == -2: // Head is up 2 and left 2
		tail.X++
		tail.Y--
	}

	return tail
}

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	commands := strings.Split(input, "\n")

	head := Vector2D{0, 0}
	tail := Vector2D{0, 0}

	recordedTailPositions := map[Vector2D]bool{}

	for _, cmd := range commands {
		action := string(cmd[0])
		count := util.StringToInt(cmd[2:])

		for i := 0; i < count; i++ {
			switch action {
			case MoveUp:
				head.Y++
			case MoveDown:
				head.Y--
			case MoveLeft:
				head.X--
			case MoveRight:
				head.X++
			}

			tail = simulateTailMovement(head, tail)

			recordedTailPositions[tail] = true
		}
	}

	return len(recordedTailPositions)
}

func solvePartTwo(input string) int {
	input = strings.Trim(input, "\n")
	commands := strings.Split(input, "\n")

	// Could have been done with an array,
	rope := make([]Vector2D, 10)

	recordedTailPositions := map[Vector2D]bool{}

	for _, cmd := range commands {
		action := string(cmd[0])
		count := util.StringToInt(cmd[2:])

		for i := 0; i < count; i++ {
			switch action {
			case MoveUp:
				rope[0].Y++
			case MoveDown:
				rope[0].Y--
			case MoveLeft:
				rope[0].X--
			case MoveRight:
				rope[0].X++
			}

			for c := 0; c < len(rope)-1; c++ {
				rope[c+1] = simulateTailMovement(rope[c], rope[c+1])
			}

			recordedTailPositions[rope[len(rope)-1]] = true
		}
	}

	return len(recordedTailPositions)
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
