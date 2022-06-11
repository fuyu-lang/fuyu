package token

// Kind is an enumeration representing the kind of a token.
type Kind uint

const (
	Illegal Kind = iota // Not allowed

	Comment // # ...

	Space   // A space (U+20)
	Newline // An LF (U+0A) or CRLF (U+0D U+0A)

	marker_identifier_begin
	Const       // Ident
	ConstHidden // _Ident
	Field       // @ident
	FieldStatic // @@ident
	Sink        // _
	Var         // ident
	VarHidden   // _ident
	marker_identifier_end

	marker_keyword_begin
	Always  // always
	As      // as
	Break   // break
	Class   // class
	Def     // def
	Do      // do
	Elif    // elif
	Else    // else
	For     // for
	From    // from
	If      // if
	In      // in
	Loop    // loop
	Panic   // panic
	Pass    // pass
	Recover // recover
	Redo    // redo
	Retry   // retry
	Return  // return
	Self    // self
	Then    // then
	Try     // try
	Unless  // unless
	Until   // until
	Where   // where
	While   // while
	marker_keyword_end

	marker_literal_begin
	Bytes             // b'...'
	Complex           // 123i
	Float             // 1.23
	Int               // 123, 0b101, 0o127 0x1f
	Nil               // nil
	StringDouble      // "..."
	StringDoubleBlock // """..."""
	StringSingle      // '...'
	StringSingleBlock // '''...'''
	Symbol            // :abc
	marker_literal_end

	marker_binary_begin
	Add        // +
	Assign     // =
	BitAnd     // &&&
	BitOr      // |||
	BitXor     // ^^^
	Compare    // <=>
	Concat     // ++
	Div        // /
	DivFloor   // //
	Eq         // ==
	Exp        // **
	Ge         // >=
	Gt         // >
	Le         // <=
	LogicAnd   // &&
	LogicOr    // ||
	Lt         // <
	Mod        // %
	Mul        // *
	Ne         // !=
	RangeExc   // ..
	RangeInc   // ..=
	ShiftLeft  // <<<
	ShiftRight // >>>
	Sub        // -
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
	Comma    // ,
	Cont     // \
	Dot      // .
	LBrace   // {
	LBrack   // [
	LParen   // (
	Pipe     // |
	PipeTo   // |>
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
	return marker_identifier_begin < k && k < marker_identifier_end
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

// IsConstant reports whether an identifier token kind is constant.
func (ident Kind) IsConstant() bool {
	return ident == Const || ident == ConstHidden
}

// IsField reports whether an identifier token kind is a field.
func (ident Kind) IsField() bool {
	return ident == Field || ident == FieldStatic
}

// IsHidden reports whether an identifier token kind is hidden.
func (ident Kind) IsHidden() bool {
	return ident == ConstHidden || ident == VarHidden
}
