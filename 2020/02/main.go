package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hurtak/advent-of-code/utils"
)

func main() {
	input := utils.ReadFileString("data.txt")

	utils.PrintResult(getValidPasswordsCount(Part1, input), getValidPasswordsCount(Part2, input))
}

func getValidPasswordsCount(funcion func(string) (bool, error), input []string) int {
	validPasswords := 0
	for _, rule := range input {
		valid, err := funcion(rule)
		if err != nil {
			panic(err)
		}
		if valid {
			validPasswords++
		}
	}

	return validPasswords
}

func getInputData(input string) (int, int, string, string, error) {
	regexString := `^(\d+)-(\d+) (\w): (\w+)$`
	r := regexp.MustCompile(regexString)

	match := r.FindStringSubmatch(input)
	if len(match) == 0 {
		return 0, 0, "", "", fmt.Errorf("input %q did not match the format %q", input, regexString)
	}

	first, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, "", "", err
	}

	second, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, "", "", err
	}

	char := match[3]
	password := match[4]

	return first, second, char, password, nil
}

func Part1(input string) (bool, error) {
	min, max, char, password, err := getInputData(input)
	if err != nil {
		return false, err
	}

	charsInPass := 0
	for _, passwordChar := range password {
		if string(passwordChar) == char {
			charsInPass++
		}
	}

	passwordValid := charsInPass >= min && charsInPass <= max

	return passwordValid, nil
}

func Part2(input string) (bool, error) {
	first, second, char, password, err := getInputData(input)
	if err != nil {
		return false, err
	}

	positionsMatched := 0
	for i, passwordChar := range password {
		if string(passwordChar) == char {
			position := i + 1
			if position == first || position == second {
				positionsMatched++
			}
		}
	}

	passwordValid := positionsMatched == 1

	return passwordValid, nil
}
