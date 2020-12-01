package main

import (
	"testing"
)

type Fixtures struct {
	Input    []int
	Expected int
}

func TestCalculateTwoEntry2020Sum(t *testing.T) {
	fixtures := []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 514579,
		},
	}
	for _, fixture := range fixtures {
		got := CalculateTwoEntry2020Sum(fixture.Input)
		if got != fixture.Expected {
			t.Errorf("CalculateTwoEntry2020Sum(%d) = %d; want %d", fixture.Input, got, fixture.Expected)
		}
	}
}

func TestCalculateThreeEntry2020Sum(t *testing.T) {
	fixtures := []Fixtures{
		{
			[]int{1721, 979, 366, 299, 675, 1456}, 241861950,
		},
	}
	for _, fixture := range fixtures {
		got := CalculateThreeEntry2020Sum(fixture.Input)
		if got != fixture.Expected {
			t.Errorf("CalculateThreeEntry2020Sum(%d) = %d; want %d", fixture.Input, got, fixture.Expected)
		}
	}
}
