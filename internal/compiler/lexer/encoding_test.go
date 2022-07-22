package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"github.com/fuyu-lang/fuyu/internal/compiler"
	"testing"
)

func TestCheckRune(t *testing.T) {
	// Test equivalence classes
	chx.NotNil(t, checkRune('\u0000'))
	chx.NotNil(t, checkRune('\u0009'))
	chx.Nil(t, checkRune('\u000A'))
	chx.NotNil(t, checkRune('\u000B'))
	chx.NotNil(t, checkRune('\u000C'))
	chx.Nil(t, checkRune('\u000D'))
	chx.NotNil(t, checkRune('\u000E'))
	chx.NotNil(t, checkRune('\u001F'))
	chx.Nil(t, checkRune('\u0020'))
	chx.Nil(t, checkRune('\u007E'))
	chx.NotNil(t, checkRune('\u007F'))
	chx.NotNil(t, checkRune('\u009F'))
	chx.Nil(t, checkRune('\u00A0'))

	// Test the limits of each codepoint size in bytes
	chx.NotNil(t, checkRune('\u0000'))
	chx.NotNil(t, checkRune('\u007F'))
	chx.NotNil(t, checkRune('\u0080'))
	chx.Nil(t, checkRune('\u07FF'))
	chx.Nil(t, checkRune('\u0800'))
	chx.Nil(t, checkRune('\uFFFF'))
	chx.Nil(t, checkRune('\U00100000'))
	chx.Nil(t, checkRune('\U0010FFFF'))
}

func TestCheckUtf8Empty(t *testing.T) {
	chx.Nil(t, checkUtf8(""))
}

func TestCheckUtf8NonEmpty(t *testing.T) {
	chx.Nil(t, checkUtf8("abc\ndef\ng h i"))
}

func TestCheckUtf8NonAscii(t *testing.T) {
	// Test 1, 2, 3, and 4 byte sequences
	chx.Nil(t, checkUtf8("aḅḈ\n1₂¾\nαβγ\n🍉"))
}

func TestCheckUtf8Invalid(t *testing.T) {
	err := *checkUtf8("abc\nde\xff\nghi").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 6, Line: 2, Col: 3,
		Expect: "", Found: "invalid bit sequence in UTF-8 string",
	}
	chx.Eq(t, errExpect, err)
}

func TestValidSourceEmcodingEmpty(t *testing.T) {
	chx.Nil(t, validSourceEncoding(""))
}

func TestValidSourceEmcodingNotEmpty(t *testing.T) {
	chx.Nil(t, validSourceEncoding("abc\ndef\ng h i"))
}

func TestValidSourceEncodingNonAscii(t *testing.T) {
	// Test 1, 2, 3, and 4 byte sequences
	chx.Nil(t, validSourceEncoding("aḅḈ\n1₂¾\nαβγ\n🍉"))
}

func TestValidSourceEncodingInvalid(t *testing.T) {
	err := *validSourceEncoding("abc\nde\xff\nghi").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 6, Line: 2, Col: 3,
		Expect: "", Found: "invalid bit sequence in UTF-8 string",
	}
	chx.Eq(t, errExpect, err)
}

func TestValidSourceEncodingInvalidMissingLinefeed(t *testing.T) {
	err := *validSourceEncoding("\rx").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 1, Line: 1, Col: 2,
		Expect: "line feed (U+000A) after carriage return (U+000D)", Found: "x",
	}
	chx.Eq(t, errExpect, err)
}

func TestValidSourceEncodingInvalidMissingLinefeedAtEnd(t *testing.T) {
	err := *validSourceEncoding("x\r").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 2, Line: 1, Col: 3,
		Expect: "line feed (U+000A) after carriage return (U+000D)", Found: "EOF",
	}
	chx.Eq(t, errExpect, err)
}

func TestValidSourceEncodingInvalidIllegalCodepoint(t *testing.T) {
	err := *validSourceEncoding("abc\tdef").(*compiler.Error)
	errExpect := compiler.Error{
		Index: 3, Line: 1, Col: 4,
		Found: "invalid UTF-8 codepoint tab (U+0009)",
	}
	chx.Eq(t, errExpect, err)
}
