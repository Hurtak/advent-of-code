package main

import (
	"github.com/hurtak/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadFileInt("data.txt")
	utils.PrintResult(Part1(input), Part2(input))
}

const target = 2020

func Part1(numbers []int) int {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1+n2 == target {
				return n1 * n2
			}
		}
	}

	return -1
}

func Part2(numbers []int) int {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			for _, n3 := range numbers {
				if n1+n2+n3 == target {
					return n1 * n2 * n3
				}
			}
		}
	}

	return -1
}
