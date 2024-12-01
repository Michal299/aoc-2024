package daytemplate

import (
	"aoc-go/utils"
	"log"
	"testing"
)

func TestPartX(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := utils.ReadInput("../resources/dayX/partX_example.txt")
		got := Part1(input)
		want := 0

		if got != want {
			t.Errorf("got: %d but want: %d", got, want)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := utils.ReadInput("../resources/dayX/partX.txt")
		got := Part1(input)
		log.Printf("PartX - got: %d\n", got)
	})
}
