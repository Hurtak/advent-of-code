package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func get2020Result(numbers []int) int {
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

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile(path.Join(cwd, "./src/01/input.txt"))

	if err != nil {
		panic(err)
	}

	list := strings.Fields(string(input))

	var listNumbers []int
	for _, element := range list {
		number, err := strconv.Atoi(element)

		if err != nil {
			continue
		}
		listNumbers = append(listNumbers, number)
	}

	res := get2020Result(listNumbers)

	fmt.Println(res)
}
