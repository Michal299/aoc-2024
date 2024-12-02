package day2

import (
	"aoc-go/utils"
	"strconv"
	"strings"
)

type report []int

func Part1(input []string) int {
	reports := utils.Map(input, mapLineToReport)
	return len(utils.Filter(reports, func(r report) bool {
		return r.isSafe(0)
	}))
}

func Part2(input []string) int {
	reports := utils.Map(input, mapLineToReport)
	return len(utils.Filter(reports, func(r report) bool {
		return r.isSafe(1)
	}))
}

func mapLineToReport(line string) report {
	return utils.Map(strings.Split(line, " "), func(s string) (val int) {
		val, _ = strconv.Atoi(s)
		return val
	})
}

func (r report) isSafe(levelErrorTolerance int) bool {
	var current, predecessor, diff int
	diff = 0
	for index := 1; index < len(r); index++ {
		current = r[index]
		predecessor = r[index-1]

		currentDiff := current - predecessor
		if utils.Abs(currentDiff) > 3 || utils.Abs(currentDiff) < 1 || diff*currentDiff < 0 {
			if levelErrorTolerance <= 0 {
				return false
			} else {
				subreportLeft := removeLevelAt(r, index)
				subreportRight := removeLevelAt(r, index-1)
				return subreportLeft.isSafe(levelErrorTolerance-1) ||
					subreportRight.isSafe(levelErrorTolerance-1) ||
					(index-2 >= 0 && removeLevelAt(r, index-2).isSafe(levelErrorTolerance-1))
			}
		}
		diff = currentDiff
	}
	return true
}

func removeLevelAt(r report, index int) (result report) {
	result = append([]int{}, r[:index]...)
	result = append(result, r[index+1:]...)
	return
}
