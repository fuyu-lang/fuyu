package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/source"
	"testing"
	"unicode/utf8"
)

// This is the test string to iterate over in terms of runes. It contains 1, 2,
// 3, and 4 byte code points in UTF-8.
const s string = "aḅḈ\n1₂¾\nαβγ\n🍉"

var length int = utf8.RuneCountInString(s)

func TestMoreRunes(t *testing.T) {
	text, _ := source.MakeText(s)
	lexer := MakeLexer(*text)
	for i := 0; i < length; i++ {
		if !lexer.moreRunes() {
			t.Errorf("expected %v more runes", length-i)
		}
		lexer.nextRune()
	}
	if lexer.moreRunes() {
		t.Errorf("did not expect more runes")
	}
}

func TestNextRune(t *testing.T) {
	text, _ := source.MakeText(s)
	lexer := MakeLexer(*text)
	expect := func(r rune, index, line, col int) {
		t.Helper()
		next := lexer.nextRune()
		if next != r {
			t.Errorf("expected rune %v; got %v", r, next)
		}
		// Index, line, and column after the call to next
		if lexer.index != index {
			t.Errorf("expected index %v; got %v", index, lexer.index)
		}
		if lexer.line != line {
			t.Errorf("expected line %v; got %v", line, lexer.line)
		}
		if lexer.col != col {
			t.Errorf("expected col %v; got %v", col, lexer.col)
		}
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
