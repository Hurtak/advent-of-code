package main

import (
	"testing"

	"github.com/hurtak/advent-of-code-2020/utils"
)

type Fixtures struct {
	Input    []int
	Expected int
}

func TestPart1(t *testing.T) {
	fixtures := []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 5145790,
		},
	}
	for _, fixture := range fixtures {
		utils.AssertEqual(t, Part1(fixture.Input), fixture.Expected)
	}
}

func TestPart2(t *testing.T) {
	fixtures := []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 241861950,
		},
	}
	for _, fixture := range fixtures {
		got := Part2(fixture.Input)
		if got != fixture.Expected {
			t.Errorf("Part2(%d) = %d; want %d", fixture.Input, got, fixture.Expected)
		}
	}
}
