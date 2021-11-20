package algorithms

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var sb strings.Builder

	if word == "" {
		return word
	}

	for i := len([]rune(word)) - 1; i >= 0; i-- {
		sb.WriteRune([]rune(word)[i])
	}

	return sb.String()
}
