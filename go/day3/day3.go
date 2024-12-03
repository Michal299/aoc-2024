package day3

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	mulRegex     string = `mul\(\d+,\d+\)`
	enableReqex  string = `do\(\)`
	disableRegex string = `don't\(\)`
)

type functionCallType int

const (
	MUL functionCallType = iota
	DO
	DONT
)

type UnknownFunction error

func Part1(input []string) int {
	regex, _ := regexp.Compile(mulRegex)
	sum := 0
	for _, line := range input {
		matchedMulCalls := regex.FindAllString(line, -1)
		for _, mulCall := range matchedMulCalls {
			sum += calculateMulString(mulCall)
		}
	}
	return sum
}

func Part2(input []string) int {
	regex, _ := regexp.Compile(mulRegex + `|` + enableReqex + `|` + disableRegex)
	sum := 0
	enabled := true
	for _, line := range input {
		functionCalls := regex.FindAllString(line, -1)
		for _, functionCall := range functionCalls {

			callType, _ := getFunctionType(functionCall)
			switch callType {
			case MUL:
				if enabled {
					sum += calculateMulString(functionCall)
				}
			case DO:
				enabled = true
			case DONT:
				enabled = false
			}
		}
	}
	return sum
}

func calculateMulString(mulCall string) int {
	parts := strings.Split(mulCall, ",")
	left, _ := strings.CutPrefix(parts[0], "mul(")
	right, _ := strings.CutSuffix(parts[1], ")")

	leftValue, _ := strconv.Atoi(left)
	rightValue, _ := strconv.Atoi(right)
	return leftValue * rightValue
}

func getFunctionType(call string) (functionCallType, error) {
	doRegex, _ := regexp.Compile(enableReqex)
	dontRegex, _ := regexp.Compile(disableRegex)
	mulRegex, _ := regexp.Compile(mulRegex)

	if doRegex.MatchString(call) {
		return DO, nil
	}
	if dontRegex.MatchString(call) {
		return DONT, nil
	}
	if mulRegex.MatchString(call) {
		return MUL, nil
	}
	return -1, errors.New("unknown function")
}
