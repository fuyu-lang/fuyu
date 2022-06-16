package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"testing"
)

func TestEmpty(t *testing.T) {
	chx.Nil(t, checkUtf8(""))
}

func TestNonEmpty(t *testing.T) {
	chx.Nil(t, checkUtf8("abc\ndef\ng h i"))
}

func TestNonAscii(t *testing.T) {
	// Test 1, 2, 3, and 4 byte sequences
	chx.Nil(t, checkUtf8("aḅḈ\n1₂¾\nαβγ\n🍉"))
}

func TestInvalid(t *testing.T) {
	err := *checkUtf8("abc\nde\xff\nghi").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 6, Line: 2, Col: 3,
		Expect: "", Found: "invalid UTF-8 at byte 6",
	}
	chx.Eq(t, errExpect, err)
}
