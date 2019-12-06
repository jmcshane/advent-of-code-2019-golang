package util

import (
	"math"
	"strconv"
	"strings"
)

func Digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func ConvertInputArray(input []string) []int {
	var inputProg = []int{}

	for _, i := range input {
		j, err := strconv.Atoi(strings.Trim(i, "\n"))
		if err != nil {
			panic(err)
		}
		inputProg = append(inputProg, j)
	}
	return inputProg
}
