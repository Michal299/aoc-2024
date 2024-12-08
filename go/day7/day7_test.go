package day7

import (
	"aoc-go/utils"
	"log"
	"testing"
)

var (
	resources = "../../resources/"
	day       = "day7"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1_example.txt")
		got := Part1(input)
		want := 3749

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
	t.Run("example part 2", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1_example.txt")
		got := Part2(input)
		want := 11387

		if got != want {
			t.Errorf("got: %d but want: %d", got, want)
		}
	})

	t.Run("actual part 2", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1.txt")
		got := Part2(input)
		log.Printf("Part2 - got: %d\n", got)
	})
}

func TestJoinOperation(t *testing.T) {
	b := 1234567890
	a := 1
	want := 11234567890
	got := JoinOperation(a, b)
	if want != got {
		t.Errorf("want: %d, got %d", want, got)
	}
}
