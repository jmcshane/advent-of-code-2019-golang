package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/input"
	"github.com/jmcshane/advent-of-code/pkg/util"
)

func main() {
	content, err := input.ReadInputFile(2)
	if err != nil {
		panic(err)
	}
	prog := strings.Split(content, ",")
	inputProg := util.ConvertInputArray(prog)
	//Part 1
	// outputProg := ProcessProg(ApplyInputs(12, 2, inputProg))
	// fmt.Printf("The first arg is %d", outputProg[0])

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			testSlice := make([]int, len(inputProg))
			copy(testSlice, inputProg)
			valuesSlice := ApplyInputs(i, j, testSlice)
			valuesSlice, err := ProcessProg(valuesSlice)
			if err == nil && len(valuesSlice) > 0 {
				if valuesSlice[0] == 19690720 {
					fmt.Println("FOUND VALUE. Noun: %d, Verb: %d", i, j)
				}
			} else {
				fmt.Printf("For noun %d and verb %d, invalid output\n", i, j)
			}
		}
	}
}
func ProcessProg(inputProg []int) ([]int, error) {
	pos := 0
	for {
		if inputProg[pos] == 99 {
			break
		}
		posFirst := inputProg[pos+1]
		posSecond := inputProg[pos+2]
		positionVal := inputProg[pos+3]
		if inputProg[pos] == 1 {
			inputProg[positionVal] = inputProg[posFirst] + inputProg[posSecond]
		} else if inputProg[pos] == 2 {
			inputProg[positionVal] = inputProg[posFirst] * inputProg[posSecond]
		} else {
			return []int{}, errors.New("Invalid operation")
		}
		pos += 4
	}
	return inputProg, nil
}

func ApplyInputs(noun, verb int, input []int) []int {
	input[1] = noun
	input[2] = verb
	return input
}
