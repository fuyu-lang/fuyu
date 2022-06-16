package compiler

import (
	"fmt"
)

// Error is an error that can occur during compilation.
type Error struct {
	Index  int    // The index of the first invalid byte in the source text.
	Line   int    // The line of the first invalid code point in the souce text.
	Col    int    // The column of the first invalid code point in the source text.
	Expect string // The expected syntax.
	Found  string // What was found.
}

// Error returns a string describing the error.
func (err Error) Error() string {
	if len(err.Expect) > 0 && len(err.Found) > 0 {
		return fmt.Sprintf("expected %v, found %v", err.Expect, err.Found)
	} else if len(err.Expect) > 0 {
		return fmt.Sprintf("expected %v", err.Expect)
	} else if len(err.Found) > 0 {
		return fmt.Sprintf("found %v", err.Found)
	} else {
		return "unknown compiler error"
	}
}

// Syntax error messages that can be produced by the compiler. There are two
// forms. The first form are those prefixed with Expect, which indicate that
// something was expected by not found. The second form are those prefixed with
// Found, which indicate that something was found but not expected.
const (
	ExpectLFAfterCR string = "linefeed (U+000A) after carriage return (U+000D)"

	FoundEOF = "EOF"
)
