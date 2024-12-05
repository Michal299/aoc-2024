package day5

import (
	"aoc-go/utils"
	"slices"
	"strconv"
	"strings"
)

type priority struct {
	before int
	after  int
}

type update []int

type priorityMap struct {
	indexes map[int]int
	matrix  [][]int
}

func Part1(input []string) int {
	prioritiesDefs, updates, uniqueNumbersCount := parseInput(input)

	sum := 0
	prioritiesGraph := preparePrioritiesMap(prioritiesDefs, uniqueNumbersCount)
	for _, update := range updates {
		if !isOrdered(prioritiesGraph, update) {
			continue
		}
		middleElement := update[len(update)/2]
		sum += middleElement
	}
	return sum
}

func Part2(input []string) int {
	prioritiesDefs, updates, uniqueNumbersCount := parseInput(input)

	sum := 0
	prioritiesGraph := preparePrioritiesMap(prioritiesDefs, uniqueNumbersCount)
	for _, update := range updates {
		if isOrdered(prioritiesGraph, update) {
			continue
		}
		sort(prioritiesGraph, update)
		middleElement := update[len(update)/2]
		sum += middleElement
	}
	return sum
}

func parseInput(input []string) (priorities []priority, updates []update, uniqueNumbersCount int) {
	addToPriority := true
	uniqueNumbers := make(map[int]bool)

	for _, line := range input {
		if len(line) == 0 {
			addToPriority = false
			continue
		}

		if addToPriority {
			parts := strings.Split(line, "|")
			base, _ := strconv.Atoi(parts[0])
			dominated, _ := strconv.Atoi(parts[1])
			priorities = append(priorities, priority{base, dominated})
			uniqueNumbers[dominated] = true
			uniqueNumbers[base] = true
		} else {
			update := utils.Map(strings.Split(line, ","), func(in string) (num int) {
				num, _ = strconv.Atoi(in)
				return
			})
			updates = append(updates, update)
		}
	}
	uniqueNumbersCount = len(uniqueNumbers)
	return
}

func preparePrioritiesMap(prioritiesDef []priority, totalNumberOfUnique int) priorityMap {
	indexes := make(map[int]int)
	priorities := make([][]int, totalNumberOfUnique)
	for i := 0; i < totalNumberOfUnique; i++ {
		priorities[i] = make([]int, totalNumberOfUnique)
	}

	for _, priorityDef := range prioritiesDef {

		a, found := indexes[priorityDef.before]
		if !found {
			a = len(indexes)
			indexes[priorityDef.before] = a
		}
		b, found := indexes[priorityDef.after]
		if !found {
			b = len(indexes)
			indexes[priorityDef.after] = b
		}

		priorities[a][b] = 1
		priorities[b][a] = -1
	}

	return priorityMap{indexes: indexes, matrix: priorities}
}

func isOrdered(priorities priorityMap, u update) bool {
	localPriorities := prepareLocalPrioritiesForUpdate(priorities, u)
	return slices.IsSortedFunc(u, getSortFunctionAccordingToPriorities(localPriorities))
}

func sort(priorities priorityMap, u update) {
	localPriorities := prepareLocalPrioritiesForUpdate(priorities, u)
	slices.SortFunc(u, getSortFunctionAccordingToPriorities(localPriorities))
}

func getSortFunctionAccordingToPriorities(p map[int]priority) func(a int, b int) int {
	return func(a int, b int) int {
		secondDiff := p[a].after - p[b].after
		firstDiff := p[b].before - p[a].before
		if firstDiff != 0 {
			return firstDiff
		}
		return secondDiff
	}
}

func prepareLocalPrioritiesForUpdate(priorities priorityMap, u update) (localPriorities map[int]priority) {
	localPriorities = make(map[int]priority)
	for _, elem := range u {
		p, found := localPriorities[elem]
		if !found {
			p = priority{0, 0}
		}
		for _, other := range u {
			if other == elem {
				continue
			}
			if priorities.matrix[priorities.indexes[elem]][priorities.indexes[other]] > 0 {
				p.before++
			}
			if priorities.matrix[priorities.indexes[elem]][priorities.indexes[other]] < 0 {
				p.after++
			}
		}
		localPriorities[elem] = p
	}
	return localPriorities
}
