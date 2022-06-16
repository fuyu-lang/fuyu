package compiler

import (
	"github.com/fuyu-lang/fuyu/internal/chx"
	"testing"
)

func TestErrorExpectFound(t *testing.T) {
	err := Error{
		Index: 50, Line: 5, Col: 8,
		Expect: "exp", Found: "fnd",
	}
	chx.Eq(t, "expected exp, found fnd", err.Error())
}

func TestErrorExpectOnly(t *testing.T) {
	err := Error{
		Index: 50, Line: 5, Col: 8,
		Expect: "exp", Found: "",
	}
	chx.Eq(t, "expected exp", err.Error())
}

func TestErrorFoundOnly(t *testing.T) {
	err := Error{
		Index: 50, Line: 5, Col: 8,
		Expect: "", Found: "fnd",
	}
	chx.Eq(t, "found fnd", err.Error())
}

func TestErrorNoInfo(t *testing.T) {
	err := Error{
		Index: 50, Line: 5, Col: 8,
		Expect: "", Found: "",
	}
	chx.Eq(t, "unknown compiler error", err.Error())
}
