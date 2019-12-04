package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/input"
)

func main() {
	content, err := input.ReadInputFile(3)
	if err != nil {
		panic(err)
	}
	paths := strings.Split(content, "\n")
	inputPaths := []Path{}
	for _, path := range paths {
		if path == "" {
			continue
		}
		instrStrings := strings.Split(path, ",")
		instructions := []Instruction{}
		for _, instr := range instrStrings {
			if instr == "" {
				continue
			}
			distance, err := strconv.Atoi(instr[1:])
			if err != nil {
				panic(err)
			}
			instruction := Instruction{
				Direction: instr[0:1],
				Distance:  distance,
			}
			instructions = append(instructions, instruction)
		}
		inputPath := Path{Instructions: instructions}
		inputPaths = append(inputPaths, inputPath)
	}
	for ix, path := range inputPaths {
		currentX := 0
		currentY := 0
		pathSegments := []LineSegment{}
		for i := 0; i < len(path.Instructions); i++ {
			instr := path.Instructions[i]
			moveFactor := 1
			if instr.Direction == "R" || instr.Direction == "L" {
				if instr.Direction == "L" {
					moveFactor = -1
				}
				segment := LineSegment{
					Start:     currentX,
					End:       currentX + moveFactor*instr.Distance,
					Stable:    currentY,
					Direction: "horizontal",
				}
				currentX = segment.End
				pathSegments = append(pathSegments, segment)
			} else {
				if instr.Direction == "D" {
					moveFactor = -1
				}
				segment := LineSegment{
					Start:     currentY,
					End:       currentY + moveFactor*instr.Distance,
					Stable:    currentX,
					Direction: "vertical",
				}
				currentY = segment.End
				pathSegments = append(pathSegments, segment)
			}
		}
		path.LineSegments = pathSegments
		inputPaths[ix] = path
	}
	// fmt.Printf("First Path.Linesegment: %d\n\n", inputPaths[0].LineSegments)
	// fmt.Printf("SEcond Path.Linesegment: %d\n\n", inputPaths[1].LineSegments)

	// fmt.Printf("Paths: %d", inputPaths)
	firstPath := inputPaths[0]
	secondPath := inputPaths[1]
	intersections := []Intersection{}
	// fmt.Printf("Segments for first: %d\n\n", firstPath.LineSegments)
	// fmt.Printf("Segments for second: %d\n\n", secondPath.LineSegments)
	distanceFirst := 0
	for _, seg := range firstPath.LineSegments {
		distanceSecond := 0
		for ix, compareSeg := range secondPath.LineSegments {
			if intersect := cross(seg, compareSeg); intersect != nil && ix != 0 {
				intersect.TotalDistance = distanceFirst + distanceSecond + intersect.IntersectionDistance
				intersections = append(intersections, *intersect)
			}
			distAdd := Abs(compareSeg.End - compareSeg.Start)
			distanceSecond += distAdd
		}
		distFirstAdd := Abs(seg.End - seg.Start)
		distanceFirst += distFirstAdd
	}
	fmt.Printf("Intersections %d", intersections)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cross(first, second LineSegment) *Intersection {
	if first.Direction == second.Direction {
		return nil
	}
	if (first.Start-second.Stable)*(first.End-second.Stable) < 0 {
		if (second.Start-first.Stable)*(second.End-first.Stable) < 0 {
			return &Intersection{
				X:                    first.Stable,
				Y:                    second.Stable,
				IntersectionDistance: Abs(first.Start-second.Stable) + Abs(second.Start-first.Stable),
			}
		}
	}
	return nil
}

type Intersection struct {
	X                    int
	Y                    int
	IntersectionDistance int
	TotalDistance        int
}

type Instruction struct {
	Direction string
	Distance  int
}

type Path struct {
	Instructions []Instruction
	LineSegments []LineSegment
}

type LineSegment struct {
	Start     int
	End       int
	Stable    int
	Direction string
}
