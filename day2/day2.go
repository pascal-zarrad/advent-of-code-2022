package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"strings"
)

const day = 2

const (
	ActionOpponentRock     = "A"
	ActionOpponentPaper    = "B"
	ActionOpponentScissors = "C"

	ActionSelfRock     = "X"
	ActionSelfPaper    = "Y"
	ActionSelfScissors = "Z"

	ExpectedOutcomeWin   = "Z"
	ExpectedOutcomeDraw  = "Y"
	ExpectedOutcomeLoose = "X"

	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3

	ScoreWin  = 6
	ScoreDraw = 3
)

func solvePartOne(input string) int {
	gameRounds := strings.Split(input, "\n")

	finalScore := 0

	for _, round := range gameRounds {
		if round == "" {
			continue
		}

		actions := strings.Split(round, " ")

		opponentAction := actions[0]
		selfAction := actions[1]

		// Evaluate points for chosen actions
		switch selfAction {
		case ActionSelfRock:
			finalScore += ScoreRock
		case ActionSelfPaper:
			finalScore += ScorePaper
		case ActionSelfScissors:
			finalScore += ScoreScissors
		}

		switch {
		case opponentAction == ActionOpponentRock && selfAction == ActionSelfPaper:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentPaper && selfAction == ActionSelfScissors:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentScissors && selfAction == ActionSelfRock:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentRock && selfAction == ActionSelfRock:
			finalScore += ScoreDraw
		case opponentAction == ActionOpponentPaper && selfAction == ActionSelfPaper:
			finalScore += ScoreDraw
		case opponentAction == ActionOpponentScissors && selfAction == ActionSelfScissors:
			finalScore += ScoreDraw
		}
	}

	return finalScore
}

func solvePartTwo(input string) int {
	gameRounds := strings.Split(input, "\n")

	finalScore := 0

	for _, round := range gameRounds {
		if round == "" {
			continue
		}

		actions := strings.Split(round, " ")

		opponentAction := actions[0]
		selfAction := actions[1]

		// Compute own action based on expected result
		switch selfAction {
		case ExpectedOutcomeWin:
			switch opponentAction {
			case ActionOpponentRock:
				selfAction = ActionSelfPaper
			case ActionOpponentPaper:
				selfAction = ActionSelfScissors
			case ActionOpponentScissors:
				selfAction = ActionSelfRock
			}
		case ExpectedOutcomeDraw:
			switch opponentAction {
			case ActionOpponentRock:
				selfAction = ActionSelfRock
			case ActionOpponentPaper:
				selfAction = ActionSelfPaper
			case ActionOpponentScissors:
				selfAction = ActionSelfScissors
			}
		case ExpectedOutcomeLoose:
			switch opponentAction {
			case ActionOpponentRock:
				selfAction = ActionSelfScissors
			case ActionOpponentPaper:
				selfAction = ActionSelfRock
			case ActionOpponentScissors:
				selfAction = ActionSelfPaper
			}
		}

		// Evaluate points for chosen actions
		switch selfAction {
		case ActionSelfRock:
			finalScore += ScoreRock
		case ActionSelfPaper:
			finalScore += ScorePaper
		case ActionSelfScissors:
			finalScore += ScoreScissors
		}

		switch {
		case opponentAction == ActionOpponentRock && selfAction == ActionSelfPaper:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentPaper && selfAction == ActionSelfScissors:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentScissors && selfAction == ActionSelfRock:
			finalScore += ScoreWin
		case opponentAction == ActionOpponentRock && selfAction == ActionSelfRock:
			finalScore += ScoreDraw
		case opponentAction == ActionOpponentPaper && selfAction == ActionSelfPaper:
			finalScore += ScoreDraw
		case opponentAction == ActionOpponentScissors && selfAction == ActionSelfScissors:
			finalScore += ScoreDraw
		}
	}

	return finalScore
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
