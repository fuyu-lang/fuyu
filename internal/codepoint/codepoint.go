package codepoint

import (
	"fmt"
	"unicode"
)

// u converts a rune to string representing it in one of the forms: U+XXXX,
// U+XXXXXX, U+XXXXXXXX, where X is an uppercase hexadecimal digit. The
// representation with the fewest leading zeroes is selected.
func u(r rune) string {
	c := uint32(r)
	switch {
	case c <= 0xFFFF:
		return fmt.Sprintf("U+%04X", c)
	case c <= 0xFFFFFF:
		return fmt.Sprintf("U+%06X", c)
	default:
		return fmt.Sprintf("U+%08X", c)
	}
}

// Humanize converts a rune to a human readable string representation.
func Humanize(r rune) string {
	if unicode.IsSpace(r) {
		switch r {
		case '\t':
			return fmt.Sprintf("tab (%s)", u(r))
		case '\n':
			return fmt.Sprintf("line feed (%s)", u(r))
		case '\r':
			return fmt.Sprintf("carriage return (%s)", u(r))
		case ' ':
			return fmt.Sprintf("space (%s)", u(r))
		default:
			return u(r)
		}
	}
	if unicode.IsPrint(r) {
		return fmt.Sprintf("%c", r)
	}
	return u(r)
}
