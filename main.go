package main

import (
	"fmt"

	depthfinder "aoc/01"
	dive "aoc/02"
	binarydiagnostic "aoc/03"
	smokebasin "aoc/09"
	"aoc/inputs"
)

func main() {
	// Day 1
	depthfinder.GetdepthIncreaseCount(inputs.DepthFinderInputs)
	depthfinder.GetDepthSumIncreaseCount(inputs.DepthFinderInputs)
	fmt.Println("-----------------")
	// Day 2
	dive.Start()
	fmt.Println("-----------------")
	// Day 3
	fmt.Printf("Day 3 - part 1 answer %d\n", binarydiagnostic.Part1())
	fmt.Printf("Day 3 - part 2 answer %d\n", binarydiagnostic.Part2())
	fmt.Println("-----------------")
	// Day 9
	smokebasin.Solution1()
}
