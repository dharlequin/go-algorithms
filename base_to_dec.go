package algorithms

import (
	"math"
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
	cs := Reverse(value)

	for i := len(value) - 1; i >= 0; i-- {
		s, err := strconv.Atoi(string(cs[i]))
		if err != nil {
			s = hexToDec(string(cs[i]))
		}
		cv += int(s) * int(math.Pow(float64(base), float64(i)))
	}

	return cv
}

func hexToDec(s string) int {
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
		return 0
	}
}
