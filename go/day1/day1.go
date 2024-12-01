package day1

import (
	"aoc-go/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	left, right := splitVertically(input)

	leftAsInts := utils.Map(left, mapToInt)
	slices.Sort(leftAsInts)
	rightAsInts := utils.Map(right, mapToInt)
	slices.Sort(rightAsInts)

	sumOfDistances := 0
	for i := 0; i < len(leftAsInts); i++ {
		sumOfDistances += (utils.Abs(leftAsInts[i] - rightAsInts[i]))
	}

	return sumOfDistances
}

func Part2(input []string) int {
	left, right := splitVertically(input)
	leftAsInts := utils.Map(left, mapToInt)
	rightAsInts := utils.Map(right, mapToInt)

	similarityScore := 0
	for _, elem := range leftAsInts {
		matching := utils.Filter(rightAsInts, func(x int) bool {
			return x == elem
		})

		similarityScore += elem * len(matching)
	}
	return similarityScore
}

func splitVertically(input []string) (left []string, right []string) {
	for _, line := range input {
		parts := strings.Split(line, " ")

		left = append(left, strings.TrimSpace(parts[0]))
		right = append(right, strings.TrimSpace(parts[len(parts)-1]))
	}
	return
}

func mapToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return
}
