// Package token defines the lexical tokens of the language.
package token

// Token is the type of and location of a token in a source.
type Token struct {
	Kind  Kind // The kind of the token.
	Index int  // The start index of the token.
	End   int  // The index immediately following the last byte of the token.
}

// Kind is an enumeration representing the kind of a token.
type Kind int

const (
	Comment Kind = iota // # ...

	Space   // A space (U+20)
	Newline // An LF (U+0A) or CRLF (U+0D U+0A)

	marker_ident_begin
	IdentArg           // @1
	IdentConst         // Ident
	IdentField         // @ident
	IdentMethodPost    // map!, empty?
	IdentMethodSpecial // [], []=, unary+, binary+, reverse+
	IdentSink          // _
	IdentStrand        // @
	IdentVar           // ident
	marker_ident_end

	marker_keyword_begin
	Abstract   // abstract
	Always     // always
	As         // as
	Async      // async
	Await      // wait
	Break      // break
	Cons       // cons
	Def        // def
	Do         // do
	Elif       // elif
	Else       // else
	If         // if
	Import     // import
	In         // in
	Let        // let
	Panic      // panic
	Pass       // pass
	Pub        // pub
	Recover    // recover
	Redo       // redo
	Retry      // retry
	Return     // return
	SelfBig    // Self
	SelfLittle // self
	Then       // then
	Trait      // trait
	Try        // try
	Type       // type
	Unless     // unless
	marker_keyword_end

	marker_literal_begin
	Bytes             // b'...'
	Complex           // 123i, 1.23i, 1.2e-3i
	Float             // 1.23, 1.2e-3
	Int               // 123, 0b101, 0o127 0x1f
	Nil               // nil
	StringDouble      // "..."
	StringDoubleBlock // """..."""
	StringSingle      // '...'
	StringSingleBlock // '''...'''
	Symbol            // $abc
	marker_literal_end

	marker_binary_begin
	Add              // +
	AddAssign        // +=
	Assign           // =
	BitAnd           // &&&
	BitAndAssign     // &&&=
	BitOr            // |||
	BitOrAssign      // |||=
	BitXor           // ^^^
	BitXorAssign     // ^^^=
	Coalesce         // ??
	CoalesceAssign   // ?=
	Compare          // <=>
	Concat           // ++
	ConcatAssign     // ++=
	Div              // /
	DivAssign        // /=
	DivFloor         // //
	DivFloorAssign   // //=
	Eq               // ==
	Exp              // **
	ExpAssign        // **=
	Ge               // >=
	GroupAssign      // :=
	Gt               // >
	Le               // <=
	LogicAnd         // &&
	LogicOr          // ||
	Lt               // <
	Mod              // %
	ModAssign        // %=
	Mul              // *
	MulAssign        // *=
	Ne               // !=
	RangeExc         // ..
	RangeInc         // ..=
	ShiftLeft        // <<<
	ShiftLeftAssign  // <<<=
	ShiftRight       // >>>
	ShiftRightAssign // >>>=
	Sub              // -
	SubAssign        // -=
	marker_binary_end

	marker_unary_begin
	BitNot    // ~~~
	Block     // &
	BlockBang // &!
	LogicNot  // !
	Neg       // -
	Pos       // +
	Splat     // *
	SplatMap  // **
	marker_unary_end

	marker_punctuation_begin
	Arrow    // =>
	Bang     // !
	Colon    // :
	Dot      // .
	LBrace   // {
	LBrack   // [
	LParen   // (
	Pipe     // |
	PipeTo   // |<>
	RBrace   // }
	RBrack   // ]
	Rparen   // )
	SplatAll // ***
	marker_punctuation_end
)

// IsComment reports whether a token kind is a comment.
func (k Kind) IsComment() bool {
	return k == Comment
}

// IsSpace reports whether a token kind is a space.
func (k Kind) IsSpace() bool {
	return k == Space
}

// IsNewline reports whether a token kind is a newline.
func (k Kind) IsNewline() bool {
	return k == Newline
}

// IsIdent reports whether a token kind is an identifier.
func (k Kind) IsIdent() bool {
	return marker_ident_begin < k && k < marker_ident_end
}

// IsKeyword reports whether a token kind is a keyword.
func (k Kind) IsKeyword() bool {
	return marker_keyword_begin < k && k < marker_keyword_end
}

// IsLiteral reports whether a token kind is a literal value.
func (k Kind) IsLiteral() bool {
	return marker_literal_begin < k && k < marker_literal_end
}

// IsBinary reports whether a token kind is a binary operator.
func (k Kind) IsBinary() bool {
	return marker_binary_begin < k && k < marker_binary_end
}

// IsUnary reports whether a token kind is a unary operator.
func (k Kind) IsUnary() bool {
	return marker_unary_begin < k && k < marker_unary_end
}

// IsPunc reports whether a token kind is a punctuation.
func (k Kind) IsPunc() bool {
	return marker_punctuation_begin < k && k < marker_punctuation_end
}
