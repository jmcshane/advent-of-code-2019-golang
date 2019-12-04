package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/jmcshane/advent-of-code/pkg/input"
)

const (
	divisiorFactor = 3
	subtractFactor = 2
)

var (
	home          string = os.Getenv("AOC_HOME")
	recursiveFuel string = os.Getenv("FUEL_RECURSE")
)

func main() {
	content, err := input.ReadInputFile(1)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for _, line := range lines {
		if strings.Trim(line, " ") != "" {
			sum += calculateFuel(line)
		}
	}
	fmt.Printf("The sum is %d\n", sum)
}

func calculateFuel(line string) int {
	inputVal, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("Error converting line %s: %v", line, err)
		panic(err)
	}
	fuelValue := (inputVal / divisiorFactor) - subtractFactor
	if strings.ToLower(recursiveFuel) == "true" {
		fuelValue += recurseCalculate(fuelValue)
	}
	return fuelValue
}

func recurseCalculate(input int) int {
	if input < 9 {
		return 0
	}
	value := (input / divisiorFactor) - subtractFactor
	return value + recurseCalculate(value)
}
