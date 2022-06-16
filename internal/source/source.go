// Package source is used to represent a valid UTF-8 encoded source text. It
// does not require that the text be a valid Fuyu program, simply that the text
// is guaranteed to be valid UTF-8.
package source

import (
	"fmt"
	"unicode/utf8"
)

// A UTF-8 encoded source text.
type Text struct {
	text string
}

// Make a new source text by checking whether it is valid UTF-8.
func MakeText(maybeText string) (*Text, error) {
	index, line, col := 0, 1, 1
	for index < len(maybeText) {
		r, size := utf8.DecodeRuneInString(maybeText[index:])
		if r == utf8.RuneError {
			if size == 0 {
				// This is fine because it is an empty string. This will also never get hit
				// because the loop will stop before this.
			} else if size == 1 {
				return nil, &Invalid{
					text:  maybeText,
					index: index, line: line, col: col,
				}
			} else {
				// This is a valid rune with a value equal to utf8.RuneError
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
	return &Text{maybeText}, nil
}

// Text gets the valid UTF-8 text.
func (text Text) Text() string {
	return text.text
}

// Invalid represents an error where a source text is invalid.
type Invalid struct {
	text  string // The source text.
	index int    // The index of the first invalid byte in the source text.
	line  int    // The line of the first invalid byte in the souce text.
	col   int    // The column of the first invalid byte in the source text.
}

// Error returns a string describing the error.
func (err Invalid) Error() string {
	return fmt.Sprintf("invalid UTF-8 sequence at byte %v", err.index)
}

// Text gets the invalid source text.
func (err Invalid) Text() string {
	return err.text
}

// Index gets the index of the first invalid byte in the source text.
func (err Invalid) Index() int {
	return err.index
}

// Line gets the line of the first invalid byte in the source text.
func (err Invalid) Line() int {
	return err.line
}

// Line gets the column line of the first invalid byte in the source text.
func (err Invalid) Col() int {
	return err.col
}
