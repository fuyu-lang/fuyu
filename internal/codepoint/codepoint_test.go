package codepoint

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"testing"
)

func TestU(t *testing.T) {
	data := map[rune](string){
		rune(0x0000):     "U+0000",
		rune(0xFFFF):     "U+FFFF",
		rune(0x10000):    "U+010000",
		rune(0xFFFFFF):   "U+FFFFFF",
		rune(0x1000000):  "U+01000000",
		rune(0x12345678): "U+12345678",
	}
	for r, expect := range data {
		chx.Eq(t, expect, u(r))
	}
}

func TestHumanize(t *testing.T) {
	data := map[rune](string){
		'\t':         "tab (U+0009)",
		'\n':         "line feed (U+000A)",
		'\r':         "carriage return (U+000D)",
		' ':          "space (U+0020)",
		rune(0x00A0): "U+00A0",
		'ϕ':          "ϕ",
		'\x01':       "U+0001",
	}
	for r, expect := range data {
		chx.Eq(t, expect, Humanize(r))
	}
}
