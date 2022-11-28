package binarydiagnostic

import (
	"strings"
)

func Part1() int {
	diagnostics := strings.Split(input, "\n")

	// Renamed the variables which maintain the "shape" of the binary input.
	n, m := len(diagnostics), len(diagnostics[0])

	// Prefer using math to maintain the binary values we build up
	// instead of leveraging strings and built-ins which parse them
	// to decimal.
	gamma := make([]byte, m)
	epsilon := make([]byte, m)

	for col := 0; col < m; col++ {

		// Whenever a map can be indexed by integers in consecutive
		// order, you can replace it with a slice (or array).
		//  - onOff[0] maintains the count of zero
		//  - onOff[1] maintains the count of one
		var onOff [2]int

		for row := 0; row < n; row++ {
			// We know that the characters are only '0' and '1' so
			// every character is a single byte and we can use 
			// naive indexing.
			bits := []byte(diagnostics[row])

			// Subtracting the lowest character from the range of character
			// values is a trick to get the integer value. This only works
			// for ASCII characters since they're always sorted ascending.
			// https://stackoverflow.com/a/3195042
			bit := bits[col] - '0'
			onOff[bit]++
		}
		if onOff[1] > onOff[0] {
			gamma[col]++
		} else {
			epsilon[col]++
		}
	}

	// Returning the answer value to allow us to test this
	// on other inputs, or do benchmarking.
	return bitsToInt(gamma) * bitsToInt(epsilon)
}

func bitsToInt(bits []byte) int {
	v := 0
	for i, b := range bits {
		// https://stackoverflow.com/a/23189744
		v += int(b) << (len(bits) - (i + 1))
	}
	return v
}
