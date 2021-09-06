package stringutil

import (
	"strings"
)

// ReverseString reverses a given string.
func ReverseString(s string) string {
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}

	return b.String()
}
