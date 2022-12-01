package main

import (
	depthfinder "aoc/01"
	dive "aoc/02"
	binarydiagnostic "aoc/03"
	bingo "aoc/04"
	smokebasin "aoc/09"
	"aoc/inputs"
	"fmt"
)

func main() {
	// Day 1
	depthfinder.GetdepthIncreaseCount(inputs.DepthFinderInputs)
	depthfinder.GetDepthSumIncreaseCount(inputs.DepthFinderInputs)
	fmt.Println("-----------------")
	// // Day 2
	dive.Start()
	fmt.Println("-----------------")
	// // Day 3
	fmt.Printf("Day 3 - part 1 answer %d\n", binarydiagnostic.Part1())
	fmt.Printf("Day 3 - part 2 answer %d\n", binarydiagnostic.Part2())
	fmt.Println("-----------------")
	// Day 4
	winningBoard := bingo.Part1()
	fmt.Printf("Day 3 - part 1 answer %d\n", winningBoard.Score)

	fmt.Println("-----------------")
	// Day 9

	smokebasin.Solution1()

}
