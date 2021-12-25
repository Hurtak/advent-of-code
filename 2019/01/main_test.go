package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixtures struct {
	Input    int
	Expected int
}

func TestPart1(t *testing.T) {
	for _, fixture := range []Fixtures{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	} {
		res := Part1(fixture.Input)
		assert.Equal(t, fixture.Expected, res, fixture.Input)
	}
}

func TestPart2(t *testing.T) {
	for _, fixture := range []Fixtures{
		{12, 2},
		{1969, 966},
		{100756, 50346},
	} {
		res := Part2(fixture.Input)
		assert.Equal(t, fixture.Expected, res, fixture.Input)
	}
}
