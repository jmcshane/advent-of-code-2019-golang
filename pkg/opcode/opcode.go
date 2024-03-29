package opcode

import (
	"fmt"

	"github.com/jmcshane/advent-of-code/pkg/util"
)

var input = []int{0}
var inputIx = 0

func ResetInput(newInput []int) {
	input = newInput
	inputIx = 0
}

func ProcessOpcodes(ix int, output *int, inputOpcodes []int) int {
	if inputOpcodes[ix] == 99 {
		if output != nil {
			return *output
		}
		panic("no current output value")
	}
	return ProcessOpcodes(ProcessOpcode(ix, output, inputOpcodes))
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

func ProcessOpcode(ix int, output *int, inputOpcodes []int) (int, *int, []int) {
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
		inputOpcodes[position] = input[inputIx]
		inputIx++
		return ix + 2, nil, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 4 {
		if util.Digit(inputOpcodes[ix], 3) == 1 {
			return ix + 2, &inputOpcodes[ix+1], inputOpcodes
		}
		return ix + 2, &inputOpcodes[inputOpcodes[ix+1]], inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 5 {
		_, a, b := getThreeArgVal(ix, inputOpcodes)
		if a != 0 {
			return b, output, inputOpcodes
		}
		return ix + 3, output, inputOpcodes
	}
	if util.Digit(inputOpcodes[ix], 1) == 6 {
		_, a, b := getThreeArgVal(ix, inputOpcodes)
		if a == 0 {
			return b, output, inputOpcodes
		}
		return ix + 3, output, inputOpcodes
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
