package main

import (
	"testing"
)

type Fixtures struct {
	Mass     int
	Expected int
}

func TestCalculateFuel(t *testing.T) {
	fixtures := []Fixtures{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, fixture := range fixtures {
		got := CalculateFuel(fixture.Mass)
		if got != fixture.Expected {
			t.Errorf("CalculateFuel(%d) = %d; want %d", fixture.Mass, got, fixture.Expected)
		}
	}
}

func TestCalculateFuel2(t *testing.T) {
	fixtures := []Fixtures{
		{12, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, fixture := range fixtures {
		got := CalculateFuel2(fixture.Mass)
		if got != fixture.Expected {
			t.Errorf("CalculateFuel2(%d) = %d; want %d", fixture.Mass, got, fixture.Expected)
		}
	}
}
