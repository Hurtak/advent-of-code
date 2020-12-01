package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("01.txt")
	if err != nil {
		panic(err)
	}

	var modules []int
	for _, element := range strings.Fields(string(input)) {
		number, err := strconv.Atoi(element)

		if err != nil {
			continue
		}
		modules = append(modules, number)
	}

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
