package utils

import (
	"fmt"
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

func PrintResult(result1 interface{}, result2 interface{}) {
	fmt.Println("Result1:", result1)
	fmt.Println("Result2:", result2)
}
