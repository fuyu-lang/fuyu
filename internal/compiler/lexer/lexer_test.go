package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"testing"
)

func TestByteOrderMark(t *testing.T) {
	chx.Eq(t, rune(0xFEFF), byteOrderMark)
}

func TestMakeLexer(t *testing.T) {
	const text string = "x = a + b * c"
	lexer, _ := MakeLexer(text)
	lexerExpect := Lexer{
		text:  text,
		index: 0, line: 1, col: 1,
		markIndex: 0, markLine: 1, markCol: 1,
	}
	chx.Eq(t, lexerExpect, *lexer)
}

func TestMakeLexerWithBom(t *testing.T) {
	const text string = "\ufeffx = a + b * c"
	lexer, _ := MakeLexer(text)
	lexerExpect := Lexer{
		text:  text,
		index: 3, line: 1, col: 1,
		markIndex: 3, markLine: 1, markCol: 1,
	}
	chx.Eq(t, lexerExpect, *lexer)
}

func TestMakeLexerInvalidUtf8(t *testing.T) {
	const text string = "x = a + b * c\ny = \xffd + e / f"
	_, err := MakeLexer(text)
	e := *err.(*compiler.Error)
	eExpect := compiler.Error{
		Index: 18, Line: 2, Col: 5,
		Expect: "", Found: "invalid bit sequence in UTF-8 string",
	}
	chx.Eq(t, eExpect, e)
}

func TestMark(t *testing.T) {
	const text string = "x = a + b * c\ny = d * e + f"
	lexer, _ := MakeLexer(text)
	for i := 0; i < 17; i++ {
		lexer.nextRune()
	}

	var lexerExpect Lexer
	lexerExpect = Lexer{
		text:  text,
		index: 17, line: 2, col: 4,
		markIndex: 0, markLine: 1, markCol: 1,
	}
	chx.Eq(t, lexerExpect, *lexer)

	lexer.mark()
	lexerExpect = Lexer{
		text:  text,
		index: 17, line: 2, col: 4,
		markIndex: 17, markLine: 2, markCol: 4,
	}
	chx.Eq(t, lexerExpect, *lexer)

	lexer.nextRune()
	lexerExpect = Lexer{
		text:  text,
		index: 18, line: 2, col: 5,
		markIndex: 17, markLine: 2, markCol: 4,
	}
	chx.Eq(t, lexerExpect, *lexer)
}
