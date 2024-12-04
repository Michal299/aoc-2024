package day4

import (
	"aoc-go/utils"
	"log"
	"testing"
)

var resources = "../../resources/"
var day = "day4"

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1_example.txt")
		got := Part1(input)
		want := 18

		if got != want {
			t.Errorf("got: %d but want: %d", got, want)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1.txt")
		got := Part1(input)
		log.Printf("Part1 - got: %d\n", got)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1_example.txt")
		got := Part2(input)
		want := 9

		if got != want {
			t.Errorf("got: %d but want: %d", got, want)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1.txt")
		got := Part2(input)
		log.Printf("Part2 - got: %d\n", got)
	})
}
