package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileStringWhole(path string) string {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(input))

	return data
}

func ReadFileString(path string) []string {
	data := ReadFileStringWhole(path)

	return strings.Split(data, "\n")
}

func ReadFileInt(path string) []int {
	list := ReadFileString(path)

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

func PrintResult(result1 interface{}, result2 interface{}) {
	fmt.Println("Result1:", result1)
	fmt.Println("Result2:", result2)
}
