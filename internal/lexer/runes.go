package lexer

import (
	"unicode/utf8"
)

// moreRunes checks that there are more runes available in the lexer.
func (lexer *Lexer) moreRunes() bool {
	return lexer.index < len(lexer.text)
}

// nextRune gets the next run in the lexer. This should never be called without
// first checking moreRunes.
func (lexer *Lexer) nextRune() rune {
	// Since the source text is known to be valid, the error is never checked.
	// This function will also never be called without checking moreRunes, so it
	// is not necessary to check that the string is empty.
	r, size := utf8.DecodeRuneInString(lexer.text[lexer.index:])
	lexer.index += size
	if r == '\n' {
		lexer.line += 1
		lexer.col = 1
	} else {
		lexer.col += 1
	}
	return r
}
