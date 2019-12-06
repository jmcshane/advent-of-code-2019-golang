package main

import (
	"fmt"
	"testing"
)

func TestBasicInput(t *testing.T) {
	var inputArray = []int{1, 1, 0, 3, 4, 0, 99}
	fmt.Printf("Output is %d", ProcessOpcodes(0, nil, inputArray))
}

func TestMultipleOutput(t *testing.T) {
	var inputArray = []int{1101, 1, 1, 1, 2, 1, 4, 0, 104, 0, 4, 0, 99}
	fmt.Printf("Output is %d", ProcessOpcodes(0, nil, inputArray))
}

func TestBasicNewOpcodes(t *testing.T) {
	var inputArray = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	SetProgInput(8)
	fmt.Printf("Output should be 1 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	SetProgInput(7)
	fmt.Printf("Output should be 0 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	fmt.Printf("Output should be 1 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	SetProgInput(9)
	fmt.Printf("Output should be 0 -> %d\n", ProcessOpcodes(0, nil, inputArray))
}

func TestBasicJumps(t *testing.T) {
	var inputArray = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	SetProgInput(0)
	fmt.Printf("Output should be 0 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Printf("Output should be 0 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	SetProgInput(98)
	fmt.Printf("Output should be 1 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	fmt.Printf("Output should be 1 -> %d\n", ProcessOpcodes(0, nil, inputArray))
}

func TestMoreComplexJump(t *testing.T) {
	inputArray := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	SetProgInput(7)
	fmt.Printf("Output should be 999 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	SetProgInput(8)
	fmt.Printf("Output should be 1000 -> %d\n", ProcessOpcodes(0, nil, inputArray))
	inputArray = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	SetProgInput(9)
	fmt.Printf("Output should be 1001 -> %d\n", ProcessOpcodes(0, nil, inputArray))
}
