package main

import (
	"fmt"

	"github.com/hurtak/advent-of-code-2020/utils"
)

func main() {
	modules := utils.ReadFileInt("data.txt")

	fuel := 0
	for _, module := range modules {
		fuel += CalculateFuel(module)
	}

	fuelAcc := 0
	for _, module := range modules {
		fuelAcc += CalculateFuel2(module)
	}

	fmt.Println(fuel, fuelAcc)
}

func CalculateFuel(mass int) int {
	return mass/3 - 2
}

func CalculateFuel2(mass int) int {
	sum := 0
	for {
		fuel := CalculateFuel(mass)
		if fuel <= 0 {
			return sum
		}
		sum += fuel

		mass = CalculateFuel(fuel)
		if mass <= 0 {
			return sum
		}
		sum += mass
	}
}
