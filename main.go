package main

import (
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
	// Day 2
	dive.Start()
	// Day 3
	binarydiagnostic.Part1()
	// Day 9
	smokebasin.Solution1()
}
