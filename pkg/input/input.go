package input

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	home string = os.Getenv("AOC_HOME")
)

func ReadInputFile(puzzleNum int) (string, error) {
	if home == "" {
		home = "/Users/jmcshan2/Projects/advent-of-code"
	}
	filename := fmt.Sprintf("%s/input/%d", home, puzzleNum)
	fmt.Printf("Reading file %s\n", filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
