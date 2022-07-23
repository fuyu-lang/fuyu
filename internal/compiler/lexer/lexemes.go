package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/codepoint"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"github.com/fuyu-lang/fuyu/internal/compiler/token"
)

// makeToken builds a token of a specific kind from the lexer. The error is
// always always nil, and is present for compatibility with other methods.
func (lexer *Lexer) makeToken(kind token.Kind) (token.Token, error) {
	t := token.Token{Kind: kind, Index: lexer.markIndex, End: lexer.index}
	return t, nil
}

// makeError builds an error from the state of the lexer. The token always has
// the kind Illegal.
func (lexer *Lexer) makeError(expect, found string) (token.Token, error) {
	err := compiler.Error{
		Index: lexer.index, Line: lexer.line, Col: lexer.col,
		Expect: expect, Found: found,
	}
	return token.Token{}, err
}

// makeErrorExpect is like makeError but does not set the found message.
func (lexer *Lexer) makeErrorExpect(expect string) (token.Token, error) {
	return lexer.makeError(expect, "")
}

// TODO Use this
// makeErrorFound is like makeError but does not set the expected message.
func (lexer *Lexer) makeErrorFound(found string) (token.Token, error) {
	return lexer.makeError("", found)
}

// takeIf advances the lexer if the condition is not satisifed. If there are no
// more runes in the lexer then nothing happens.
func (lexer *Lexer) takeIf(cond func(rune) bool) {
	if lexer.moreRunes() && cond(lexer.peekRune()) {
		lexer.nextRune()
	}
}

// takeWhile advances the lexer until either there are no more runes in the
// lexer or the condition is not satisifed.
func (lexer *Lexer) takeWhile(cond func(rune) bool) {
	for ; lexer.moreRunes() && cond(lexer.peekRune()); lexer.nextRune() {
	}
}

// lexComment assumes that the lexer is sitting on the first rune of a comment
// and consumes the rest of the current line, not including the newline.
func (lexer *Lexer) lexComment() (token.Token, error) {
	lexer.takeWhile(func(r rune) bool { return r != '\r' && r != '\n' })
	return lexer.makeToken(token.Comment)
}

// lexSpace assumes that the lexer is sitting on a space and consumes up to the
// first non-space rune.
func (lexer *Lexer) lexSpace() (token.Token, error) {
	lexer.takeWhile(func(r rune) bool { return r == ' ' })
	return lexer.makeToken(token.Space)
}

// lexNewline assumes that the lexer is sitting on a newline and consumes one
// newline. A newline is considered to be a linefeed (U+000A) optionally
// preceded by a carriage return (U+000D).
func (lexer *Lexer) lexNewline() (token.Token, error) {
	if lexer.nextRune() == '\r' {
		if lexer.moreRunes() {
			c := lexer.peekRune()
			if c != '\n' {
				return lexer.makeError(
					compiler.ExpectLFAfterCR, codepoint.Humanize(c))
			} else {
				lexer.nextRune()
			}
		} else {
			return lexer.makeError(compiler.ExpectLFAfterCR, compiler.FoundEOF)
		}
	}
	return lexer.makeToken(token.Newline)
}

// lexWord assumes that the lexer is sitting on the first rune of a word token,
// which are identifiers, keywords, boolean literals, and nil literals, and
// consumes the entire word.
func (lexer *Lexer) lexWord() (token.Token, error) {
	// TODO Classify as kind: Identifier, keyword, bool, nil
	panic("Not implemented")
}

// lexNumber assumes that the lexer is sitting on the first rune of a number
// literal and consumes the entire number.
func (lexer *Lexer) lexNumber() (token.Token, error) {
	if lexer.peekRune() == '0' {
		lexer.nextRune()
		if lexer.moreRunes() {
			r := lexer.peekRune()
			switch r {
			case 'b':
				lexer.nextRune()
				return lexBin(lexer)
			case 'o':
				lexer.nextRune()
				return lexOct(lexer)
			case 'x':
				lexer.nextRune()
				return lexHex(lexer)
				// TODO Default or is that implicit?
			}
		}
	}
	return lexDec(lexer)
}

// lexBin consumes the rest of a binary integer literal after the 0b.
func lexBin(lexer *Lexer) (token.Token, error) {
	seenDigit := false
	lexer.takeWhile(func(r rune) bool {
		if r == '0' || r == '1' {
			seenDigit = true
			return true
		}
		return r == '_'
	})
	if seenDigit {
		// TODO Return token.IntBin?
		return lexer.makeToken(token.Int)
	}
	return lexer.makeErrorExpect(compiler.ExpectDigitInInt)
}

// lexOct consumes the rest of a octal integer literal after the 0c.
func lexOct(lexer *Lexer) (token.Token, error) {
	seenDigit := false
	lexer.takeWhile(func(r rune) bool {
		if '0' <= r && r <= '7' {
			seenDigit = true
			return true
		}
		return r == '_'
	})
	if seenDigit {
		// TODO Return token.IntOct?
		return lexer.makeToken(token.Int)
	}
	return lexer.makeErrorExpect(compiler.ExpectDigitInInt)
}

// lexHex consumes the rest of a hexdecimal integer literal after the 0x.
func lexHex(lexer *Lexer) (token.Token, error) {
	seenDigit := false
	lexer.takeWhile(func(r rune) bool {
		if ('0' <= r && r <= '9') || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F') {
			seenDigit = true
			return true
		}
		return r == '_'
	})
	if seenDigit {
		// TODO Return token.IntHex?
		return lexer.makeToken(token.Int)
	}
	return lexer.makeErrorExpect(compiler.ExpectDigitInInt)
}

// TODO Docs
func lexDec(lexer *Lexer) (token.Token, error) {
	panic("Not implemented") // TODO
}

// lexString assumes that the lexer is sitting on the first rune of a string
// literal and consumes the entire string.
//
// Strings which support interpolation must be processed separately to access
// their sub-tokens.
func (lexer *Lexer) lexString() (token.Token, error) {
	panic("Not implemented") // TODO
}

// lexBytes assumes that the lexer is sitting on the first rune of a bytes
// literal and consumes the entire bytes literal.
func (lexer *Lexer) lexBytes() (token.Token, error) {
	panic("Not implemented") // TODO
}

// lexSymbol assumes that the lexer is sitting on the first rune of a symbol
// literal and consumes the entire symbol.
func (lexer *Lexer) lexSymbol() (token.Token, error) {
	panic("Not implemented") // TODO
}

// lexOpPunc assumes that the lexer is sitting on the first rune of an operator
// or punctuation and consumes the entire operator or punctuation.
func (lexer *Lexer) lexOpPunc() (token.Token, error) {
	panic("Not implemented") // TODO
}
