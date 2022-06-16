package source

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	const source string = ""
	text, err := MakeText(source)
	if text.Text() != source {
		t.Error("incorrect text stored")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestNonEmpty(t *testing.T) {
	const source string = "abc\ndef\ng h i"
	text, err := MakeText(source)
	if text.Text() != source {
		t.Error("incorrect text stored")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestInvalid(t *testing.T) {
	const source string = "abc\nde\xff\nghi"
	text, err := MakeText(source)
	if text != nil {
		t.Error("text should be nil")
	}
	if err == nil {
		t.Error("expect an error")
	}
	if err.Error() != "invalid UTF-8 sequence at byte 6" {
		t.Error("incorrect error message:", err.Error())
	}
	e := err.(*Invalid)
	if e.Text() != source {
		t.Error("incorrect text stored")
	}
	if e.Index() != 6 {
		t.Error("expected error at index 6")
	}
	if e.Line() != 2 {
		t.Error("expected error at line 2")
	}
	if e.Col() != 3 {
		t.Error("expected error at column 3")
	}
}
