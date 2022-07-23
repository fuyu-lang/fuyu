package lexer

import (
	"unicode/utf8"
)

// moreRunes checks that there are more runes available in the lexer.
func (lexer *Lexer) moreRunes() bool {
	return lexer.index < len(lexer.text)
}

// nextRune gets the next rune in the lexer and advances it. This should never
// be called without first checking moreRunes.
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

// peekRune is like nextRune but does not advance the lexer.
func (lexer *Lexer) peekRune() rune {
	r, _ := utf8.DecodeRuneInString(lexer.text[lexer.index:])
	return r
}

// peekNthRune peeks ahead n runes, starting at 1. Unlike peekRune, this is
// safe and called be called at any time without checking moreRunes.
func (lexer *Lexer) peekNthRune(n int) rune {
	index := lexer.index
	var r rune = 0
	var size int
	for i := 0; i < n; i++ {
		r, size = utf8.DecodeRuneInString(lexer.text[index:])
		if size == 0 {
			return 0
		}
		index += size
	}
	return r
}
