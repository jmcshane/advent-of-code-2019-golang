package main

import (
	"fmt"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/input"
	"github.com/jmcshane/advent-of-code/pkg/opcode"
	"github.com/jmcshane/advent-of-code/pkg/util"
)

var (
	progInput = 5
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
	output := Run(inputOpcodes)
	fmt.Printf("output is %d", output)
}

func Run(inputOpcodes []int) int {
	opcode.ResetInput([]int{progInput})
	output := opcode.ProcessOpcodes(0, nil, inputOpcodes)
	return output
}
