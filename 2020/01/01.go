package main

import (
	"fmt"

	"github.com/hurtak/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadFileInt("01.txt")

	fmt.Println(CalculateTwoEntry2020Sum(input))
	fmt.Println(CalculateThreeEntry2020Sum(input))
}

const target = 2020

func CalculateTwoEntry2020Sum(numbers []int) int {
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			if number1+number2 == target {
				return number1 * number2
			}
		}
	}

	return -1
}

func CalculateThreeEntry2020Sum(numbers []int) int {
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			for _, number3 := range numbers {
				if number1+number2+number3 == target {
					return number1 * number2 * number3
				}
			}
		}
	}

	return -1
}
