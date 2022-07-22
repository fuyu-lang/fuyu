package lexer

import (
	"errors"
	"fmt"
	"github.com/fuyu-lang/fuyu/internal/codepoint"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"unicode/utf8"
)

// checkRune checks that a rune is one of the valid runes that is allowed to
// appear in a source text. The following are not allowed to appear anywhere:
//
// - U+0000 (NUL) to U+0009 (TAB)
// - U+000B (VT) to U+000C (FF)
// - U+000E (SO) to U+001F (US)
// - U+007F (DEL) to U+009F (APC)
func checkRune(r rune) error {
	if (0x000 <= r && r <= 0x0009) ||
		(0x000B <= r && r <= 0x000C) ||
		(0x000E <= r && r <= 0x001F) ||
		(0x007F <= r && r <= 0x009F) {
		return errors.New(codepoint.Humanize(r))
	}
	return nil
}

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
					Found: "invalid bit sequence in UTF-8 string",
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

// checkRune checks that a rune is one of the valid runes that is allowed to
// appear in a source text. The following are not allowed to appear anywhere:
// - the input is a valid UTF-8 encoded string
// - The following codepoints ranges do not appear anywhere in the string:
//  - U+0000 (NUL) to U+0009 (TAB)
//  - U+000B (VT) to U+000C (FF)
//  - U+000E (SO) to U+001F (US)
//  - U+007F (DEL) to U+009F (APC)
// - Every carriage return (U+000D) is immediately followed by a linefeed
//   (U+000A).
func validSourceEncoding(maybeUtf8 string) error {
	if err := checkUtf8(maybeUtf8); err != nil {
		return err
	}
	index, line, col := 0, 1, 1
	prev := '\u0000'
	for index < len(maybeUtf8) {
		r, size := utf8.DecodeRuneInString(maybeUtf8[index:])
		if err := checkRune(r); err != nil {
			return &compiler.Error{
				Index: index, Line: line, Col: col,
				Found: fmt.Sprintf("invalid UTF-8 codepoint %s", err.Error()),
			}
		}
		if r == '\n' {
			line += 1
			col = 1
		} else {
			// an \r was not followed by an \n
			if prev == '\r' {
				return &compiler.Error{
					Index: index, Line: line, Col: col,
					Expect: compiler.ExpectLFAfterCR,
					Found:  codepoint.Humanize(r),
				}
			}
			col += 1
		}
		index += size
		prev = r
	}
	// The last character was an \r and therefore not followed by an \n
	if prev == '\r' {
		return &compiler.Error{
			Index: index, Line: line, Col: col,
			Expect: compiler.ExpectLFAfterCR,
			Found:  compiler.FoundEOF,
		}
	}
	return nil
}
