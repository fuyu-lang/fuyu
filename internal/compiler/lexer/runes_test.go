package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"testing"
	"unicode/utf8"
)

// This is the test string to iterate over in terms of runes. It contains 1, 2,
// 3, and 4 byte code points in UTF-8.
const text string = "aḅḈ\n1₂¾\nαβγ\n🍉"

var length int = utf8.RuneCountInString(text)

func TestMoreRunes(t *testing.T) {
	lexer, _ := MakeLexer(text)
	for i := 0; i < length; i++ {
		chx.True(t, lexer.moreRunes())
		lexer.nextRune()
	}
	chx.False(t, lexer.moreRunes())
}

func TestNextRune(t *testing.T) {
	lexer, _ := MakeLexer(text)
	expect := func(r rune, index, line, col int) {
		t.Helper()
		next := lexer.nextRune()
		chx.Eq(t, r, next, "rune")
		chx.Eq(t, index, lexer.index, "index")
		chx.Eq(t, line, lexer.line, "line")
		chx.Eq(t, col, lexer.col, "col")
	}
	expect('a', 1, 1, 2)
	expect('ḅ', 4, 1, 3)
	expect('Ḉ', 7, 1, 4)
	expect('\n', 8, 2, 1)
	expect('1', 9, 2, 2)
	expect('₂', 12, 2, 3)
	expect('¾', 14, 2, 4)
	expect('\n', 15, 3, 1)
	expect('α', 17, 3, 2)
	expect('β', 19, 3, 3)
	expect('γ', 21, 3, 4)
	expect('\n', 22, 4, 1)
	expect('🍉', 26, 4, 2)
}

func TestPeekRune(t *testing.T) {
	lexer, _ := MakeLexer(text)
	expect := func(r rune, index, line, col int) {
		t.Helper()
		peek := lexer.peekRune()
		chx.Eq(t, r, peek, "rune")
		chx.Eq(t, index, lexer.index, "index")
		chx.Eq(t, line, lexer.line, "line")
		chx.Eq(t, col, lexer.col, "col")
		lexer.nextRune()
	}
	expect('a', 0, 1, 1)
	expect('ḅ', 1, 1, 2)
	expect('Ḉ', 4, 1, 3)
	expect('\n', 7, 1, 4)
	expect('1', 8, 2, 1)
	expect('₂', 9, 2, 2)
	expect('¾', 12, 2, 3)
	expect('\n', 14, 2, 4)
	expect('α', 15, 3, 1)
	expect('β', 17, 3, 2)
	expect('γ', 19, 3, 3)
	expect('\n', 21, 3, 4)
	expect('🍉', 22, 4, 1)
}
