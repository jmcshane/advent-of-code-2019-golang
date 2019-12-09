package main

import (
	"fmt"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/opcode"

	"github.com/jmcshane/advent-of-code/pkg/input"
	"github.com/jmcshane/advent-of-code/pkg/util"
)

const initialInput = 0

var (
	InputMin = 5
	InputMax = 9
)

func SetMinMax(min, max int) {
	InputMin = min
	InputMax = max
}

func main() {
	content, err := input.ReadInputFile(7)
	if err != nil {
		panic(err)
	}
	opcodes := strings.Split(content, ",")
	inputOpcodes := util.ConvertInputArray(opcodes)
	output := Run(inputOpcodes)
	fmt.Printf("output is %d", output)
}

func Run(inputOpcodes []int) int {
	permutations := Permutations()
	outputValues := []int{}
	output := initialInput
	for _, permutation := range permutations {
		for _, val := range permutation {
			inputCopy := make([]int, len(inputOpcodes))
			copy(inputCopy, inputOpcodes)
			input := []int{val, output}
			opcode.ResetInput(input)
			output = opcode.ProcessOpcodes(0, nil, inputCopy)
		}
		outputValues = append(outputValues, output)
		output = 0
	}
	return Max(outputValues)
}

func Max(array []int) int {
	if len(array) == 0 {
		return 0
	}
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func Permutations() [][]int {
	output := [][]int{}
	for a := InputMin; a < InputMax; a++ {
		for b := InputMin; b < InputMax; b++ {
			for c := InputMin; c < InputMax; c++ {
				for d := InputMin; d < InputMax; d++ {
					for e := InputMin; e < InputMax; e++ {
						if a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e {
							output = append(output, []int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}
	return output
}
