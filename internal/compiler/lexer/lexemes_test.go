package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"github.com/fuyu-lang/fuyu/internal/compiler/token"
	"testing"
)

func TestMakeToken(t *testing.T) {
	lexer, _ := MakeLexer("a b c")
	lexer.nextRune()
	lexer.nextRune()
	lexer.mark()
	lexer.nextRune()
	tok, err := lexer.makeToken(token.IdentVar)
	tokExpect := token.Token{Kind: token.IdentVar, Index: 2, End: 3}
	chx.Eq(t, tokExpect, tok)
	chx.Nil(t, err)
}

func TestMakeError(t *testing.T) {
	lexer, _ := MakeLexer("a b c")
	lexer.nextRune()
	lexer.nextRune()
	lexer.mark()
	lexer.nextRune()
	exp, fnd := "exp", "fnd"
	tok, err := lexer.makeError(exp, fnd)
	errExpect := compiler.Error{
		Index: 3, Line: 1, Col: 4,
		Expect: exp, Found: fnd,
	}
	chx.Eq(t, token.Token{}, tok)
	chx.Eq(t, errExpect, err)
}

func TestMakeErrorExpect(t *testing.T) {
	lexer, _ := MakeLexer("a b c")
	lexer.nextRune()
	lexer.nextRune()
	lexer.mark()
	lexer.nextRune()
	exp := "exp"
	tok, err := lexer.makeErrorExpect(exp)
	errExpect := compiler.Error{
		Index: 3, Line: 1, Col: 4,
		Expect: exp, Found: "",
	}
	chx.Eq(t, token.Token{}, tok)
	chx.Eq(t, errExpect, err)
}

func TestMakeErrorFound(t *testing.T) {
	lexer, _ := MakeLexer("a b c")
	lexer.nextRune()
	lexer.nextRune()
	lexer.mark()
	lexer.nextRune()
	fnd := "fnd"
	tok, err := lexer.makeErrorFound(fnd)
	errExpect := compiler.Error{
		Index: 3, Line: 1, Col: 4,
		Expect: "", Found: fnd,
	}
	chx.Eq(t, token.Token{}, tok)
	chx.Eq(t, errExpect, err)
}

func TestTakeIf(t *testing.T) {
	text := "ab"
	lexer, _ := MakeLexer(text)
	lexerExpect := Lexer{
		text:  text,
		index: 1, line: 1, col: 2,
		markIndex: 0, markLine: 1, markCol: 1,
	}

	// This advances the lexer
	lexer.takeIf(func(r rune) bool { return r == 'a' })
	chx.Eq(t, lexerExpect, *lexer)

	// This does not advance the lexer
	lexer.takeIf(func(r rune) bool { return r == 'a' })
	chx.Eq(t, lexerExpect, *lexer)
}

func TestTakeWhile(t *testing.T) {
	text := "aaab"
	lexer, _ := MakeLexer(text)
	lexerExpect := Lexer{
		text:  text,
		index: 3, line: 1, col: 4,
		markIndex: 0, markLine: 1, markCol: 1,
	}

	// This advances the lexer
	lexer.takeWhile(func(r rune) bool { return r == 'a' })
	chx.Eq(t, lexerExpect, *lexer)

	// This does not advance the lexer
	lexer.takeWhile(func(r rune) bool { return r == 'a' })
	chx.Eq(t, lexerExpect, *lexer)
}

// testLex runs a lexer operation on a source text and asserts that the
// specified token is found.
func testLex(
	t *testing.T,
	text string,
	op func(lexer *Lexer) (token.Token, error),
	tokExpect token.Token,
) {
	t.Helper()
	lexer, _ := MakeLexer(text)
	tok, err := op(lexer)
	chx.Eq(t, tokExpect, tok)
	chx.Nil(t, err)
}

