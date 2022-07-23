package chx

import (
	"testing"
)

func TestNil(t *testing.T) {
	Nil(t, nil)
}

func TestNotNil(t *testing.T) {
	x := 5
	p := &x
	NotNil(t, p)
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}

func TestEq(t *testing.T) {
	Eq(t, 5, 5)
}

func TestNe(t *testing.T) {
	Ne(t, 5, 6)
}
