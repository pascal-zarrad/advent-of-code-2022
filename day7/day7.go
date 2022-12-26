package main

import (
	"github.com/pascal-zarrad/advent-of-code-2022/util"
	"regexp"
	"strings"
)

const day = 7

func solvePartOne(input string) int {
	input = strings.Trim(input, "\n")
	cli := strings.Split(input, "\n")

	directoryRootPattern, err := regexp.Compile(`\$\scd\s/`)
	if nil != err {
		panic(err)
	}
	directoryUpPattern, err := regexp.Compile(`\$\scd\s\.\.`)
	if nil != err {
		panic(err)
	}
	directoryDownPattern, err := regexp.Compile(`\$\scd\s([a-bA-z]+)`)
	if nil != err {
		panic(err)
	}
	fileListPattern, err := regexp.Compile(`\$\sls`)
	if nil != err {
		panic(err)
	}
	outputPattern, err := regexp.Compile(`(.+)\s([a-zA-z]+)`)
	if nil != err {
		panic(err)
	}

	directoryLevel := 0
	directoryContentTotal := 0
	directorySubtotals := map[int]int{}
	for _, cmd := range cli {
		// Either handle command or view
		if strings.HasPrefix(cmd, "$") {
			if directoryRootPattern.MatchString(cmd) {
				directoryLevel = 0

				continue
			}

			if directoryUpPattern.MatchString(cmd) {
				if directorySubtotals[directoryLevel] <= 100000 {
					directoryContentTotal += directorySubtotals[directoryLevel]
				}

				directorySubtotals[directoryLevel-1] += directorySubtotals[directoryLevel]

				directoryLevel--

				continue
			}

			if directoryDownPattern.MatchString(cmd) {
				directoryLevel++

				continue
			}

			// Only ls command left, ignore it
			if fileListPattern.MatchString(cmd) {
				directorySubtotals[directoryLevel] = 0
			}

			continue
		}

		lsOut := outputPattern.FindAllStringSubmatch(cmd, -1)

		if 2 < len(lsOut) {
			continue
		}

		if lsOut[0][1] == "dir" {
			continue
		}

		directorySubtotals[directoryLevel] += util.StringToInt(lsOut[0][1])
	}

	if !directoryDownPattern.MatchString(cli[len(cli)-1]) && directorySubtotals[directoryLevel] <= 100000 {
		directoryContentTotal += directorySubtotals[directoryLevel]
	}

	return directoryContentTotal
}

func solvePartTwo(input string) int {
	input = strings.Trim(input, "\n")
	cli := strings.Split(input, "\n")

	maxDiskStorage := 70000000
	requiredDiskStorage := 30000000

	directoryRootPattern, err := regexp.Compile(`\$\scd\s/`)
	if nil != err {
		panic(err)
	}
	directoryUpPattern, err := regexp.Compile(`\$\scd\s\.\.`)
	if nil != err {
		panic(err)
	}
	directoryDownPattern, err := regexp.Compile(`\$\scd\s([a-bA-z]+)`)
	if nil != err {
		panic(err)
	}
	fileListPattern, err := regexp.Compile(`\$\sls`)
	if nil != err {
		panic(err)
	}
	outputPattern, err := regexp.Compile(`(.+)\s([a-zA-z]+)`)
	if nil != err {
		panic(err)
	}

	diskUsageSum := 0
	for _, cmd := range cli {
		if strings.HasPrefix(cmd, "$") {
			continue
		}

		if outputPattern.MatchString(cmd) {
			lsOut := outputPattern.FindAllStringSubmatch(cmd, -1)
			if 2 < len(lsOut) {
				continue
			}
			if lsOut[0][1] == "dir" {
				continue
			}

			diskUsageSum += util.StringToInt(lsOut[0][1])
		}
	}

	diskSpaceFree := maxDiskStorage - diskUsageSum
	minSpaceRequired := requiredDiskStorage - diskSpaceFree

	directoryLevel := 0
	directorySubtotals := map[int]int{}
	actualBestSpace := maxDiskStorage
	for _, cmd := range cli {
		// Either handle command or view
		if strings.HasPrefix(cmd, "$") {
			if directoryRootPattern.MatchString(cmd) {
				directoryLevel = 0

				continue
			}

			if directoryUpPattern.MatchString(cmd) {
				if directorySubtotals[directoryLevel] >= minSpaceRequired && directorySubtotals[directoryLevel] < actualBestSpace {
					actualBestSpace = directorySubtotals[directoryLevel]
				}

				directorySubtotals[directoryLevel-1] += directorySubtotals[directoryLevel]

				directoryLevel--

				continue
			}

			if directoryDownPattern.MatchString(cmd) {
				directoryLevel++

				continue
			}

			// Only ls command left, ignore it
			if fileListPattern.MatchString(cmd) {
				directorySubtotals[directoryLevel] = 0
			}

			continue
		}

		lsOut := outputPattern.FindAllStringSubmatch(cmd, -1)

		if 2 < len(lsOut) {
			continue
		}

		if lsOut[0][1] == "dir" {
			continue
		}

		directorySubtotals[directoryLevel] += util.StringToInt(lsOut[0][1])
	}

	if !directoryDownPattern.MatchString(cli[len(cli)-1]) && directorySubtotals[directoryLevel] >= minSpaceRequired && directorySubtotals[directoryLevel] < actualBestSpace {
		actualBestSpace = directorySubtotals[directoryLevel]
	}

	return actualBestSpace
}

func main() {
	input := util.ReadInput(day)

	util.PrintResult(1, solvePartOne(input))
	util.PrintResult(2, solvePartTwo(input))
}
