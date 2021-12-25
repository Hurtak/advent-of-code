package main

func CalculateFuel(mass int) int {
	return mass/3 - 2
}

func Part1(mass int) int {
	return CalculateFuel(mass)
}

func Part2(mass int) int {
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
