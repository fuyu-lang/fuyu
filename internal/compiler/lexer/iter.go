package lexer

import (
	"github.com/fuyu-lang/fuyu/internal/compiler/token"
)

// NextToken advances the lexer and produces the next token in the source.
//
// When a token is found, the tok return value is valid, found is set to true,
// and err is nil.
//
// When there are no more tokens in the source text, the tok return value is
// not valid, found is set to false, and err is nil.
//
// When an error occurs, the tok return value is not valid, found is true, and
// err is non-nil.
//
// When there are no more tokens (or errors), found will be false.
func (lexer *Lexer) NextToken() (tok token.Token, found bool, err error) {
	// TODO Implement this
	return token.Token{}, false, nil
}