// testLexError runs an lexer operation on a source text and asserts that the
// specified error is generated.
func testLexError(
	t *testing.T,
	text string,
	op func(lexer *Lexer) (token.Token, error),
	errExpect compiler.Error,
) {
	t.Helper()
	lexer, err := MakeLexer(text)
	if err != nil {
		e := err.(*compiler.Error)
		chx.Eq(t, errExpect, *e)
		return
	}
	tok, err := op(lexer)
	chx.Eq(t, token.Token{}, tok)
	chx.NotNil(t, err)
	chx.Eq(t, errExpect, err)
}

func TestLexComment(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexComment()
	}
	testLex(t, "# a b c\nx y z", op,
		token.Token{Kind: token.Comment, Index: 0, End: 7})
	testLex(t, "# a b c", op,
		token.Token{Kind: token.Comment, Index: 0, End: 7})
}

func TestLexSpace(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexSpace()
	}
	testLex(t, "    a b c", op,
		token.Token{Kind: token.Space, Index: 0, End: 4})
	testLex(t, "    ", op,
		token.Token{Kind: token.Space, Index: 0, End: 4})
}

func TestLexNewline(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexNewline()
	}
	testLex(t, "\nx", op,
		token.Token{Kind: token.Newline, Index: 0, End: 1})
	testLex(t, "\r\nx", op,
		token.Token{Kind: token.Newline, Index: 0, End: 2})
	testLexError(t, "\rx", op,
		compiler.Error{
			Index: 1, Line: 1, Col: 2,
			Expect: compiler.ExpectLFAfterCR, Found: "x",
		})
}

func TestLexNumberBin(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexNumber()
	}
	testLex(t, "0b0", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0b1", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0b0101", op,
		token.Token{Kind: token.Int, Index: 0, End: 6})
	testLex(t, "0b_0", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0b0_", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0b0____1", op,
		token.Token{Kind: token.Int, Index: 0, End: 8})
	testLexError(t, "0b", op,
		compiler.Error{
			Index: 2, Line: 1, Col: 3,
			Expect: compiler.ExpectDigitInInt,
		})
	testLexError(t, "0b_", op,
		compiler.Error{
			Index: 3, Line: 1, Col: 4,
			Expect: compiler.ExpectDigitInInt,
		})
}

func TestLexNumberOct(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexNumber()
	}
	testLex(t, "0o0", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0o7", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0o01234567", op,
		token.Token{Kind: token.Int, Index: 0, End: 10})
	testLex(t, "0o_0", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0o0_", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0o0____7", op,
		token.Token{Kind: token.Int, Index: 0, End: 8})
	testLexError(t, "0o", op,
		compiler.Error{
			Index: 2, Line: 1, Col: 3,
			Expect: compiler.ExpectDigitInInt,
		})
	testLexError(t, "0o_", op,
		compiler.Error{
			Index: 3, Line: 1, Col: 4,
			Expect: compiler.ExpectDigitInInt,
		})
}

func TestLexNumberHex(t *testing.T) {
	op := func(lexer *Lexer) (token.Token, error) {
		return lexer.lexNumber()
	}
	testLex(t, "0x0", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0xF", op,
		token.Token{Kind: token.Int, Index: 0, End: 3})
	testLex(t, "0x0123456789abcdefABCDEF", op,
		token.Token{Kind: token.Int, Index: 0, End: 24})
	testLex(t, "0x_0", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0x0_", op,
		token.Token{Kind: token.Int, Index: 0, End: 4})
	testLex(t, "0x0____f", op,
		token.Token{Kind: token.Int, Index: 0, End: 8})
	testLexError(t, "0x", op,
		compiler.Error{
			Index: 2, Line: 1, Col: 3,
			Expect: compiler.ExpectDigitInInt,
		})
	testLexError(t, "0x_", op,
		compiler.Error{
			Index: 3, Line: 1, Col: 4,
			Expect: compiler.ExpectDigitInInt,
		})
}

func TestLexNumberDec(t *testing.T) {
	// TODO
}
