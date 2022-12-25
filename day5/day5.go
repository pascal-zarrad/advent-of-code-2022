package main

import (
	"fmt"
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"regexp"
	"strings"
)

const day = 5

func solvePartOne(input string) string {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	layout := map[int][]string{}
	commands := make([]string, 0)
	placeCount := 1

	containerRegex, err := regexp.Compile(`[\s|\[](\s|[A-Z])[\s|\]]{1,2}`)
	if nil != err {
		panic(err)
	}

	// First split commands from layout
	tempLayoutLines := make([]string, 0)
	collectingLayout := true
	for _, line := range lines {
		if line == "" {
			collectingLayout = false
			continue
		}

		if collectingLayout {
			tempLayoutLines = append(tempLayoutLines, line)
			continue
		}

		commands = append(commands, line)
	}

	// Now collect how many slots we have
	for _, char := range tempLayoutLines[len(tempLayoutLines)-1] {
		if char == ' ' {
			continue
		}

		layout[placeCount] = make([]string, 0)
		placeCount++
	}

	// Finally, fill in containers into slots
	for i := len(tempLayoutLines) - 2; i >= 0; i-- {
		containersSM := containerRegex.FindAllStringSubmatch(tempLayoutLines[i], -1)
		containers := make([]string, 0)
		for _, match := range containersSM {
			containers = append(containers, match[1])
		}

		for j := 1; j <= len(layout); j++ {
			if j <= len(containers) && containers[j-1] != " " {
				layout[j] = append(layout[j], containers[j-1])
			}
		}
	}

	// Apply crane actions
	moveRegex, err := regexp.Compile(`\d+`)
	if nil != err {
		panic(err)
	}

	for _, command := range commands {
		actionNumbers := moveRegex.FindAllString(command, 3)

		amount := util.StringToInt(actionNumbers[0])
		from := util.StringToInt(actionNumbers[1])
		to := util.StringToInt(actionNumbers[2])

		for i := 0; i < amount; i++ {
			element := layout[from][len(layout[from])-1]
			layout[from] = layout[from][:len(layout[from])-1]
			layout[to] = append(layout[to], element)
		}
	}

	result := ""
	for i := 1; i < len(layout)+1; i++ {
		result = fmt.Sprintf("%s%s", result, layout[i][len(layout[i])-1])
	}

	return result
}

func solvePartTwo(input string) string {
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")

	layout := map[int][]string{}
	commands := make([]string, 0)
	placeCount := 1

	containerRegex, err := regexp.Compile(`[\s|\[](\s|[A-Z])[\s|\]]{1,2}`)
	if nil != err {
		panic(err)
	}

	// First split commands from layout
	tempLayoutLines := make([]string, 0)
	collectingLayout := true
	for _, line := range lines {
		if line == "" {
			collectingLayout = false
			continue
		}

		if collectingLayout {
			tempLayoutLines = append(tempLayoutLines, line)
			continue
		}

		commands = append(commands, line)
	}

	// Now collect how many slots we have
	for _, char := range tempLayoutLines[len(tempLayoutLines)-1] {
		if char == ' ' {
			continue
		}

		layout[placeCount] = make([]string, 0)
		placeCount++
	}

	// Finally, fill in containers into slots
	for i := len(tempLayoutLines) - 2; i >= 0; i-- {
		containersSM := containerRegex.FindAllStringSubmatch(tempLayoutLines[i], -1)
		containers := make([]string, 0)
		for _, match := range containersSM {
			containers = append(containers, match[1])
		}

		for j := 1; j <= len(layout); j++ {
			if j <= len(containers) && containers[j-1] != " " {
				layout[j] = append(layout[j], containers[j-1])
			}
		}
	}

	// Apply crane actions
	moveRegex, err := regexp.Compile(`\d+`)
	if nil != err {
		panic(err)
	}

	for _, command := range commands {
		actionNumbers := moveRegex.FindAllString(command, 3)

		amount := util.StringToInt(actionNumbers[0])
		from := util.StringToInt(actionNumbers[1])
		to := util.StringToInt(actionNumbers[2])

		layout[to] = append(layout[to], layout[from][len(layout[from])-amount:]...)
		layout[from] = layout[from][:len(layout[from])-amount]
	}

	result := ""
	for i := 1; i < len(layout)+1; i++ {
		result = fmt.Sprintf("%s%s", result, layout[i][len(layout[i])-1])
	}

	return result
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
