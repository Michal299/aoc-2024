package day4

import (
	"aoc-go/utils"
)

type coordinate struct {
	row    int
	column int
}

type direction coordinate

var neighboursDirections = []direction{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func Part1(input []string) int {
	runeTable := utils.Map(input, func(line string) []rune {
		return []rune(line)
	})
	return countSequences([]rune("XMAS"), runeTable)
}

func countSequences(seq []rune, input [][]rune) int {
	sum := 0

	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			sum += countSequencesFrom(seq, input, coordinate{row, column})
		}
	}
	return sum
}

func countSequencesFrom(seq []rune, input [][]rune, from coordinate) int {
	if seq[0] != input[from.row][from.column] {
		return 0
	}

	sum := 0
	for _, direction := range neighboursDirections {
		if findSequenceInDirectionFrom(seq, input, direction, from) {
			sum++
		}
	}
	return sum
}

func findSequenceInDirectionFrom(seq []rune, input [][]rune, d direction, from coordinate) bool {
	if from.row < 0 || from.row >= len(input) || from.column < 0 || from.column >= len(input[0]) {
		return false
	}

	if len(seq) <= 0 || seq[0] != input[from.row][from.column] {
		return false
	}
	if len(seq) == 1 && seq[0] == input[from.row][from.column] {
		return true
	}
	nextRow := from.row + d.row
	nextColumn := from.column + d.column
	if nextRow < 0 || nextRow >= len(input) || nextColumn < 0 || nextColumn >= len(input[0]) {
		return false
	}
	return findSequenceInDirectionFrom(seq[1:], input, d, coordinate{nextRow, nextColumn})
}

func Part2(input []string) int {
	runeTable := utils.Map(input, func(line string) []rune {
		return []rune(line)
	})

	sum := 0

	for row := 0; row < len(runeTable); row++ {
		for column := 0; column < len(runeTable[row]); column++ {
			if isXMas(runeTable, coordinate{row, column}) {
				sum++
			}
		}
	}
	return sum
}

func isXMas(input [][]rune, loc coordinate) bool {
	if input[loc.row][loc.column] != 'A' {
		return false
	}
	possibleDirections := []direction{
		{1, 1},   // to the bottom right
		{1, -1},  // to the bottom left
		{-1, -1}, // to the top left
		{-1, 1},  // to the top right
	}
	countFound := 0
	for _, d := range possibleDirections {
		startCoordinate := coordinate{loc.row - d.row, loc.column - d.column} // one step back according the direction to check
		if findSequenceInDirectionFrom([]rune("MAS"), input, d, startCoordinate) {
			countFound++
		}
	}
	return countFound >= 2
}
