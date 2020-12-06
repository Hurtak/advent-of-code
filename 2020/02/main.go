package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hurtak/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadFileString("data.txt")

	validPasswords := 0
	for _, rule := range input {
		valid, err := Part1(rule)
		if err != nil {
			panic(err)
		}
		if valid {
			validPasswords++
		}
	}

	utils.PrintResult(validPasswords, 0)
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
