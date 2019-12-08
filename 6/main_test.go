package main

import "testing"

func TestBasic(t *testing.T) {
	input := []string{"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	CountNodes(input)
}
