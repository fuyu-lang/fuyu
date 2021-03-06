// Package token defines the lexical tokens of the language.
package token

// Token is the type of and location of a token in a source.
type Token struct {
	Kind  Kind // The kind of the token.
	Index uint // The start index of the token.
	End   uint // The index immediately following the last byte of the token.
}
