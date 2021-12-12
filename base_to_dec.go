package algorithms

import (
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {

	var cv int
	// base^bit place in value
	multiplier := 1

	for i := len(value) - 1; i >= 0; i-- {
		s := hexToDec(value[i])
		cv += s * multiplier
		multiplier *= base
	}

	return cv
}

func hexToDec(b byte) int {
	s := string(b)

	switch s {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		n, _ := strconv.Atoi(s)
		return n
	}
}
