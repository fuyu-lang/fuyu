package lexer

import (
	"fmt"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"unicode/utf8"
)

// checkUtf8 checks that a string contains only valid UTF-8.
func checkUtf8(maybeUtf8 string) error {
	index, line, col := 0, 1, 1
	for index < len(maybeUtf8) {
		r, size := utf8.DecodeRuneInString(maybeUtf8[index:])
		if r == utf8.RuneError {
			if size == 0 {
				// This is fine because it is an empty string. This will
				// also never get hit because the loop will stop before
				// this happens.
			} else if size == 1 {
				return &compiler.Error{
					Index: index, Line: line, Col: col,
					Found: fmt.Sprintf("invalid UTF-8 at byte %v", index),
				}
			} else {
				// This is a valid rune with a value equal to
				// utf8.RuneError
			}
		}
		if r == '\n' {
			line += 1
			col = 1
		} else {
			col += 1
		}
		index += size
	}
	return nil
}
