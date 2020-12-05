package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileInt(path string) []int {
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

	return numbers

}
