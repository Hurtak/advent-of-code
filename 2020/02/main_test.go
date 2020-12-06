package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixtures struct {
	Input    string
	Expected bool
}

func TestPart1(t *testing.T) {
	for _, fixture := range []Fixtures{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
		{"10-20 c: cccccccccc", true},
	} {
		res, err := Part1(fixture.Input)
		assert.Equal(t, fixture.Expected, res, fixture.Input, err)
	}
}

func TestPart2(t *testing.T) {
	for _, fixture := range []Fixtures{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
		{"1-2 a: axaaaa", true},
	} {
		res, err := Part2(fixture.Input)
		assert.Equal(t, fixture.Expected, res, fixture.Input, err)
	}
}
