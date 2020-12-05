package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixtures struct {
	Input    []int
	Expected int
}

func TestPart1(t *testing.T) {
	for _, fixture := range []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 514579,
		},
	} {
		assert.Equal(t, Part1(fixture.Input), fixture.Expected)
	}
}

func TestPart2(t *testing.T) {
	for _, fixture := range []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 241861950,
		},
	} {
		assert.Equal(t, Part2(fixture.Input), fixture.Expected)
	}
}
