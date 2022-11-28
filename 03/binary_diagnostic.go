package binarydiagnostic

import (
	"aoc/inputs"
	"aoc/utils"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	diagnosticsSlice := strings.Split(inputs.Day3, "\n")
	diagnostic := strings.Split(diagnosticsSlice[0], "")
	rowLen := len(diagnostic)
	var gammaRate bytes.Buffer
	var epsilonRate bytes.Buffer
	for i := 0; i < rowLen; i++ {
		m := make(map[string]int)
		m["on"] = 0
		m["off"] = 0
		for j := 0; j < len(diagnosticsSlice); j++ {
			diagnosticBits := strings.Split(diagnosticsSlice[j], "")
			if diagnosticBits[i] == "0" {
				m["off"]++
			}
			if diagnosticBits[i] == "1" {
				m["on"]++
			}
		}
		if m["on"] > m["off"] {
			gammaRate.WriteString("1")
			epsilonRate.WriteString("0")
		} else {
			gammaRate.WriteString("0")
			epsilonRate.WriteString("1")
		}
	}

	gammaDecimal, err := strconv.ParseInt(gammaRate.String(), 2, 64)
	utils.ErrorCheck(err)
	epsilonDecimal, err := strconv.ParseInt(epsilonRate.String(), 2, 64)
	utils.ErrorCheck(err)

	subPower := gammaDecimal * epsilonDecimal

	fmt.Printf("Day 3 - part 1 answer %d\n", subPower)

}
