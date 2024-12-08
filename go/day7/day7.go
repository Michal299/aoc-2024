package day7

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type equation struct {
	targetValue int
	numbers     []int
}

type opertator int

const (
	mul opertator = iota
	add
	join
)

func (e *equation) isValid2(op []opertator) bool {

	getOperatorIndex := func(combination, numberOfOperators, position int) (i int) {
		if position == 0 {
			i = combination % numberOfOperators
			return
		}
		i = (combination / pow(numberOfOperators, position)) % numberOfOperators
		return
	}

	currentCombination := 0
	totalNumberOfPermutations := pow(len(op), len(e.numbers)-1)
	for i := 0; i < totalNumberOfPermutations; i++ {
		currentValue := e.numbers[0]
		for index, num := range e.numbers[1:] {
			operatorIndex := getOperatorIndex(currentCombination, len(op), index)
			switch op[operatorIndex] {
			case mul:
				currentValue = currentValue * num
			case add:
				currentValue = currentValue + num
			case join:
				currentValue = JoinOperation(currentValue, num)
			}

			if currentValue > e.targetValue {
				break
			}
		}

		if currentValue == e.targetValue {
			return true
		}
		currentCombination++
	}
	return false
}
func Part1(input []string) int {
	data := utils.Map(input, mapLineToEquation)
	sum := 0
	for _, eq := range data {
		if eq.isValid2([]opertator{add, mul}) {
			sum += eq.targetValue
		}
	}
	return sum
}

func Part2(input []string) int {
	data := utils.Map(input, mapLineToEquation)
	sum := 0
	for _, eq := range data {
		if eq.isValid2([]opertator{add, mul, join}) {
			sum += eq.targetValue
		}
	}
	return sum
}

func mapLineToEquation(line string) equation {
	parts := strings.Split(line, ":")
	targetValue, err := strconv.Atoi(parts[0])
	rawNumbers := strings.Trim(parts[1], " ")
	if err != nil {
		panic(fmt.Sprintf("Failure when processing an target value from input: %q", line))
	}
	numbers := utils.Map(strings.Split(rawNumbers, " "), func(val string) int {
		v, e := strconv.Atoi(val)
		if e != nil {
			panic(fmt.Sprintf("Failure when processing the numbers: %q", val))
		}
		return v
	})
	return equation{targetValue, numbers}
}

func pow(base, exp int) int {
	r := 1
	for i := 0; i < exp; i++ {
		r = r * base
	}
	return r
}

func JoinOperation(a, b int) (r int) {
	if b < 10 {
		r = a*10 + b
		return
	}
	return JoinOperation(a, b/10)*10 + b%10
}
