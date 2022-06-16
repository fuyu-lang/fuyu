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

// testLex runs an lexer operation on a source text and asserts that the
// specified token is found.
func testLex(
	t *testing.T,
	text string,
	tokExpect token.Token,
	op func(lexer *Lexer) (token.Token, error),
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
	errExpect compiler.Error,
	op func(lexer *Lexer) (token.Token, error),
) {
	t.Helper()
	lexer, _ := MakeLexer(text)
	tok, err := op(lexer)
	chx.Eq(t, token.Token{}, tok)
	chx.NotNil(t, err)
	chx.Eq(t, errExpect, err)
}

func TestLexComment(t *testing.T) {
	testLex(t, "# a b c\nx y z",
		token.Token{Kind: token.Comment, Index: 0, End: 7},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexComment()
		})
	testLex(t, "# a b c",
		token.Token{Kind: token.Comment, Index: 0, End: 7},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexComment()
		})
}

func TestLexSpace(t *testing.T) {
	testLex(t, "    a b c",
		token.Token{Kind: token.Space, Index: 0, End: 4},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexSpace()
		})
	testLex(t, "    ",
		token.Token{Kind: token.Space, Index: 0, End: 4},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexSpace()
		})
}

func TestLexNewline(t *testing.T) {
	testLex(t, "\nx",
		token.Token{Kind: token.Newline, Index: 0, End: 1},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexNewline()
		})
	testLex(t, "\r\nx",
		token.Token{Kind: token.Newline, Index: 0, End: 2},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexNewline()
		})
	testLexError(t, "\rx",
		compiler.Error{
			Index: 1, Line: 1, Col: 2,
			Expect: compiler.ExpectLFAfterCR, Found: "x",
		},
		func(lexer *Lexer) (token.Token, error) {
			return lexer.lexNewline()
		})
}
