package day9

import (
	"aoc-go/utils"
	"log"
	"testing"
)

var (
	resources = "../../resources/"
	day       = "day9"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1_example.txt")
		got := Part1(input)
		want := 1928

		if got != want {
			t.Errorf("got: %d but want: %d", got, want)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := utils.ReadInput(resources + day + "/part1.txt")
		got := Part1(input)
		log.Printf("PartX - got: %d\n", got)
	})
}
