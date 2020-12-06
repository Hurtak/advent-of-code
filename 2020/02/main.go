package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hurtak/advent-of-code-2020/utils"
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

func Part1(input string) (bool, error) {
	regexString := `^(\d+)-(\d+) (\w): (\w+)$`
	r := regexp.MustCompile(regexString)

	match := r.FindStringSubmatch(input)
	if len(match) == 0 {
		return false, fmt.Errorf("input %q did not match the format %q", input, regexString)
	}

	min, err := strconv.Atoi(match[1])
	if err != nil {
		return false, err
	}

	max, err := strconv.Atoi(match[2])
	if err != nil {
		return false, err
	}

	char := match[3]
	password := match[4]

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
	regexString := `^(\d+)-(\d+) (\w): (\w+)$`
	r := regexp.MustCompile(regexString)

	match := r.FindStringSubmatch(input)
	if len(match) == 0 {
		return false, fmt.Errorf("input %q did not match the format %q", input, regexString)
	}

	min, err := strconv.Atoi(match[1])
	if err != nil {
		return false, err
	}

	max, err := strconv.Atoi(match[2])
	if err != nil {
		return false, err
	}

	char := match[3]
	password := match[4]

	positionsMatched := 0
	for i, passwordChar := range password {
		if string(passwordChar) == char {
			pos := i + 1
			if pos == min || pos == max {
				positionsMatched++
			}
		}
	}

	passwordValid := positionsMatched == 1

	return passwordValid, nil
}
