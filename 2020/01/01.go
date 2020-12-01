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

	list := strings.Fields(string(input))
	var numbers []int
	for _, element := range list {
		number, err := strconv.Atoi(element)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}

	res1 := CalculateTwoEntry2020Sum(numbers)
	fmt.Println(res1)
	res2 := CalculateThreeEntry2020Sum(numbers)
	fmt.Println(res2)
}

func CalculateTwoEntry2020Sum(numbers []int) int {
	for i1, number1 := range numbers {
		for i2, number2 := range numbers {
			if i1 == i2 {
				continue
			}
			if number1+number2 == 2020 {
				return number1 * number2
			}
		}
	}

	return -1
}

func CalculateThreeEntry2020Sum(numbers []int) int {
	for i1, number1 := range numbers {
		for i2, number2 := range numbers {
			for i3, number3 := range numbers {
				if i1 == i2 || i1 == i3 || i2 == i3 {
					continue
				}
				if number1+number2+number3 == 2020 {
					return number1 * number2 * number3
				}
			}
		}
	}

	return -1
}
