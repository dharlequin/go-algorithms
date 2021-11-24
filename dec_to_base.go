package algorithms

import (
	"fmt"
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	var sb strings.Builder

	for dec > 0 {
		r := dec % base

		sb.WriteString(fmt.Sprintf("%X", r))

		dec = dec / base
	}

	return Reverse(sb.String())
}