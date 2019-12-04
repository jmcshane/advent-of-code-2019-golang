package main

import (
	"fmt"
	"math"
)

const (
	start int = 134792
	end   int = 675810
	// start int = 134772
	// end   int = 134778
	// start int = 111121
	// end   int = 111123
)

func main() {
	count := 0
	for val := start; val <= end; val++ {
		if checkValue(val) {
			count++
		}
	}
	fmt.Printf("The count is %d\n", count)
}

func checkValue(value int) bool {
	digits := []int{}
	for i := 1; math.Pow(10, float64(i-1)) < float64(value); i++ {
		digits = append(digits, digit(value, i))
	}
	foundDup := false
	for ix := range digits {
		if ix == 0 {
			continue
		}
		if digits[ix]-digits[ix-1] > 0 {
			return false
		}
		if digits[ix]-digits[ix-1] == 0 {
			if ix < 2 || digits[ix-1] != digits[ix-2] {
				if ix > len(digits)-2 || digits[ix] != digits[ix+1] {
					foundDup = true
				}
			}
		}
	}
	return foundDup
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
