package main

import (
	"fmt"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/input"
	"github.com/jmcshane/advent-of-code/pkg/util"
)

var (
	progInput int = 5
)

func SetProgInput(input int) {
	progInput = input
}

func main() {
	content, err := input.ReadInputFile(5)
	if err != nil {
		panic(err)
	}
	opcodes := strings.Split(content, ",")
	inputOpcodes := util.ConvertInputArray(opcodes)
	output := ProcessOpcodes(0, nil, inputOpcodes)
	fmt.Printf("output is %d", output)
}

func ProcessOpcodes(ix int, output *int, inputOpcodes []int) int {
	if inputOpcodes[ix] == 99 {
		if output != nil {
			return *output
		}
		panic("no current output value")
	}
	if output != nil && *output != 0 {
		fmt.Printf("Invalid output at operation before index %d", ix)
		panic("Failure of output")
	}
	return ProcessOpcodes(ProcessOpcode(ix, inputOpcodes))
}

func getThreeArgVal(ix int, inputOpcodes []int) (int, int, int) {
	first := inputOpcodes[ix+1]
	second := inputOpcodes[ix+2]
	position := inputOpcodes[ix+3]
	if util.Digit(inputOpcodes[ix], 3) == 0 {
		first = inputOpcodes[first]
	}
	if util.Digit(inputOpcodes[ix], 4) == 0 {
		second = inputOpcodes[second]
	}
	return position, first, second
}

func ProcessOpcode(ix int, inputOpcodes []int) (int, *int, []int) {
	fmt.Printf("Process codes %d at index %d\n", inputOpcodes, ix)
	if util.Digit(inputOpcodes[ix], 1) == 1 {
		modifyIX, a, b := getThreeArgVal(ix, inputOpcodes)
		inputOpcodes[modifyIX] = a + b
		return ix + 4, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 2 {
		modifyIX, a, b := getThreeArgVal(ix, inputOpcodes)
		inputOpcodes[modifyIX] = a * b
		return ix + 4, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 3 {
		if inputOpcodes[ix] != 3 {
			fmt.Printf("Not sure what to do with input %d. Next input is %d", inputOpcodes[ix], inputOpcodes[ix+1])
			panic("Invalid opcode")
		}
		position := inputOpcodes[ix+1]
		inputOpcodes[position] = progInput
		return ix + 2, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 4 {
		output := &inputOpcodes[inputOpcodes[ix+1]]
		if util.Digit(inputOpcodes[ix], 3) == 1 {
			output = &inputOpcodes[ix+1]
		}
		return ix + 2, output, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 5 {
		_, a, b := getThreeArgVal(ix, inputOpcodes)
		if a != 0 {
			return b, nil, inputOpcodes
		}
		return ix + 3, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 6 {
		_, a, b := getThreeArgVal(ix, inputOpcodes)
		if a == 0 {
			return b, nil, inputOpcodes
		}
		return ix + 3, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 7 {
		modifyIX, a, b := getThreeArgVal(ix, inputOpcodes)
		if a < b {
			inputOpcodes[modifyIX] = 1
		} else {
			inputOpcodes[modifyIX] = 0
		}
		return ix + 4, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 8 {
		modifyIX, a, b := getThreeArgVal(ix, inputOpcodes)
		if a == b {
			inputOpcodes[modifyIX] = 1
		} else {
			inputOpcodes[modifyIX] = 0
		}
		return ix + 4, nil, inputOpcodes
	}
	return 0, nil, []int{}
}
