// Package lexer provides a lexer for producing tokens of a Fuyu source text.
package lexer

import (
	"unicode/utf8"
)

// The value of a byte order mark.
const byteOrderMark rune = 0xFEFF

// Lexer is a lexer over some source text which provide tokens in the source.
type Lexer struct {
	text      string // A valid UTF-8 encoded source text.
	index     int    // The current index of the lexer in bytes.
	line      int    // The current line of the lexer.
	col       int    // The current column of the lexer.
	markIndex int    // The marked index of the lexer in bytes.
	markLine  int    // The marked line of the lexer.
	markCol   int    // The markd column of the lexer.
}

// MakeLexer creates a new lexer from a source text. This fails when the input
// text is not valid UTF-8.
func MakeLexer(text string) (*Lexer, error) {
	err := checkUtf8(text)
	if err != nil {
		return nil, err
	}
	lexer := &Lexer{
		text:  text,
		index: 0, line: 1, col: 1,
		markIndex: 0, markLine: 1, markCol: 1,
	}
	r, size := utf8.DecodeRuneInString(lexer.text)
	if r == byteOrderMark {
		lexer.index = size
		lexer.mark()
	}
	return lexer, nil
}

// mark stores the current index, line, and column for later use.
func (lexer *Lexer) mark() {
	lexer.markIndex = lexer.index
	lexer.markLine = lexer.line
	lexer.markCol = lexer.col
}
