package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/source"
	"testing"
)

func TestByteOrderMark(t *testing.T) {
	if byteOrderMark != 0xFEFF {
		t.Errorf("incorrect byte order mark")
	}
}

func TestMakeLexer(t *testing.T) {
	const s string = "x = a + b * c"
	text, _ := source.MakeText(s)
	lexer := MakeLexer(*text)
	expected := Lexer{
		text:  text.Text(),
		index: 0, line: 1, col: 1,
		markIndex: 0, markLine: 1, markCol: 1,
	}
	if *lexer != expected {
		t.Errorf("expected %#v; got %#v", expected, lexer)
	}
}

func TestMakeLexerWithBom(t *testing.T) {
	const s string = "\ufeffx = a + b * c"
	text, _ := source.MakeText(s)
	lexer := MakeLexer(*text)
	expected := Lexer{
		text:  text.Text(),
		index: 3, line: 1, col: 1,
		markIndex: 3, markLine: 1, markCol: 1,
	}
	if *lexer != expected {
		t.Errorf("expected %#v; got %#v", expected, lexer)
	}
}

func TestMark(t *testing.T) {
	const s string = "x = a + b * c\ny = d * e + f"
	text, _ := source.MakeText(s)
	lexer := MakeLexer(*text)
	for i := 0; i < 17; i++ {
		lexer.nextRune()
	}

	var expected Lexer
	expected = Lexer{
		text:  text.Text(),
		index: 17, line: 2, col: 4,
		markIndex: 0, markLine: 1, markCol: 1,
	}
	if *lexer != expected {
		t.Errorf("expected %#v; got %#v", expected, lexer)
	}

	lexer.mark()
	expected = Lexer{
		text:  text.Text(),
		index: 17, line: 2, col: 4,
		markIndex: 17, markLine: 2, markCol: 4,
	}
	if *lexer != expected {
		t.Errorf("expected %#v; got %#v", expected, lexer)
	}

	lexer.nextRune()
	expected = Lexer{
		text:  text.Text(),
		index: 18, line: 2, col: 5,
		markIndex: 17, markLine: 2, markCol: 4,
	}
	if *lexer != expected {
		t.Errorf("expected %#v; got %#v", expected, lexer)
	}
}
