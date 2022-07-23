// Package chx provides simplified interfaces for making assertions and
// printing failure information in unit tests.
package chx

import (
	"fmt"
	"testing"
)

// makeNote prepares the optional note that is printed with a failed test.
func makeNote(note []interface{}) string {
	if len(note) == 0 {
		return "no note given"
	}
	noteString, isString := note[0].(string)
	if !isString {
		return fmt.Sprintf("note was not a string, got: %T", note[0])
	}
	if len(note) == 1 {
		return noteString
	}
	return fmt.Sprintf(noteString, note[1:]...)
}

// fmtExpect formats a message showing expected and actual values along with an
// optional note to display for failed tests.
func fmtExpect(expect, value interface{}, note []interface{}) string {
	s := fmt.Sprintf("expected %T %#v, got %T %#v", expect, expect, value, value)
	if len(note) > 0 {
		s = fmt.Sprintf("%s (%s)", s, makeNote(note))
	}
	return s
}

// fmtExpectDesc is like fmtExpect except it formats the contents of the the
// expected string as a descriptive text, not a literal value.
func fmtExpectDesc(expect string, value interface{}, note []interface{}) string {
	s := fmt.Sprintf("expected %s, got %T %#v", expect, value, value)
	if len(note) > 0 {
		s = fmt.Sprintf("%s (%s)", s, makeNote(note))
	}
	return s
}

// Nil asserts that a value is nil.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func Nil(t *testing.T, value interface{}, note ...interface{}) {
	t.Helper()
	if value != nil {
		t.Errorf(fmtExpect(nil, value, note))
	}
}

// NotNil asserts that a value is not nil.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func NotNil(t *testing.T, value interface{}, note ...interface{}) {
	t.Helper()
	if value == nil {
		t.Errorf(fmtExpectDesc("not <nil>", value, note))
	}
}

// True asserts that a value is true.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func True(t *testing.T, value bool, note ...interface{}) {
	t.Helper()
	if !value {
		t.Errorf(fmtExpect(true, value, note))
	}
}

// False asserts that a value is false.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func False(t *testing.T, value bool, note ...interface{}) {
	t.Helper()
	if value {
		t.Errorf(fmtExpect(false, value, note))
	}
}

// Eq asserts that a expected and actual values are equal by value.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func Eq(t *testing.T, expect, value interface{}, note ...interface{}) {
	t.Helper()
	if expect != value {
		t.Errorf(fmtExpect(expect, value, note))
	}
}

// Ne asserts that a expected and actual values are inequal by value.
//
// An optional note can be provided to display when the assertion fails. The
// arguments of the note are passed to fmt.Sprintf.
func Ne(t *testing.T, expect, value interface{}, note ...interface{}) {
	t.Helper()
	if expect == value {
		t.Errorf(fmtExpect(expect, value, note))
	}
}
