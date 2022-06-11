package token

import "testing"

func TestIsComment(t *testing.T) {
	if Illegal.IsComment() {
		t.Error("Illegal should not be a comment")
	}
	if !Comment.IsComment() {
		t.Error("Comment should be a comment")
	}
	if Space.IsComment() {
		t.Error("Space should not be a comment")
	}
	if Newline.IsComment() {
		t.Error("Newline should not be a comment")
	}
	if Const.IsComment() {
		t.Error("Const should not be a comment")
	}
	if ConstHidden.IsComment() {
		t.Error("ConstHidden should not be a comment")
	}
	if Field.IsComment() {
		t.Error("Field should not be a comment")
	}
	if FieldStatic.IsComment() {
		t.Error("FieldStaticshould not be a comment")
	}
	if Sink.IsComment() {
		t.Error("Sink should not be a comment")
	}
	if Var.IsComment() {
		t.Error("Var should not be a comment")
	}
	if VarHidden.IsComment() {
		t.Error("VarHidden should not be a comment")
	}
	if Always.IsComment() {
		t.Error("Always should not be a comment")
	}
	if As.IsComment() {
		t.Error("As should not be a comment")
	}
	if Break.IsComment() {
		t.Error("Break should not be a comment")
	}
	if Class.IsComment() {
		t.Error("Class should not be a comment")
	}
	if Def.IsComment() {
		t.Error("Def should not be a comment")
	}
	if Do.IsComment() {
		t.Error("Do should not be a comment")
	}
	if Elif.IsComment() {
		t.Error("Elif should not be a comment")
	}
	if Else.IsComment() {
		t.Error("Else should not be a comment")
	}
	if For.IsComment() {
		t.Error("For should not be a comment")
	}
	if From.IsComment() {
		t.Error("From should not be a comment")
	}
	if If.IsComment() {
		t.Error("If should not be a comment")
	}
	if In.IsComment() {
		t.Error("In should not be a comment")
	}
	if Loop.IsComment() {
		t.Error("Loop should not be a comment")
	}
	if Panic.IsComment() {
		t.Error("Panic should not be a comment")
	}
	if Pass.IsComment() {
		t.Error("Pass should not be a comment")
	}
	if Recover.IsComment() {
		t.Error("Recover should not be a comment")
	}
	if Redo.IsComment() {
		t.Error("Redo should not be a comment")
	}
	if Retry.IsComment() {
		t.Error("Retry should not be a comment")
	}
	if Return.IsComment() {
		t.Error("Return should not be a comment")
	}
	if Self.IsComment() {
		t.Error("Self should not be a comment")
	}
	if Then.IsComment() {
		t.Error("Then should not be a comment")
	}
	if Try.IsComment() {
		t.Error("Try should not be a comment")
	}
	if Unless.IsComment() {
		t.Error("Unless should not be a comment")
	}
	if Until.IsComment() {
		t.Error("Until should not be a comment")
	}
	if Where.IsComment() {
		t.Error("Where should not be a comment")
	}
	if While.IsComment() {
		t.Error("While should not be a comment")
	}
	if Bytes.IsComment() {
		t.Error("Bytes should not be a comment")
	}
	if Complex.IsComment() {
		t.Error("Complex should not be a comment")
	}
	if Float.IsComment() {
		t.Error("Float should not be a comment")
	}
	if Int.IsComment() {
		t.Error("Int should not be a comment")
	}
	if Nil.IsComment() {
		t.Error("Nil should not be a comment")
	}
	if StringDouble.IsComment() {
		t.Error("StringDouble should not be a comment")
	}
	if StringDoubleBlock.IsComment() {
		t.Error("StringDoubleBlock should not be a comment")
	}
	if StringSingle.IsComment() {
		t.Error("StringSingle should not be a comment")
	}
	if StringSingleBlock.IsComment() {
		t.Error("StringSingleBlock should not be a comment")
	}
	if Symbol.IsComment() {
		t.Error("Symbol should not be a comment")
	}
	if Add.IsComment() {
		t.Error("Add should not be a comment")
	}
	if Assign.IsComment() {
		t.Error("Assign should not be a comment")
	}
	if BitAnd.IsComment() {
		t.Error("BitAnd should not be a comment")
	}
	if BitOr.IsComment() {
		t.Error("BitOr should not be a comment")
	}
	if BitXor.IsComment() {
		t.Error("BitXor should not be a comment")
	}
	if Compare.IsComment() {
		t.Error("Compare should not be a comment")
	}
	if Concat.IsComment() {
		t.Error("Concat should not be a comment")
	}
	if Div.IsComment() {
		t.Error("Div should not be a comment")
	}
	if DivFloor.IsComment() {
		t.Error("DivFloor should not be a comment")
	}
	if Eq.IsComment() {
		t.Error("Eq should not be a comment")
	}
	if Exp.IsComment() {
		t.Error("Exp should not be a comment")
	}
	if Ge.IsComment() {
		t.Error("Ge should not be a comment")
	}
	if Gt.IsComment() {
		t.Error("Gt should not be a comment")
	}
	if Le.IsComment() {
		t.Error("Le should not be a comment")
	}
	if LogicAnd.IsComment() {
		t.Error("LogicAnd should not be a comment")
	}
	if LogicOr.IsComment() {
		t.Error("LogicOr should not be a comment")
	}
	if Lt.IsComment() {
		t.Error("Lt should not be a comment")
	}
	if Mod.IsComment() {
		t.Error("Mod should not be a comment")
	}
	if Mul.IsComment() {
		t.Error("Mul should not be a comment")
	}
	if Ne.IsComment() {
		t.Error("Ne should not be a comment")
	}
	if RangeExc.IsComment() {
		t.Error("RangeExc should not be a comment")
	}
	if RangeInc.IsComment() {
		t.Error("RangeInc should not be a comment")
	}
	if ShiftLeft.IsComment() {
		t.Error("ShiftLeft should not be a comment")
	}
	if ShiftRight.IsComment() {
		t.Error("ShiftRight should not be a comment")
	}
	if Sub.IsComment() {
		t.Error("Sub should not be a comment")
	}
	if BitNot.IsComment() {
		t.Error("BitNot should not be a comment")
	}
	if Block.IsComment() {
		t.Error("Block should not be a comment")
	}
	if BlockBang.IsComment() {
		t.Error("BlockBang should not be a comment")
	}
	if LogicNot.IsComment() {
		t.Error("LogicNot should not be a comment")
	}
	if Neg.IsComment() {
		t.Error("Neg should not be a comment")
	}
	if Pos.IsComment() {
		t.Error("Pos should not be a comment")
	}
	if Splat.IsComment() {
		t.Error("Splat should not be a comment")
	}
	if SplatMap.IsComment() {
		t.Error("SplatMap should not be a comment")
	}
	if Arrow.IsComment() {
		t.Error("Arrow should not be a comment")
	}
	if Bang.IsComment() {
		t.Error("Bang should not be a comment")
	}
	if Colon.IsComment() {
		t.Error("Colon should not be a comment")
	}
	if Comma.IsComment() {
		t.Error("Comma should not be a comment")
	}
	if Cont.IsComment() {
		t.Error("Cont should not be a comment")
	}
	if Dot.IsComment() {
		t.Error("Dot should not be a comment")
	}
	if LBrace.IsComment() {
		t.Error("LBrace should not be a comment")
	}
	if LBrack.IsComment() {
		t.Error("LBrack should not be a comment")
	}
	if LParen.IsComment() {
		t.Error("LParen should not be a comment")
	}
	if Pipe.IsComment() {
		t.Error("Pipe should not be a comment")
	}
	if PipeTo.IsComment() {
		t.Error("PipeTo should not be a comment")
	}
	if RBrace.IsComment() {
		t.Error("RBrace should not be a comment")
	}
	if RBrack.IsComment() {
		t.Error("RBrack should not be a comment")
	}
	if Rparen.IsComment() {
		t.Error("Rparen should not be a comment")
	}
	if SplatAll.IsComment() {
		t.Error("SplatAll should not be a comment")
	}
}

func TestIsSpace(t *testing.T) {
	if Illegal.IsSpace() {
		t.Error("Illegal should not be a space")
	}
	if Comment.IsSpace() {
		t.Error("Comment should not be a space")
	}
	if !Space.IsSpace() {
		t.Error("Space should not be a space")
	}
	if Newline.IsSpace() {
		t.Error("Newline should not be a space")
	}
	if Const.IsSpace() {
		t.Error("Const should not be a space")
	}
	if ConstHidden.IsSpace() {
		t.Error("ConstHidden should not be a space")
	}
	if Field.IsSpace() {
		t.Error("Field should not be a space")
	}
	if FieldStatic.IsSpace() {
		t.Error("FieldStaticshould not be a space")
	}
	if Sink.IsSpace() {
		t.Error("Sink should not be a space")
	}
	if Var.IsSpace() {
		t.Error("Var should not be a space")
	}
	if VarHidden.IsSpace() {
		t.Error("VarHidden should not be a space")
	}
	if Always.IsSpace() {
		t.Error("Always should not be a space")
	}
	if As.IsSpace() {
		t.Error("As should not be a space")
	}
	if Break.IsSpace() {
		t.Error("Break should not be a space")
	}
	if Class.IsSpace() {
		t.Error("Class should not be a space")
	}
	if Def.IsSpace() {
		t.Error("Def should not be a space")
	}
	if Do.IsSpace() {
		t.Error("Do should not be a space")
	}
	if Elif.IsSpace() {
		t.Error("Elif should not be a space")
	}
	if Else.IsSpace() {
		t.Error("Else should not be a space")
	}
	if For.IsSpace() {
		t.Error("For should not be a space")
	}
	if From.IsSpace() {
		t.Error("From should not be a space")
	}
	if If.IsSpace() {
		t.Error("If should not be a space")
	}
	if In.IsSpace() {
		t.Error("In should not be a space")
	}
	if Loop.IsSpace() {
		t.Error("Loop should not be a space")
	}
	if Panic.IsSpace() {
		t.Error("Panic should not be a space")
	}
	if Pass.IsSpace() {
		t.Error("Pass should not be a space")
	}
	if Recover.IsSpace() {
		t.Error("Recover should not be a space")
	}
	if Redo.IsSpace() {
		t.Error("Redo should not be a space")
	}
	if Retry.IsSpace() {
		t.Error("Retry should not be a space")
	}
	if Return.IsSpace() {
		t.Error("Return should not be a space")
	}
	if Self.IsSpace() {
		t.Error("Self should not be a space")
	}
	if Then.IsSpace() {
		t.Error("Then should not be a space")
	}
	if Try.IsSpace() {
		t.Error("Try should not be a space")
	}
	if Unless.IsSpace() {
		t.Error("Unless should not be a space")
	}
	if Until.IsSpace() {
		t.Error("Until should not be a space")
	}
	if Where.IsSpace() {
		t.Error("Where should not be a space")
	}
	if While.IsSpace() {
		t.Error("While should not be a space")
	}
	if Bytes.IsSpace() {
		t.Error("Bytes should not be a space")
	}
	if Complex.IsSpace() {
		t.Error("Complex should not be a space")
	}
	if Float.IsSpace() {
		t.Error("Float should not be a space")
	}
	if Int.IsSpace() {
		t.Error("Int should not be a space")
	}
	if Nil.IsSpace() {
		t.Error("Nil should not be a space")
	}
	if StringDouble.IsSpace() {
		t.Error("StringDouble should not be a space")
	}
	if StringDoubleBlock.IsSpace() {
		t.Error("StringDoubleBlock should not be a space")
	}
	if StringSingle.IsSpace() {
		t.Error("StringSingle should not be a space")
	}
	if StringSingleBlock.IsSpace() {
		t.Error("StringSingleBlock should not be a space")
	}
	if Symbol.IsSpace() {
		t.Error("Symbol should not be a space")
	}
	if Add.IsSpace() {
		t.Error("Add should not be a space")
	}
	if Assign.IsSpace() {
		t.Error("Assign should not be a space")
	}
	if BitAnd.IsSpace() {
		t.Error("BitAnd should not be a space")
	}
	if BitOr.IsSpace() {
		t.Error("BitOr should not be a space")
	}
	if BitXor.IsSpace() {
		t.Error("BitXor should not be a space")
	}
	if Compare.IsSpace() {
		t.Error("Compare should not be a space")
	}
	if Concat.IsSpace() {
		t.Error("Concat should not be a space")
	}
	if Div.IsSpace() {
		t.Error("Div should not be a space")
	}
	if DivFloor.IsSpace() {
		t.Error("DivFloor should not be a space")
	}
	if Eq.IsSpace() {
		t.Error("Eq should not be a space")
	}
	if Exp.IsSpace() {
		t.Error("Exp should not be a space")
	}
	if Ge.IsSpace() {
		t.Error("Ge should not be a space")
	}
	if Gt.IsSpace() {
		t.Error("Gt should not be a space")
	}
	if Le.IsSpace() {
		t.Error("Le should not be a space")
	}
	if LogicAnd.IsSpace() {
		t.Error("LogicAnd should not be a space")
	}
	if LogicOr.IsSpace() {
		t.Error("LogicOr should not be a space")
	}
	if Lt.IsSpace() {
		t.Error("Lt should not be a space")
	}
	if Mod.IsSpace() {
		t.Error("Mod should not be a space")
	}
	if Mul.IsSpace() {
		t.Error("Mul should not be a space")
	}
	if Ne.IsSpace() {
		t.Error("Ne should not be a space")
	}
	if RangeExc.IsSpace() {
		t.Error("RangeExc should not be a space")
	}
	if RangeInc.IsSpace() {
		t.Error("RangeInc should not be a space")
	}
	if ShiftLeft.IsSpace() {
		t.Error("ShiftLeft should not be a space")
	}
	if ShiftRight.IsSpace() {
		t.Error("ShiftRight should not be a space")
	}
	if Sub.IsSpace() {
		t.Error("Sub should not be a space")
	}
	if BitNot.IsSpace() {
		t.Error("BitNot should not be a space")
	}
	if Block.IsSpace() {
		t.Error("Block should not be a space")
	}
	if BlockBang.IsSpace() {
		t.Error("BlockBang should not be a space")
	}
	if LogicNot.IsSpace() {
		t.Error("LogicNot should not be a space")
	}
	if Neg.IsSpace() {
		t.Error("Neg should not be a space")
	}
	if Pos.IsSpace() {
		t.Error("Pos should not be a space")
	}
	if Splat.IsSpace() {
		t.Error("Splat should not be a space")
	}
	if SplatMap.IsSpace() {
		t.Error("SplatMap should not be a space")
	}
	if Arrow.IsSpace() {
		t.Error("Arrow should not be a space")
	}
	if Bang.IsSpace() {
		t.Error("Bang should not be a space")
	}
	if Colon.IsSpace() {
		t.Error("Colon should not be a space")
	}
	if Comma.IsSpace() {
		t.Error("Comma should not be a space")
	}
	if Cont.IsSpace() {
		t.Error("Cont should not be a space")
	}
	if Dot.IsSpace() {
		t.Error("Dot should not be a space")
	}
	if LBrace.IsSpace() {
		t.Error("LBrace should not be a space")
	}
	if LBrack.IsSpace() {
		t.Error("LBrack should not be a space")
	}
	if LParen.IsSpace() {
		t.Error("LParen should not be a space")
	}
	if Pipe.IsSpace() {
		t.Error("Pipe should not be a space")
	}
	if PipeTo.IsSpace() {
		t.Error("PipeTo should not be a space")
	}
	if RBrace.IsSpace() {
		t.Error("RBrace should not be a space")
	}
	if RBrack.IsSpace() {
		t.Error("RBrack should not be a space")
	}
	if Rparen.IsSpace() {
		t.Error("Rparen should not be a space")
	}
	if SplatAll.IsSpace() {
		t.Error("SplatAll should not be a space")
	}
}

func TestIsNewline(t *testing.T) {
	if Illegal.IsNewline() {
		t.Error("Illegal should not be a newline")
	}
	if Comment.IsNewline() {
		t.Error("Comment should not be a newline")
	}
	if Space.IsNewline() {
		t.Error("Space should not be a newline")
	}
	if !Newline.IsNewline() {
		t.Error("Newline should be a newline")
	}
	if Const.IsNewline() {
		t.Error("Const should not be a newline")
	}
	if ConstHidden.IsNewline() {
		t.Error("ConstHidden should not be a newline")
	}
	if Field.IsNewline() {
		t.Error("Field should not be a newline")
	}
	if FieldStatic.IsNewline() {
		t.Error("FieldStaticshould not be a newline")
	}
	if Sink.IsNewline() {
		t.Error("Sink should not be a newline")
	}
	if Var.IsNewline() {
		t.Error("Var should not be a newline")
	}
	if VarHidden.IsNewline() {
		t.Error("VarHidden should not be a newline")
	}
	if Always.IsNewline() {
		t.Error("Always should not be a newline")
	}
	if As.IsNewline() {
		t.Error("As should not be a newline")
	}
	if Break.IsNewline() {
		t.Error("Break should not be a newline")
	}
	if Class.IsNewline() {
		t.Error("Class should not be a newline")
	}
	if Def.IsNewline() {
		t.Error("Def should not be a newline")
	}
	if Do.IsNewline() {
		t.Error("Do should not be a newline")
	}
	if Elif.IsNewline() {
		t.Error("Elif should not be a newline")
	}
	if Else.IsNewline() {
		t.Error("Else should not be a newline")
	}
	if For.IsNewline() {
		t.Error("For should not be a newline")
	}
	if From.IsNewline() {
		t.Error("From should not be a newline")
	}
	if If.IsNewline() {
		t.Error("If should not be a newline")
	}
	if In.IsNewline() {
		t.Error("In should not be a newline")
	}
	if Loop.IsNewline() {
		t.Error("Loop should not be a newline")
	}
	if Panic.IsNewline() {
		t.Error("Panic should not be a newline")
	}
	if Pass.IsNewline() {
		t.Error("Pass should not be a newline")
	}
	if Recover.IsNewline() {
		t.Error("Recover should not be a newline")
	}
	if Redo.IsNewline() {
		t.Error("Redo should not be a newline")
	}
	if Retry.IsNewline() {
		t.Error("Retry should not be a newline")
	}
	if Return.IsNewline() {
		t.Error("Return should not be a newline")
	}
	if Self.IsNewline() {
		t.Error("Self should not be a newline")
	}
	if Then.IsNewline() {
		t.Error("Then should not be a newline")
	}
	if Try.IsNewline() {
		t.Error("Try should not be a newline")
	}
	if Unless.IsNewline() {
		t.Error("Unless should not be a newline")
	}
	if Until.IsNewline() {
		t.Error("Until should not be a newline")
	}
	if Where.IsNewline() {
		t.Error("Where should not be a newline")
	}
	if While.IsNewline() {
		t.Error("While should not be a newline")
	}
	if Bytes.IsNewline() {
		t.Error("Bytes should not be a newline")
	}
	if Complex.IsNewline() {
		t.Error("Complex should not be a newline")
	}
	if Float.IsNewline() {
		t.Error("Float should not be a newline")
	}
	if Int.IsNewline() {
		t.Error("Int should not be a newline")
	}
	if Nil.IsNewline() {
		t.Error("Nil should not be a newline")
	}
	if StringDouble.IsNewline() {
		t.Error("StringDouble should not be a newline")
	}
	if StringDoubleBlock.IsNewline() {
		t.Error("StringDoubleBlock should not be a newline")
	}
	if StringSingle.IsNewline() {
		t.Error("StringSingle should not be a newline")
	}
	if StringSingleBlock.IsNewline() {
		t.Error("StringSingleBlock should not be a newline")
	}
	if Symbol.IsNewline() {
		t.Error("Symbol should not be a newline")
	}
	if Add.IsNewline() {
		t.Error("Add should not be a newline")
	}
	if Assign.IsNewline() {
		t.Error("Assign should not be a newline")
	}
	if BitAnd.IsNewline() {
		t.Error("BitAnd should not be a newline")
	}
	if BitOr.IsNewline() {
		t.Error("BitOr should not be a newline")
	}
	if BitXor.IsNewline() {
		t.Error("BitXor should not be a newline")
	}
	if Compare.IsNewline() {
		t.Error("Compare should not be a newline")
	}
	if Concat.IsNewline() {
		t.Error("Concat should not be a newline")
	}
	if Div.IsNewline() {
		t.Error("Div should not be a newline")
	}
	if DivFloor.IsNewline() {
		t.Error("DivFloor should not be a newline")
	}
	if Eq.IsNewline() {
		t.Error("Eq should not be a newline")
	}
	if Exp.IsNewline() {
		t.Error("Exp should not be a newline")
	}
	if Ge.IsNewline() {
		t.Error("Ge should not be a newline")
	}
	if Gt.IsNewline() {
		t.Error("Gt should not be a newline")
	}
	if Le.IsNewline() {
		t.Error("Le should not be a newline")
	}
	if LogicAnd.IsNewline() {
		t.Error("LogicAnd should not be a newline")
	}
	if LogicOr.IsNewline() {
		t.Error("LogicOr should not be a newline")
	}
	if Lt.IsNewline() {
		t.Error("Lt should not be a newline")
	}
	if Mod.IsNewline() {
		t.Error("Mod should not be a newline")
	}
	if Mul.IsNewline() {
		t.Error("Mul should not be a newline")
	}
	if Ne.IsNewline() {
		t.Error("Ne should not be a newline")
	}
	if RangeExc.IsNewline() {
		t.Error("RangeExc should not be a newline")
	}
	if RangeInc.IsNewline() {
		t.Error("RangeInc should not be a newline")
	}
	if ShiftLeft.IsNewline() {
		t.Error("ShiftLeft should not be a newline")
	}
	if ShiftRight.IsNewline() {
		t.Error("ShiftRight should not be a newline")
	}
	if Sub.IsNewline() {
		t.Error("Sub should not be a newline")
	}
	if BitNot.IsNewline() {
		t.Error("BitNot should not be a newline")
	}
	if Block.IsNewline() {
		t.Error("Block should not be a newline")
	}
	if BlockBang.IsNewline() {
		t.Error("BlockBang should not be a newline")
	}
	if LogicNot.IsNewline() {
		t.Error("LogicNot should not be a newline")
	}
	if Neg.IsNewline() {
		t.Error("Neg should not be a newline")
	}
	if Pos.IsNewline() {
		t.Error("Pos should not be a newline")
	}
	if Splat.IsNewline() {
		t.Error("Splat should not be a newline")
	}
	if SplatMap.IsNewline() {
		t.Error("SplatMap should not be a newline")
	}
	if Arrow.IsNewline() {
		t.Error("Arrow should not be a newline")
	}
	if Bang.IsNewline() {
		t.Error("Bang should not be a newline")
	}
	if Colon.IsNewline() {
		t.Error("Colon should not be a newline")
	}
	if Comma.IsNewline() {
		t.Error("Comma should not be a newline")
	}
	if Cont.IsNewline() {
		t.Error("Cont should not be a newline")
	}
	if Dot.IsNewline() {
		t.Error("Dot should not be a newline")
	}
	if LBrace.IsNewline() {
		t.Error("LBrace should not be a newline")
	}
	if LBrack.IsNewline() {
		t.Error("LBrack should not be a newline")
	}
	if LParen.IsNewline() {
		t.Error("LParen should not be a newline")
	}
	if Pipe.IsNewline() {
		t.Error("Pipe should not be a newline")
	}
	if PipeTo.IsNewline() {
		t.Error("PipeTo should not be a newline")
	}
	if RBrace.IsNewline() {
		t.Error("RBrace should not be a newline")
	}
	if RBrack.IsNewline() {
		t.Error("RBrack should not be a newline")
	}
	if Rparen.IsNewline() {
		t.Error("Rparen should not be a newline")
	}
	if SplatAll.IsNewline() {
		t.Error("SplatAll should not be a newline")
	}
}

func TestIsIdent(t *testing.T) {
	if Illegal.IsIdent() {
		t.Error("Illegal should not be an identifier")
	}
	if Comment.IsIdent() {
		t.Error("Comment should not be an identifier")
	}
	if Space.IsIdent() {
		t.Error("Space should not be an identifier")
	}
	if Newline.IsIdent() {
		t.Error("Newline should not be an identifier")
	}
	if !Const.IsIdent() {
		t.Error("Const should be an identifier")
	}
	if !ConstHidden.IsIdent() {
		t.Error("ConstHidden should be an identifier")
	}
	if !Field.IsIdent() {
		t.Error("Field should be an identifier")
	}
	if !FieldStatic.IsIdent() {
		t.Error("FieldStaticshould be an identifier")
	}
	if !Sink.IsIdent() {
		t.Error("Sink should be an identifier")
	}
	if !Var.IsIdent() {
		t.Error("Var should be an identifier")
	}
	if !VarHidden.IsIdent() {
		t.Error("VarHidden should be an identifier")
	}
	if Always.IsIdent() {
		t.Error("Always should not be an identifier")
	}
	if As.IsIdent() {
		t.Error("As should not be an identifier")
	}
	if Break.IsIdent() {
		t.Error("Break should not be an identifier")
	}
	if Class.IsIdent() {
		t.Error("Class should not be an identifier")
	}
	if Def.IsIdent() {
		t.Error("Def should not be an identifier")
	}
	if Do.IsIdent() {
		t.Error("Do should not be an identifier")
	}
	if Elif.IsIdent() {
		t.Error("Elif should not be an identifier")
	}
	if Else.IsIdent() {
		t.Error("Else should not be an identifier")
	}
	if For.IsIdent() {
		t.Error("For should not be an identifier")
	}
	if From.IsIdent() {
		t.Error("From should not be an identifier")
	}
	if If.IsIdent() {
		t.Error("If should not be an identifier")
	}
	if In.IsIdent() {
		t.Error("In should not be an identifier")
	}
	if Loop.IsIdent() {
		t.Error("Loop should not be an identifier")
	}
	if Panic.IsIdent() {
		t.Error("Panic should not be an identifier")
	}
	if Pass.IsIdent() {
		t.Error("Pass should not be an identifier")
	}
	if Recover.IsIdent() {
		t.Error("Recover should not be an identifier")
	}
	if Redo.IsIdent() {
		t.Error("Redo should not be an identifier")
	}
	if Retry.IsIdent() {
		t.Error("Retry should not be an identifier")
	}
	if Return.IsIdent() {
		t.Error("Return should not be an identifier")
	}
	if Self.IsIdent() {
		t.Error("Self should not be an identifier")
	}
	if Then.IsIdent() {
		t.Error("Then should not be an identifier")
	}
	if Try.IsIdent() {
		t.Error("Try should not be an identifier")
	}
	if Unless.IsIdent() {
		t.Error("Unless should not be an identifier")
	}
	if Until.IsIdent() {
		t.Error("Until should not be an identifier")
	}
	if Where.IsIdent() {
		t.Error("Where should not be an identifier")
	}
	if While.IsIdent() {
		t.Error("While should not be an identifier")
	}
	if Bytes.IsIdent() {
		t.Error("Bytes should not be an identifier")
	}
	if Complex.IsIdent() {
		t.Error("Complex should not be an identifier")
	}
	if Float.IsIdent() {
		t.Error("Float should not be an identifier")
	}
	if Int.IsIdent() {
		t.Error("Int should not be an identifier")
	}
	if Nil.IsIdent() {
		t.Error("Nil should not be an identifier")
	}
	if StringDouble.IsIdent() {
		t.Error("StringDouble should not be an identifier")
	}
	if StringDoubleBlock.IsIdent() {
		t.Error("StringDoubleBlock should not be an identifier")
	}
	if StringSingle.IsIdent() {
		t.Error("StringSingle should not be an identifier")
	}
	if StringSingleBlock.IsIdent() {
		t.Error("StringSingleBlock should not be an identifier")
	}
	if Symbol.IsIdent() {
		t.Error("Symbol should not be an identifier")
	}
	if Add.IsIdent() {
		t.Error("Add should not be an identifier")
	}
	if Assign.IsIdent() {
		t.Error("Assign should not be an identifier")
	}
	if BitAnd.IsIdent() {
		t.Error("BitAnd should not be an identifier")
	}
	if BitOr.IsIdent() {
		t.Error("BitOr should not be an identifier")
	}
	if BitXor.IsIdent() {
		t.Error("BitXor should not be an identifier")
	}
	if Compare.IsIdent() {
		t.Error("Compare should not be an identifier")
	}
	if Concat.IsIdent() {
		t.Error("Concat should not be an identifier")
	}
	if Div.IsIdent() {
		t.Error("Div should not be an identifier")
	}
	if DivFloor.IsIdent() {
		t.Error("DivFloor should not be an identifier")
	}
	if Eq.IsIdent() {
		t.Error("Eq should not be an identifier")
	}
	if Exp.IsIdent() {
		t.Error("Exp should not be an identifier")
	}
	if Ge.IsIdent() {
		t.Error("Ge should not be an identifier")
	}
	if Gt.IsIdent() {
		t.Error("Gt should not be an identifier")
	}
	if Le.IsIdent() {
		t.Error("Le should not be an identifier")
	}
	if LogicAnd.IsIdent() {
		t.Error("LogicAnd should not be an identifier")
	}
	if LogicOr.IsIdent() {
		t.Error("LogicOr should not be an identifier")
	}
	if Lt.IsIdent() {
		t.Error("Lt should not be an identifier")
	}
	if Mod.IsIdent() {
		t.Error("Mod should not be an identifier")
	}
	if Mul.IsIdent() {
		t.Error("Mul should not be an identifier")
	}
	if Ne.IsIdent() {
		t.Error("Ne should not be an identifier")
	}
	if RangeExc.IsIdent() {
		t.Error("RangeExc should not be an identifier")
	}
	if RangeInc.IsIdent() {
		t.Error("RangeInc should not be an identifier")
	}
	if ShiftLeft.IsIdent() {
		t.Error("ShiftLeft should not be an identifier")
	}
	if ShiftRight.IsIdent() {
		t.Error("ShiftRight should not be an identifier")
	}
	if Sub.IsIdent() {
		t.Error("Sub should not be an identifier")
	}
	if BitNot.IsIdent() {
		t.Error("BitNot should not be an identifier")
	}
	if Block.IsIdent() {
		t.Error("Block should not be an identifier")
	}
	if BlockBang.IsIdent() {
		t.Error("BlockBang should not be an identifier")
	}
	if LogicNot.IsIdent() {
		t.Error("LogicNot should not be an identifier")
	}
	if Neg.IsIdent() {
		t.Error("Neg should not be an identifier")
	}
	if Pos.IsIdent() {
		t.Error("Pos should not be an identifier")
	}
	if Splat.IsIdent() {
		t.Error("Splat should not be an identifier")
	}
	if SplatMap.IsIdent() {
		t.Error("SplatMap should not be an identifier")
	}
	if Arrow.IsIdent() {
		t.Error("Arrow should not be an identifier")
	}
	if Bang.IsIdent() {
		t.Error("Bang should not be an identifier")
	}
	if Colon.IsIdent() {
		t.Error("Colon should not be an identifier")
	}
	if Comma.IsIdent() {
		t.Error("Comma should not be an identifier")
	}
	if Cont.IsIdent() {
		t.Error("Cont should not be an identifier")
	}
	if Dot.IsIdent() {
		t.Error("Dot should not be an identifier")
	}
	if LBrace.IsIdent() {
		t.Error("LBrace should not be an identifier")
	}
	if LBrack.IsIdent() {
		t.Error("LBrack should not be an identifier")
	}
	if LParen.IsIdent() {
		t.Error("LParen should not be an identifier")
	}
	if Pipe.IsIdent() {
		t.Error("Pipe should not be an identifier")
	}
	if PipeTo.IsIdent() {
		t.Error("PipeTo should not be an identifier")
	}
	if RBrace.IsIdent() {
		t.Error("RBrace should not be an identifier")
	}
	if RBrack.IsIdent() {
		t.Error("RBrack should not be an identifier")
	}
	if Rparen.IsIdent() {
		t.Error("Rparen should not be an identifier")
	}
	if SplatAll.IsIdent() {
		t.Error("SplatAll should not be an identifier")
	}
}

func TestIsKeyword(t *testing.T) {
	if Illegal.IsKeyword() {
		t.Error("Illegal should not be a keyword")
	}
	if Comment.IsKeyword() {
		t.Error("Comment should not be a keyword")
	}
	if Space.IsKeyword() {
		t.Error("Space should not be a keyword")
	}
	if Newline.IsKeyword() {
		t.Error("Newline should not be a keyword")
	}
	if Const.IsKeyword() {
		t.Error("Const should not be a keyword")
	}
	if ConstHidden.IsKeyword() {
		t.Error("ConstHidden should not be a keyword")
	}
	if Field.IsKeyword() {
		t.Error("Field should not be a keyword")
	}
	if FieldStatic.IsKeyword() {
		t.Error("FieldStaticshould not be a keyword")
	}
	if Sink.IsKeyword() {
		t.Error("Sink should not be a keyword")
	}
	if Var.IsKeyword() {
		t.Error("Var should not be a keyword")
	}
	if VarHidden.IsKeyword() {
		t.Error("VarHidden should not be a keyword")
	}
	if !Always.IsKeyword() {
		t.Error("Always should be a keyword")
	}
	if !As.IsKeyword() {
		t.Error("As should be a keyword")
	}
	if !Break.IsKeyword() {
		t.Error("Break should be a keyword")
	}
	if !Class.IsKeyword() {
		t.Error("Class should be a keyword")
	}
	if !Def.IsKeyword() {
		t.Error("Def should be a keyword")
	}
	if !Do.IsKeyword() {
		t.Error("Do should be a keyword")
	}
	if !Elif.IsKeyword() {
		t.Error("Elif should be a keyword")
	}
	if !Else.IsKeyword() {
		t.Error("Else should be a keyword")
	}
	if !For.IsKeyword() {
		t.Error("For should be a keyword")
	}
	if !From.IsKeyword() {
		t.Error("From should be a keyword")
	}
	if !If.IsKeyword() {
		t.Error("If should be a keyword")
	}
	if !In.IsKeyword() {
		t.Error("In should be a keyword")
	}
	if !Loop.IsKeyword() {
		t.Error("Loop should be a keyword")
	}
	if !Panic.IsKeyword() {
		t.Error("Panic should be a keyword")
	}
	if !Pass.IsKeyword() {
		t.Error("Pass should be a keyword")
	}
	if !Recover.IsKeyword() {
		t.Error("Recover should be a keyword")
	}
	if !Redo.IsKeyword() {
		t.Error("Redo should be a keyword")
	}
	if !Retry.IsKeyword() {
		t.Error("Retry should be a keyword")
	}
	if !Return.IsKeyword() {
		t.Error("Return should be a keyword")
	}
	if !Self.IsKeyword() {
		t.Error("Self should be a keyword")
	}
	if !Then.IsKeyword() {
		t.Error("Then should be a keyword")
	}
	if !Try.IsKeyword() {
		t.Error("Try should be a keyword")
	}
	if !Unless.IsKeyword() {
		t.Error("Unless should be a keyword")
	}
	if !Until.IsKeyword() {
		t.Error("Until should be a keyword")
	}
	if !Where.IsKeyword() {
		t.Error("Where should be a keyword")
	}
	if !While.IsKeyword() {
		t.Error("While should be a keyword")
	}
	if Bytes.IsKeyword() {
		t.Error("Bytes should not be a keyword")
	}
	if Complex.IsKeyword() {
		t.Error("Complex should not be a keyword")
	}
	if Float.IsKeyword() {
		t.Error("Float should not be a keyword")
	}
	if Int.IsKeyword() {
		t.Error("Int should not be a keyword")
	}
	if Nil.IsKeyword() {
		t.Error("Nil should not be a keyword")
	}
	if StringDouble.IsKeyword() {
		t.Error("StringDouble should not be a keyword")
	}
	if StringDoubleBlock.IsKeyword() {
		t.Error("StringDoubleBlock should not be a keyword")
	}
	if StringSingle.IsKeyword() {
		t.Error("StringSingle should not be a keyword")
	}
	if StringSingleBlock.IsKeyword() {
		t.Error("StringSingleBlock should not be a keyword")
	}
	if Symbol.IsKeyword() {
		t.Error("Symbol should not be a keyword")
	}
	if Add.IsKeyword() {
		t.Error("Add should not be a keyword")
	}
	if Assign.IsKeyword() {
		t.Error("Assign should not be a keyword")
	}
	if BitAnd.IsKeyword() {
		t.Error("BitAnd should not be a keyword")
	}
	if BitOr.IsKeyword() {
		t.Error("BitOr should not be a keyword")
	}
	if BitXor.IsKeyword() {
		t.Error("BitXor should not be a keyword")
	}
	if Compare.IsKeyword() {
		t.Error("Compare should not be a keyword")
	}
	if Concat.IsKeyword() {
		t.Error("Concat should not be a keyword")
	}
	if Div.IsKeyword() {
		t.Error("Div should not be a keyword")
	}
	if DivFloor.IsKeyword() {
		t.Error("DivFloor should not be a keyword")
	}
	if Eq.IsKeyword() {
		t.Error("Eq should not be a keyword")
	}
	if Exp.IsKeyword() {
		t.Error("Exp should not be a keyword")
	}
	if Ge.IsKeyword() {
		t.Error("Ge should not be a keyword")
	}
	if Gt.IsKeyword() {
		t.Error("Gt should not be a keyword")
	}
	if Le.IsKeyword() {
		t.Error("Le should not be a keyword")
	}
	if LogicAnd.IsKeyword() {
		t.Error("LogicAnd should not be a keyword")
	}
	if LogicOr.IsKeyword() {
		t.Error("LogicOr should not be a keyword")
	}
	if Lt.IsKeyword() {
		t.Error("Lt should not be a keyword")
	}
	if Mod.IsKeyword() {
		t.Error("Mod should not be a keyword")
	}
	if Mul.IsKeyword() {
		t.Error("Mul should not be a keyword")
	}
	if Ne.IsKeyword() {
		t.Error("Ne should not be a keyword")
	}
	if RangeExc.IsKeyword() {
		t.Error("RangeExc should not be a keyword")
	}
	if RangeInc.IsKeyword() {
		t.Error("RangeInc should not be a keyword")
	}
	if ShiftLeft.IsKeyword() {
		t.Error("ShiftLeft should not be a keyword")
	}
	if ShiftRight.IsKeyword() {
		t.Error("ShiftRight should not be a keyword")
	}
	if Sub.IsKeyword() {
		t.Error("Sub should not be a keyword")
	}
	if BitNot.IsKeyword() {
		t.Error("BitNot should not be a keyword")
	}
	if Block.IsKeyword() {
		t.Error("Block should not be a keyword")
	}
	if BlockBang.IsKeyword() {
		t.Error("BlockBang should not be a keyword")
	}
	if LogicNot.IsKeyword() {
		t.Error("LogicNot should not be a keyword")
	}
	if Neg.IsKeyword() {
		t.Error("Neg should not be a keyword")
	}
	if Pos.IsKeyword() {
		t.Error("Pos should not be a keyword")
	}
	if Splat.IsKeyword() {
		t.Error("Splat should not be a keyword")
	}
	if SplatMap.IsKeyword() {
		t.Error("SplatMap should not be a keyword")
	}
	if Arrow.IsKeyword() {
		t.Error("Arrow should not be a keyword")
	}
	if Bang.IsKeyword() {
		t.Error("Bang should not be a keyword")
	}
	if Colon.IsKeyword() {
		t.Error("Colon should not be a keyword")
	}
	if Comma.IsKeyword() {
		t.Error("Comma should not be a keyword")
	}
	if Cont.IsKeyword() {
		t.Error("Cont should not be a keyword")
	}
	if Dot.IsKeyword() {
		t.Error("Dot should not be a keyword")
	}
	if LBrace.IsKeyword() {
		t.Error("LBrace should not be a keyword")
	}
	if LBrack.IsKeyword() {
		t.Error("LBrack should not be a keyword")
	}
	if LParen.IsKeyword() {
		t.Error("LParen should not be a keyword")
	}
	if Pipe.IsKeyword() {
		t.Error("Pipe should not be a keyword")
	}
	if PipeTo.IsKeyword() {
		t.Error("PipeTo should not be a keyword")
	}
	if RBrace.IsKeyword() {
		t.Error("RBrace should not be a keyword")
	}
	if RBrack.IsKeyword() {
		t.Error("RBrack should not be a keyword")
	}
	if Rparen.IsKeyword() {
		t.Error("Rparen should not be a keyword")
	}
	if SplatAll.IsKeyword() {
		t.Error("SplatAll should not be a keyword")
	}
}

func TestIsLiteral(t *testing.T) {
	if Illegal.IsLiteral() {
		t.Error("Illegal should not be a literal")
	}
	if Comment.IsLiteral() {
		t.Error("Comment should not be a literal")
	}
	if Space.IsLiteral() {
		t.Error("Space should not be a literal")
	}
	if Newline.IsLiteral() {
		t.Error("Newline should not be a literal")
	}
	if Const.IsLiteral() {
		t.Error("Const should not be a literal")
	}
	if ConstHidden.IsLiteral() {
		t.Error("ConstHidden should not be a literal")
	}
	if Field.IsLiteral() {
		t.Error("Field should not be a literal")
	}
	if FieldStatic.IsLiteral() {
		t.Error("FieldStaticshould not be a literal")
	}
	if Sink.IsLiteral() {
		t.Error("Sink should not be a literal")
	}
	if Var.IsLiteral() {
		t.Error("Var should not be a literal")
	}
	if VarHidden.IsLiteral() {
		t.Error("VarHidden should not be a literal")
	}
	if Always.IsLiteral() {
		t.Error("Always should not be a literal")
	}
	if As.IsLiteral() {
		t.Error("As should not be a literal")
	}
	if Break.IsLiteral() {
		t.Error("Break should not be a literal")
	}
	if Class.IsLiteral() {
		t.Error("Class should not be a literal")
	}
	if Def.IsLiteral() {
		t.Error("Def should not be a literal")
	}
	if Do.IsLiteral() {
		t.Error("Do should not be a literal")
	}
	if Elif.IsLiteral() {
		t.Error("Elif should not be a literal")
	}
	if Else.IsLiteral() {
		t.Error("Else should not be a literal")
	}
	if For.IsLiteral() {
		t.Error("For should not be a literal")
	}
	if From.IsLiteral() {
		t.Error("From should not be a literal")
	}
	if If.IsLiteral() {
		t.Error("If should not be a literal")
	}
	if In.IsLiteral() {
		t.Error("In should not be a literal")
	}
	if Loop.IsLiteral() {
		t.Error("Loop should not be a literal")
	}
	if Panic.IsLiteral() {
		t.Error("Panic should not be a literal")
	}
	if Pass.IsLiteral() {
		t.Error("Pass should not be a literal")
	}
	if Recover.IsLiteral() {
		t.Error("Recover should not be a literal")
	}
	if Redo.IsLiteral() {
		t.Error("Redo should not be a literal")
	}
	if Retry.IsLiteral() {
		t.Error("Retry should not be a literal")
	}
	if Return.IsLiteral() {
		t.Error("Return should not be a literal")
	}
	if Self.IsLiteral() {
		t.Error("Self should not be a literal")
	}
	if Then.IsLiteral() {
		t.Error("Then should not be a literal")
	}
	if Try.IsLiteral() {
		t.Error("Try should not be a literal")
	}
	if Unless.IsLiteral() {
		t.Error("Unless should not be a literal")
	}
	if Until.IsLiteral() {
		t.Error("Until should not be a literal")
	}
	if Where.IsLiteral() {
		t.Error("Where should not be a literal")
	}
	if While.IsLiteral() {
		t.Error("While should not be a literal")
	}
	if !Bytes.IsLiteral() {
		t.Error("Bytes should be a literal")
	}
	if !Complex.IsLiteral() {
		t.Error("Complex should be a literal")
	}
	if !Float.IsLiteral() {
		t.Error("Float should be a literal")
	}
	if !Int.IsLiteral() {
		t.Error("Int should be a literal")
	}
	if !Nil.IsLiteral() {
		t.Error("Nil should be a literal")
	}
	if !StringDouble.IsLiteral() {
		t.Error("StringDouble should be a literal")
	}
	if !StringDoubleBlock.IsLiteral() {
		t.Error("StringDoubleBlock should be a literal")
	}
	if !StringSingle.IsLiteral() {
		t.Error("StringSingle should be a literal")
	}
	if !StringSingleBlock.IsLiteral() {
		t.Error("StringSingleBlock should be a literal")
	}
	if !Symbol.IsLiteral() {
		t.Error("Symbol should be a literal")
	}
	if Add.IsLiteral() {
		t.Error("Add should not be a literal")
	}
	if Assign.IsLiteral() {
		t.Error("Assign should not be a literal")
	}
	if BitAnd.IsLiteral() {
		t.Error("BitAnd should not be a literal")
	}
	if BitOr.IsLiteral() {
		t.Error("BitOr should not be a literal")
	}
	if BitXor.IsLiteral() {
		t.Error("BitXor should not be a literal")
	}
	if Compare.IsLiteral() {
		t.Error("Compare should not be a literal")
	}
	if Concat.IsLiteral() {
		t.Error("Concat should not be a literal")
	}
	if Div.IsLiteral() {
		t.Error("Div should not be a literal")
	}
	if DivFloor.IsLiteral() {
		t.Error("DivFloor should not be a literal")
	}
	if Eq.IsLiteral() {
		t.Error("Eq should not be a literal")
	}
	if Exp.IsLiteral() {
		t.Error("Exp should not be a literal")
	}
	if Ge.IsLiteral() {
		t.Error("Ge should not be a literal")
	}
	if Gt.IsLiteral() {
		t.Error("Gt should not be a literal")
	}
	if Le.IsLiteral() {
		t.Error("Le should not be a literal")
	}
	if LogicAnd.IsLiteral() {
		t.Error("LogicAnd should not be a literal")
	}
	if LogicOr.IsLiteral() {
		t.Error("LogicOr should not be a literal")
	}
	if Lt.IsLiteral() {
		t.Error("Lt should not be a literal")
	}
	if Mod.IsLiteral() {
		t.Error("Mod should not be a literal")
	}
	if Mul.IsLiteral() {
		t.Error("Mul should not be a literal")
	}
	if Ne.IsLiteral() {
		t.Error("Ne should not be a literal")
	}
	if RangeExc.IsLiteral() {
		t.Error("RangeExc should not be a literal")
	}
	if RangeInc.IsLiteral() {
		t.Error("RangeInc should not be a literal")
	}
	if ShiftLeft.IsLiteral() {
		t.Error("ShiftLeft should not be a literal")
	}
	if ShiftRight.IsLiteral() {
		t.Error("ShiftRight should not be a literal")
	}
	if Sub.IsLiteral() {
		t.Error("Sub should not be a literal")
	}
	if BitNot.IsLiteral() {
		t.Error("BitNot should not be a literal")
	}
	if Block.IsLiteral() {
		t.Error("Block should not be a literal")
	}
	if BlockBang.IsLiteral() {
		t.Error("BlockBang should not be a literal")
	}
	if LogicNot.IsLiteral() {
		t.Error("LogicNot should not be a literal")
	}
	if Neg.IsLiteral() {
		t.Error("Neg should not be a literal")
	}
	if Pos.IsLiteral() {
		t.Error("Pos should not be a literal")
	}
	if Splat.IsLiteral() {
		t.Error("Splat should not be a literal")
	}
	if SplatMap.IsLiteral() {
		t.Error("SplatMap should not be a literal")
	}
	if Arrow.IsLiteral() {
		t.Error("Arrow should not be a literal")
	}
	if Bang.IsLiteral() {
		t.Error("Bang should not be a literal")
	}
	if Colon.IsLiteral() {
		t.Error("Colon should not be a literal")
	}
	if Comma.IsLiteral() {
		t.Error("Comma should not be a literal")
	}
	if Cont.IsLiteral() {
		t.Error("Cont should not be a literal")
	}
	if Dot.IsLiteral() {
		t.Error("Dot should not be a literal")
	}
	if LBrace.IsLiteral() {
		t.Error("LBrace should not be a literal")
	}
	if LBrack.IsLiteral() {
		t.Error("LBrack should not be a literal")
	}
	if LParen.IsLiteral() {
		t.Error("LParen should not be a literal")
	}
	if Pipe.IsLiteral() {
		t.Error("Pipe should not be a literal")
	}
	if PipeTo.IsLiteral() {
		t.Error("PipeTo should not be a literal")
	}
	if RBrace.IsLiteral() {
		t.Error("RBrace should not be a literal")
	}
	if RBrack.IsLiteral() {
		t.Error("RBrack should not be a literal")
	}
	if Rparen.IsLiteral() {
		t.Error("Rparen should not be a literal")
	}
	if SplatAll.IsLiteral() {
		t.Error("SplatAll should not be a literal")
	}
}

func TestIsBinary(t *testing.T) {
	if Illegal.IsBinary() {
		t.Error("Illegal should not be a binary operator")
	}
	if Comment.IsBinary() {
		t.Error("Comment should not be a binary operator")
	}
	if Space.IsBinary() {
		t.Error("Space should not be a binary operator")
	}
	if Newline.IsBinary() {
		t.Error("Newline should not be a binary operator")
	}
	if Const.IsBinary() {
		t.Error("Const should not be a binary operator")
	}
	if ConstHidden.IsBinary() {
		t.Error("ConstHidden should not be a binary operator")
	}
	if Field.IsBinary() {
		t.Error("Field should not be a binary operator")
	}
	if FieldStatic.IsBinary() {
		t.Error("FieldStaticshould not be a binary operator")
	}
	if Sink.IsBinary() {
		t.Error("Sink should not be a binary operator")
	}
	if Var.IsBinary() {
		t.Error("Var should not be a binary operator")
	}
	if VarHidden.IsBinary() {
		t.Error("VarHidden should not be a binary operator")
	}
	if Always.IsBinary() {
		t.Error("Always should not be a binary operator")
	}
	if As.IsBinary() {
		t.Error("As should not be a binary operator")
	}
	if Break.IsBinary() {
		t.Error("Break should not be a binary operator")
	}
	if Class.IsBinary() {
		t.Error("Class should not be a binary operator")
	}
	if Def.IsBinary() {
		t.Error("Def should not be a binary operator")
	}
	if Do.IsBinary() {
		t.Error("Do should not be a binary operator")
	}
	if Elif.IsBinary() {
		t.Error("Elif should not be a binary operator")
	}
	if Else.IsBinary() {
		t.Error("Else should not be a binary operator")
	}
	if For.IsBinary() {
		t.Error("For should not be a binary operator")
	}
	if From.IsBinary() {
		t.Error("From should not be a binary operator")
	}
	if If.IsBinary() {
		t.Error("If should not be a binary operator")
	}
	if In.IsBinary() {
		t.Error("In should not be a binary operator")
	}
	if Loop.IsBinary() {
		t.Error("Loop should not be a binary operator")
	}
	if Panic.IsBinary() {
		t.Error("Panic should not be a binary operator")
	}
	if Pass.IsBinary() {
		t.Error("Pass should not be a binary operator")
	}
	if Recover.IsBinary() {
		t.Error("Recover should not be a binary operator")
	}
	if Redo.IsBinary() {
		t.Error("Redo should not be a binary operator")
	}
	if Retry.IsBinary() {
		t.Error("Retry should not be a binary operator")
	}
	if Return.IsBinary() {
		t.Error("Return should not be a binary operator")
	}
	if Self.IsBinary() {
		t.Error("Self should not be a binary operator")
	}
	if Then.IsBinary() {
		t.Error("Then should not be a binary operator")
	}
	if Try.IsBinary() {
		t.Error("Try should not be a binary operator")
	}
	if Unless.IsBinary() {
		t.Error("Unless should not be a binary operator")
	}
	if Until.IsBinary() {
		t.Error("Until should not be a binary operator")
	}
	if Where.IsBinary() {
		t.Error("Where should not be a binary operator")
	}
	if While.IsBinary() {
		t.Error("While should not be a binary operator")
	}
	if Bytes.IsBinary() {
		t.Error("Bytes should not be a binary operator")
	}
	if Complex.IsBinary() {
		t.Error("Complex should not be a binary operator")
	}
	if Float.IsBinary() {
		t.Error("Float should not be a binary operator")
	}
	if Int.IsBinary() {
		t.Error("Int should not be a binary operator")
	}
	if Nil.IsBinary() {
		t.Error("Nil should not be a binary operator")
	}
	if StringDouble.IsBinary() {
		t.Error("StringDouble should not be a binary operator")
	}
	if StringDoubleBlock.IsBinary() {
		t.Error("StringDoubleBlock should not be a binary operator")
	}
	if StringSingle.IsBinary() {
		t.Error("StringSingle should not be a binary operator")
	}
	if StringSingleBlock.IsBinary() {
		t.Error("StringSingleBlock should not be a binary operator")
	}
	if Symbol.IsBinary() {
		t.Error("Symbol should not be a binary operator")
	}
	if !Add.IsBinary() {
		t.Error("Add should be a binary operator")
	}
	if !Assign.IsBinary() {
		t.Error("Assign should be a binary operator")
	}
	if !BitAnd.IsBinary() {
		t.Error("BitAnd should be a binary operator")
	}
	if !BitOr.IsBinary() {
		t.Error("BitOr should be a binary operator")
	}
	if !BitXor.IsBinary() {
		t.Error("BitXor should be a binary operator")
	}
	if !Compare.IsBinary() {
		t.Error("Compare should be a binary operator")
	}
	if !Concat.IsBinary() {
		t.Error("Concat should be a binary operator")
	}
	if !Div.IsBinary() {
		t.Error("Div should be a binary operator")
	}
	if !DivFloor.IsBinary() {
		t.Error("DivFloor should be a binary operator")
	}
	if !Eq.IsBinary() {
		t.Error("Eq should be a binary operator")
	}
	if !Exp.IsBinary() {
		t.Error("Exp should be a binary operator")
	}
	if !Ge.IsBinary() {
		t.Error("Ge should be a binary operator")
	}
	if !Gt.IsBinary() {
		t.Error("Gt should be a binary operator")
	}
	if !Le.IsBinary() {
		t.Error("Le should be a binary operator")
	}
	if !LogicAnd.IsBinary() {
		t.Error("LogicAnd should be a binary operator")
	}
	if !LogicOr.IsBinary() {
		t.Error("LogicOr should be a binary operator")
	}
	if !Lt.IsBinary() {
		t.Error("Lt should be a binary operator")
	}
	if !Mod.IsBinary() {
		t.Error("Mod should be a binary operator")
	}
	if !Mul.IsBinary() {
		t.Error("Mul should be a binary operator")
	}
	if !Ne.IsBinary() {
		t.Error("Ne should be a binary operator")
	}
	if !RangeExc.IsBinary() {
		t.Error("RangeExc should be a binary operator")
	}
	if !RangeInc.IsBinary() {
		t.Error("RangeInc should be a binary operator")
	}
	if !ShiftLeft.IsBinary() {
		t.Error("ShiftLeft should be a binary operator")
	}
	if !ShiftRight.IsBinary() {
		t.Error("ShiftRight should be a binary operator")
	}
	if !Sub.IsBinary() {
		t.Error("Sub should be a binary operator")
	}
	if BitNot.IsBinary() {
		t.Error("BitNot should not be a binary operator")
	}
	if Block.IsBinary() {
		t.Error("Block should not be a binary operator")
	}
	if BlockBang.IsBinary() {
		t.Error("BlockBang should not be a binary operator")
	}
	if LogicNot.IsBinary() {
		t.Error("LogicNot should not be a binary operator")
	}
	if Neg.IsBinary() {
		t.Error("Neg should not be a binary operator")
	}
	if Pos.IsBinary() {
		t.Error("Pos should not be a binary operator")
	}
	if Splat.IsBinary() {
		t.Error("Splat should not be a binary operator")
	}
	if SplatMap.IsBinary() {
		t.Error("SplatMap should not be a binary operator")
	}
	if Arrow.IsBinary() {
		t.Error("Arrow should not be a binary operator")
	}
	if Bang.IsBinary() {
		t.Error("Bang should not be a binary operator")
	}
	if Colon.IsBinary() {
		t.Error("Colon should not be a binary operator")
	}
	if Comma.IsBinary() {
		t.Error("Comma should not be a binary operator")
	}
	if Cont.IsBinary() {
		t.Error("Cont should not be a binary operator")
	}
	if Dot.IsBinary() {
		t.Error("Dot should not be a binary operator")
	}
	if LBrace.IsBinary() {
		t.Error("LBrace should not be a binary operator")
	}
	if LBrack.IsBinary() {
		t.Error("LBrack should not be a binary operator")
	}
	if LParen.IsBinary() {
		t.Error("LParen should not be a binary operator")
	}
	if Pipe.IsBinary() {
		t.Error("Pipe should not be a binary operator")
	}
	if PipeTo.IsBinary() {
		t.Error("PipeTo should not be a binary operator")
	}
	if RBrace.IsBinary() {
		t.Error("RBrace should not be a binary operator")
	}
	if RBrack.IsBinary() {
		t.Error("RBrack should not be a binary operator")
	}
	if Rparen.IsBinary() {
		t.Error("Rparen should not be a binary operator")
	}
	if SplatAll.IsBinary() {
		t.Error("SplatAll should not be a binary operator")
	}
}

func TestIsUnary(t *testing.T) {
	if Illegal.IsUnary() {
		t.Error("Illegal should not be a unary operator")
	}
	if Comment.IsUnary() {
		t.Error("Comment should not be a unary operator")
	}
	if Space.IsUnary() {
		t.Error("Space should not be a unary operator")
	}
	if Newline.IsUnary() {
		t.Error("Newline should not be a unary operator")
	}
	if Const.IsUnary() {
		t.Error("Const should not be a unary operator")
	}
	if ConstHidden.IsUnary() {
		t.Error("ConstHidden should not be a unary operator")
	}
	if Field.IsUnary() {
		t.Error("Field should not be a unary operator")
	}
	if FieldStatic.IsUnary() {
		t.Error("FieldStaticshould not be a unary operator")
	}
	if Sink.IsUnary() {
		t.Error("Sink should not be a unary operator")
	}
	if Var.IsUnary() {
		t.Error("Var should not be a unary operator")
	}
	if VarHidden.IsUnary() {
		t.Error("VarHidden should not be a unary operator")
	}
	if Always.IsUnary() {
		t.Error("Always should not be a unary operator")
	}
	if As.IsUnary() {
		t.Error("As should not be a unary operator")
	}
	if Break.IsUnary() {
		t.Error("Break should not be a unary operator")
	}
	if Class.IsUnary() {
		t.Error("Class should not be a unary operator")
	}
	if Def.IsUnary() {
		t.Error("Def should not be a unary operator")
	}
	if Do.IsUnary() {
		t.Error("Do should not be a unary operator")
	}
	if Elif.IsUnary() {
		t.Error("Elif should not be a unary operator")
	}
	if Else.IsUnary() {
		t.Error("Else should not be a unary operator")
	}
	if For.IsUnary() {
		t.Error("For should not be a unary operator")
	}
	if From.IsUnary() {
		t.Error("From should not be a unary operator")
	}
	if If.IsUnary() {
		t.Error("If should not be a unary operator")
	}
	if In.IsUnary() {
		t.Error("In should not be a unary operator")
	}
	if Loop.IsUnary() {
		t.Error("Loop should not be a unary operator")
	}
	if Panic.IsUnary() {
		t.Error("Panic should not be a unary operator")
	}
	if Pass.IsUnary() {
		t.Error("Pass should not be a unary operator")
	}
	if Recover.IsUnary() {
		t.Error("Recover should not be a unary operator")
	}
	if Redo.IsUnary() {
		t.Error("Redo should not be a unary operator")
	}
	if Retry.IsUnary() {
		t.Error("Retry should not be a unary operator")
	}
	if Return.IsUnary() {
		t.Error("Return should not be a unary operator")
	}
	if Self.IsUnary() {
		t.Error("Self should not be a unary operator")
	}
	if Then.IsUnary() {
		t.Error("Then should not be a unary operator")
	}
	if Try.IsUnary() {
		t.Error("Try should not be a unary operator")
	}
	if Unless.IsUnary() {
		t.Error("Unless should not be a unary operator")
	}
	if Until.IsUnary() {
		t.Error("Until should not be a unary operator")
	}
	if Where.IsUnary() {
		t.Error("Where should not be a unary operator")
	}
	if While.IsUnary() {
		t.Error("While should not be a unary operator")
	}
	if Bytes.IsUnary() {
		t.Error("Bytes should not be a unary operator")
	}
	if Complex.IsUnary() {
		t.Error("Complex should not be a unary operator")
	}
	if Float.IsUnary() {
		t.Error("Float should not be a unary operator")
	}
	if Int.IsUnary() {
		t.Error("Int should not be a unary operator")
	}
	if Nil.IsUnary() {
		t.Error("Nil should not be a unary operator")
	}
	if StringDouble.IsUnary() {
		t.Error("StringDouble should not be a unary operator")
	}
	if StringDoubleBlock.IsUnary() {
		t.Error("StringDoubleBlock should not be a unary operator")
	}
	if StringSingle.IsUnary() {
		t.Error("StringSingle should not be a unary operator")
	}
	if StringSingleBlock.IsUnary() {
		t.Error("StringSingleBlock should not be a unary operator")
	}
	if Symbol.IsUnary() {
		t.Error("Symbol should not be a unary operator")
	}
	if Add.IsUnary() {
		t.Error("Add should not be a unary operator")
	}
	if Assign.IsUnary() {
		t.Error("Assign should not be a unary operator")
	}
	if BitAnd.IsUnary() {
		t.Error("BitAnd should not be a unary operator")
	}
	if BitOr.IsUnary() {
		t.Error("BitOr should not be a unary operator")
	}
	if BitXor.IsUnary() {
		t.Error("BitXor should not be a unary operator")
	}
	if Compare.IsUnary() {
		t.Error("Compare should not be a unary operator")
	}
	if Concat.IsUnary() {
		t.Error("Concat should not be a unary operator")
	}
	if Div.IsUnary() {
		t.Error("Div should not be a unary operator")
	}
	if DivFloor.IsUnary() {
		t.Error("DivFloor should not be a unary operator")
	}
	if Eq.IsUnary() {
		t.Error("Eq should not be a unary operator")
	}
	if Exp.IsUnary() {
		t.Error("Exp should not be a unary operator")
	}
	if Ge.IsUnary() {
		t.Error("Ge should not be a unary operator")
	}
	if Gt.IsUnary() {
		t.Error("Gt should not be a unary operator")
	}
	if Le.IsUnary() {
		t.Error("Le should not be a unary operator")
	}
	if LogicAnd.IsUnary() {
		t.Error("LogicAnd should not be a unary operator")
	}
	if LogicOr.IsUnary() {
		t.Error("LogicOr should not be a unary operator")
	}
	if Lt.IsUnary() {
		t.Error("Lt should not be a unary operator")
	}
	if Mod.IsUnary() {
		t.Error("Mod should not be a unary operator")
	}
	if Mul.IsUnary() {
		t.Error("Mul should not be a unary operator")
	}
	if Ne.IsUnary() {
		t.Error("Ne should not be a unary operator")
	}
	if RangeExc.IsUnary() {
		t.Error("RangeExc should not be a unary operator")
	}
	if RangeInc.IsUnary() {
		t.Error("RangeInc should not be a unary operator")
	}
	if ShiftLeft.IsUnary() {
		t.Error("ShiftLeft should not be a unary operator")
	}
	if ShiftRight.IsUnary() {
		t.Error("ShiftRight should not be a unary operator")
	}
	if Sub.IsUnary() {
		t.Error("Sub should not be a unary operator")
	}
	if !BitNot.IsUnary() {
		t.Error("BitNot should be a unary operator")
	}
	if !Block.IsUnary() {
		t.Error("Block should be a unary operator")
	}
	if !BlockBang.IsUnary() {
		t.Error("BlockBang should be a unary operator")
	}
	if !LogicNot.IsUnary() {
		t.Error("LogicNot should be a unary operator")
	}
	if !Neg.IsUnary() {
		t.Error("Neg should be a unary operator")
	}
	if !Pos.IsUnary() {
		t.Error("Pos should be a unary operator")
	}
	if !Splat.IsUnary() {
		t.Error("Splat should be a unary operator")
	}
	if !SplatMap.IsUnary() {
		t.Error("SplatMap should be a unary operator")
	}
	if Arrow.IsUnary() {
		t.Error("Arrow should not be a unary operator")
	}
	if Bang.IsUnary() {
		t.Error("Bang should not be a unary operator")
	}
	if Colon.IsUnary() {
		t.Error("Colon should not be a unary operator")
	}
	if Comma.IsUnary() {
		t.Error("Comma should not be a unary operator")
	}
	if Cont.IsUnary() {
		t.Error("Cont should not be a unary operator")
	}
	if Dot.IsUnary() {
		t.Error("Dot should not be a unary operator")
	}
	if LBrace.IsUnary() {
		t.Error("LBrace should not be a unary operator")
	}
	if LBrack.IsUnary() {
		t.Error("LBrack should not be a unary operator")
	}
	if LParen.IsUnary() {
		t.Error("LParen should not be a unary operator")
	}
	if Pipe.IsUnary() {
		t.Error("Pipe should not be a unary operator")
	}
	if PipeTo.IsUnary() {
		t.Error("PipeTo should not be a unary operator")
	}
	if RBrace.IsUnary() {
		t.Error("RBrace should not be a unary operator")
	}
	if RBrack.IsUnary() {
		t.Error("RBrack should not be a unary operator")
	}
	if Rparen.IsUnary() {
		t.Error("Rparen should not be a unary operator")
	}
	if SplatAll.IsUnary() {
		t.Error("SplatAll should not be a unary operator")
	}
}

func TestIsPunc(t *testing.T) {
	if Illegal.IsPunc() {
		t.Error("Illegal should not be a punctuation")
	}
	if Comment.IsPunc() {
		t.Error("Comment should not be a punctuation")
	}
	if Space.IsPunc() {
		t.Error("Space should not be a punctuation")
	}
	if Newline.IsPunc() {
		t.Error("Newline should not be a punctuation")
	}
	if Const.IsPunc() {
		t.Error("Const should not be a punctuation")
	}
	if ConstHidden.IsPunc() {
		t.Error("ConstHidden should not be a punctuation")
	}
	if Field.IsPunc() {
		t.Error("Field should not be a punctuation")
	}
	if FieldStatic.IsPunc() {
		t.Error("FieldStaticshould not be a punctuation")
	}
	if Sink.IsPunc() {
		t.Error("Sink should not be a punctuation")
	}
	if Var.IsPunc() {
		t.Error("Var should not be a punctuation")
	}
	if VarHidden.IsPunc() {
		t.Error("VarHidden should not be a punctuation")
	}
	if Always.IsPunc() {
		t.Error("Always should not be a punctuation")
	}
	if As.IsPunc() {
		t.Error("As should not be a punctuation")
	}
	if Break.IsPunc() {
		t.Error("Break should not be a punctuation")
	}
	if Class.IsPunc() {
		t.Error("Class should not be a punctuation")
	}
	if Def.IsPunc() {
		t.Error("Def should not be a punctuation")
	}
	if Do.IsPunc() {
		t.Error("Do should not be a punctuation")
	}
	if Elif.IsPunc() {
		t.Error("Elif should not be a punctuation")
	}
	if Else.IsPunc() {
		t.Error("Else should not be a punctuation")
	}
	if For.IsPunc() {
		t.Error("For should not be a punctuation")
	}
	if From.IsPunc() {
		t.Error("From should not be a punctuation")
	}
	if If.IsPunc() {
		t.Error("If should not be a punctuation")
	}
	if In.IsPunc() {
		t.Error("In should not be a punctuation")
	}
	if Loop.IsPunc() {
		t.Error("Loop should not be a punctuation")
	}
	if Panic.IsPunc() {
		t.Error("Panic should not be a punctuation")
	}
	if Pass.IsPunc() {
		t.Error("Pass should not be a punctuation")
	}
	if Recover.IsPunc() {
		t.Error("Recover should not be a punctuation")
	}
	if Redo.IsPunc() {
		t.Error("Redo should not be a punctuation")
	}
	if Retry.IsPunc() {
		t.Error("Retry should not be a punctuation")
	}
	if Return.IsPunc() {
		t.Error("Return should not be a punctuation")
	}
	if Self.IsPunc() {
		t.Error("Self should not be a punctuation")
	}
	if Then.IsPunc() {
		t.Error("Then should not be a punctuation")
	}
	if Try.IsPunc() {
		t.Error("Try should not be a punctuation")
	}
	if Unless.IsPunc() {
		t.Error("Unless should not be a punctuation")
	}
	if Until.IsPunc() {
		t.Error("Until should not be a punctuation")
	}
	if Where.IsPunc() {
		t.Error("Where should not be a punctuation")
	}
	if While.IsPunc() {
		t.Error("While should not be a punctuation")
	}
	if Bytes.IsPunc() {
		t.Error("Bytes should not be a punctuation")
	}
	if Complex.IsPunc() {
		t.Error("Complex should not be a punctuation")
	}
	if Float.IsPunc() {
		t.Error("Float should not be a punctuation")
	}
	if Int.IsPunc() {
		t.Error("Int should not be a punctuation")
	}
	if Nil.IsPunc() {
		t.Error("Nil should not be a punctuation")
	}
	if StringDouble.IsPunc() {
		t.Error("StringDouble should not be a punctuation")
	}
	if StringDoubleBlock.IsPunc() {
		t.Error("StringDoubleBlock should not be a punctuation")
	}
	if StringSingle.IsPunc() {
		t.Error("StringSingle should not be a punctuation")
	}
	if StringSingleBlock.IsPunc() {
		t.Error("StringSingleBlock should not be a punctuation")
	}
	if Symbol.IsPunc() {
		t.Error("Symbol should not be a punctuation")
	}
	if Add.IsPunc() {
		t.Error("Add should not be a punctuation")
	}
	if Assign.IsPunc() {
		t.Error("Assign should not be a punctuation")
	}
	if BitAnd.IsPunc() {
		t.Error("BitAnd should not be a punctuation")
	}
	if BitOr.IsPunc() {
		t.Error("BitOr should not be a punctuation")
	}
	if BitXor.IsPunc() {
		t.Error("BitXor should not be a punctuation")
	}
	if Compare.IsPunc() {
		t.Error("Compare should not be a punctuation")
	}
	if Concat.IsPunc() {
		t.Error("Concat should not be a punctuation")
	}
	if Div.IsPunc() {
		t.Error("Div should not be a punctuation")
	}
	if DivFloor.IsPunc() {
		t.Error("DivFloor should not be a punctuation")
	}
	if Eq.IsPunc() {
		t.Error("Eq should not be a punctuation")
	}
	if Exp.IsPunc() {
		t.Error("Exp should not be a punctuation")
	}
	if Ge.IsPunc() {
		t.Error("Ge should not be a punctuation")
	}
	if Gt.IsPunc() {
		t.Error("Gt should not be a punctuation")
	}
	if Le.IsPunc() {
		t.Error("Le should not be a punctuation")
	}
	if LogicAnd.IsPunc() {
		t.Error("LogicAnd should not be a punctuation")
	}
	if LogicOr.IsPunc() {
		t.Error("LogicOr should not be a punctuation")
	}
	if Lt.IsPunc() {
		t.Error("Lt should not be a punctuation")
	}
	if Mod.IsPunc() {
		t.Error("Mod should not be a punctuation")
	}
	if Mul.IsPunc() {
		t.Error("Mul should not be a punctuation")
	}
	if Ne.IsPunc() {
		t.Error("Ne should not be a punctuation")
	}
	if RangeExc.IsPunc() {
		t.Error("RangeExc should not be a punctuation")
	}
	if RangeInc.IsPunc() {
		t.Error("RangeInc should not be a punctuation")
	}
	if ShiftLeft.IsPunc() {
		t.Error("ShiftLeft should not be a punctuation")
	}
	if ShiftRight.IsPunc() {
		t.Error("ShiftRight should not be a punctuation")
	}
	if Sub.IsPunc() {
		t.Error("Sub should not be a punctuation")
	}
	if BitNot.IsPunc() {
		t.Error("BitNot should not be a punctuation")
	}
	if Block.IsPunc() {
		t.Error("Block should not be a punctuation")
	}
	if BlockBang.IsPunc() {
		t.Error("BlockBang should not be a punctuation")
	}
	if LogicNot.IsPunc() {
		t.Error("LogicNot should not be a punctuation")
	}
	if Neg.IsPunc() {
		t.Error("Neg should not be a punctuation")
	}
	if Pos.IsPunc() {
		t.Error("Pos should not be a punctuation")
	}
	if Splat.IsPunc() {
		t.Error("Splat should not be a punctuation")
	}
	if SplatMap.IsPunc() {
		t.Error("SplatMap should not be a punctuation")
	}
	if !Arrow.IsPunc() {
		t.Error("Arrow should be a punctuation")
	}
	if !Bang.IsPunc() {
		t.Error("Bang should be a punctuation")
	}
	if !Colon.IsPunc() {
		t.Error("Colon should be a punctuation")
	}
	if !Comma.IsPunc() {
		t.Error("Comma should be a punctuation")
	}
	if !Cont.IsPunc() {
		t.Error("Cont should be a punctuation")
	}
	if !Dot.IsPunc() {
		t.Error("Dot should be a punctuation")
	}
	if !LBrace.IsPunc() {
		t.Error("LBrace should be a punctuation")
	}
	if !LBrack.IsPunc() {
		t.Error("LBrack should be a punctuation")
	}
	if !LParen.IsPunc() {
		t.Error("LParen should be a punctuation")
	}
	if !Pipe.IsPunc() {
		t.Error("Pipe should be a punctuation")
	}
	if !PipeTo.IsPunc() {
		t.Error("PipeTo should be a punctuation")
	}
	if !RBrace.IsPunc() {
		t.Error("RBrace should be a punctuation")
	}
	if !RBrack.IsPunc() {
		t.Error("RBrack should be a punctuation")
	}
	if !Rparen.IsPunc() {
		t.Error("Rparen should be a punctuation")
	}
	if !SplatAll.IsPunc() {
		t.Error("SplatAll should be a punctuation")
	}
}

func TestIsConstant(t *testing.T) {
	if Illegal.IsConstant() {
		t.Error("Illegal should not be a constant")
	}
	if Comment.IsConstant() {
		t.Error("Comment should not be a constant")
	}
	if Space.IsConstant() {
		t.Error("Space should not be a constant")
	}
	if Newline.IsConstant() {
		t.Error("Newline should not be a constant")
	}
	if !Const.IsConstant() {
		t.Error("Const should be a constant")
	}
	if !ConstHidden.IsConstant() {
		t.Error("ConstHidden should be a constant")
	}
	if Field.IsConstant() {
		t.Error("Field should not be a constant")
	}
	if FieldStatic.IsConstant() {
		t.Error("FieldStaticshould not be a constant")
	}
	if Sink.IsConstant() {
		t.Error("Sink should not be a constant")
	}
	if Var.IsConstant() {
		t.Error("Var should not be a constant")
	}
	if VarHidden.IsConstant() {
		t.Error("VarHidden should not be a constant")
	}
	if Always.IsConstant() {
		t.Error("Always should not be a constant")
	}
	if As.IsConstant() {
		t.Error("As should not be a constant")
	}
	if Break.IsConstant() {
		t.Error("Break should not be a constant")
	}
	if Class.IsConstant() {
		t.Error("Class should not be a constant")
	}
	if Def.IsConstant() {
		t.Error("Def should not be a constant")
	}
	if Do.IsConstant() {
		t.Error("Do should not be a constant")
	}
	if Elif.IsConstant() {
		t.Error("Elif should not be a constant")
	}
	if Else.IsConstant() {
		t.Error("Else should not be a constant")
	}
	if For.IsConstant() {
		t.Error("For should not be a constant")
	}
	if From.IsConstant() {
		t.Error("From should not be a constant")
	}
	if If.IsConstant() {
		t.Error("If should not be a constant")
	}
	if In.IsConstant() {
		t.Error("In should not be a constant")
	}
	if Loop.IsConstant() {
		t.Error("Loop should not be a constant")
	}
	if Panic.IsConstant() {
		t.Error("Panic should not be a constant")
	}
	if Pass.IsConstant() {
		t.Error("Pass should not be a constant")
	}
	if Recover.IsConstant() {
		t.Error("Recover should not be a constant")
	}
	if Redo.IsConstant() {
		t.Error("Redo should not be a constant")
	}
	if Retry.IsConstant() {
		t.Error("Retry should not be a constant")
	}
	if Return.IsConstant() {
		t.Error("Return should not be a constant")
	}
	if Self.IsConstant() {
		t.Error("Self should not be a constant")
	}
	if Then.IsConstant() {
		t.Error("Then should not be a constant")
	}
	if Try.IsConstant() {
		t.Error("Try should not be a constant")
	}
	if Unless.IsConstant() {
		t.Error("Unless should not be a constant")
	}
	if Until.IsConstant() {
		t.Error("Until should not be a constant")
	}
	if Where.IsConstant() {
		t.Error("Where should not be a constant")
	}
	if While.IsConstant() {
		t.Error("While should not be a constant")
	}
	if Bytes.IsConstant() {
		t.Error("Bytes should not be a constant")
	}
	if Complex.IsConstant() {
		t.Error("Complex should not be a constant")
	}
	if Float.IsConstant() {
		t.Error("Float should not be a constant")
	}
	if Int.IsConstant() {
		t.Error("Int should not be a constant")
	}
	if Nil.IsConstant() {
		t.Error("Nil should not be a constant")
	}
	if StringDouble.IsConstant() {
		t.Error("StringDouble should not be a constant")
	}
	if StringDoubleBlock.IsConstant() {
		t.Error("StringDoubleBlock should not be a constant")
	}
	if StringSingle.IsConstant() {
		t.Error("StringSingle should not be a constant")
	}
	if StringSingleBlock.IsConstant() {
		t.Error("StringSingleBlock should not be a constant")
	}
	if Symbol.IsConstant() {
		t.Error("Symbol should not be a constant")
	}
	if Add.IsConstant() {
		t.Error("Add should not be a constant")
	}
	if Assign.IsConstant() {
		t.Error("Assign should not be a constant")
	}
	if BitAnd.IsConstant() {
		t.Error("BitAnd should not be a constant")
	}
	if BitOr.IsConstant() {
		t.Error("BitOr should not be a constant")
	}
	if BitXor.IsConstant() {
		t.Error("BitXor should not be a constant")
	}
	if Compare.IsConstant() {
		t.Error("Compare should not be a constant")
	}
	if Concat.IsConstant() {
		t.Error("Concat should not be a constant")
	}
	if Div.IsConstant() {
		t.Error("Div should not be a constant")
	}
	if DivFloor.IsConstant() {
		t.Error("DivFloor should not be a constant")
	}
	if Eq.IsConstant() {
		t.Error("Eq should not be a constant")
	}
	if Exp.IsConstant() {
		t.Error("Exp should not be a constant")
	}
	if Ge.IsConstant() {
		t.Error("Ge should not be a constant")
	}
	if Gt.IsConstant() {
		t.Error("Gt should not be a constant")
	}
	if Le.IsConstant() {
		t.Error("Le should not be a constant")
	}
	if LogicAnd.IsConstant() {
		t.Error("LogicAnd should not be a constant")
	}
	if LogicOr.IsConstant() {
		t.Error("LogicOr should not be a constant")
	}
	if Lt.IsConstant() {
		t.Error("Lt should not be a constant")
	}
	if Mod.IsConstant() {
		t.Error("Mod should not be a constant")
	}
	if Mul.IsConstant() {
		t.Error("Mul should not be a constant")
	}
	if Ne.IsConstant() {
		t.Error("Ne should not be a constant")
	}
	if RangeExc.IsConstant() {
		t.Error("RangeExc should not be a constant")
	}
	if RangeInc.IsConstant() {
		t.Error("RangeInc should not be a constant")
	}
	if ShiftLeft.IsConstant() {
		t.Error("ShiftLeft should not be a constant")
	}
	if ShiftRight.IsConstant() {
		t.Error("ShiftRight should not be a constant")
	}
	if Sub.IsConstant() {
		t.Error("Sub should not be a constant")
	}
	if BitNot.IsConstant() {
		t.Error("BitNot should not be a constant")
	}
	if Block.IsConstant() {
		t.Error("Block should not be a constant")
	}
	if BlockBang.IsConstant() {
		t.Error("BlockBang should not be a constant")
	}
	if LogicNot.IsConstant() {
		t.Error("LogicNot should not be a constant")
	}
	if Neg.IsConstant() {
		t.Error("Neg should not be a constant")
	}
	if Pos.IsConstant() {
		t.Error("Pos should not be a constant")
	}
	if Splat.IsConstant() {
		t.Error("Splat should not be a constant")
	}
	if SplatMap.IsConstant() {
		t.Error("SplatMap should not be a constant")
	}
	if Arrow.IsConstant() {
		t.Error("Arrow should not be a constant")
	}
	if Bang.IsConstant() {
		t.Error("Bang should not be a constant")
	}
	if Colon.IsConstant() {
		t.Error("Colon should not be a constant")
	}
	if Comma.IsConstant() {
		t.Error("Comma should not be a constant")
	}
	if Cont.IsConstant() {
		t.Error("Cont should not be a constant")
	}
	if Dot.IsConstant() {
		t.Error("Dot should not be a constant")
	}
	if LBrace.IsConstant() {
		t.Error("LBrace should not be a constant")
	}
	if LBrack.IsConstant() {
		t.Error("LBrack should not be a constant")
	}
	if LParen.IsConstant() {
		t.Error("LParen should not be a constant")
	}
	if Pipe.IsConstant() {
		t.Error("Pipe should not be a constant")
	}
	if PipeTo.IsConstant() {
		t.Error("PipeTo should not be a constant")
	}
	if RBrace.IsConstant() {
		t.Error("RBrace should not be a constant")
	}
	if RBrack.IsConstant() {
		t.Error("RBrack should not be a constant")
	}
	if Rparen.IsConstant() {
		t.Error("Rparen should not be a constant")
	}
	if SplatAll.IsConstant() {
		t.Error("SplatAll should not be a constant")
	}
}

func TestIsField(t *testing.T) {
	if Illegal.IsField() {
		t.Error("Illegal should not be a field")
	}
	if Comment.IsField() {
		t.Error("Comment should not be a field")
	}
	if Space.IsField() {
		t.Error("Space should not be a field")
	}
	if Newline.IsField() {
		t.Error("Newline should not be a field")
	}
	if Const.IsField() {
		t.Error("Const should not be a field")
	}
	if ConstHidden.IsField() {
		t.Error("ConstHidden should not be a field")
	}
	if !Field.IsField() {
		t.Error("Field should be a field")
	}
	if !FieldStatic.IsField() {
		t.Error("FieldStaticshould be a field")
	}
	if Sink.IsField() {
		t.Error("Sink should not be a field")
	}
	if Var.IsField() {
		t.Error("Var should not be a field")
	}
	if VarHidden.IsField() {
		t.Error("VarHidden should not be a field")
	}
	if Always.IsField() {
		t.Error("Always should not be a field")
	}
	if As.IsField() {
		t.Error("As should not be a field")
	}
	if Break.IsField() {
		t.Error("Break should not be a field")
	}
	if Class.IsField() {
		t.Error("Class should not be a field")
	}
	if Def.IsField() {
		t.Error("Def should not be a field")
	}
	if Do.IsField() {
		t.Error("Do should not be a field")
	}
	if Elif.IsField() {
		t.Error("Elif should not be a field")
	}
	if Else.IsField() {
		t.Error("Else should not be a field")
	}
	if For.IsField() {
		t.Error("For should not be a field")
	}
	if From.IsField() {
		t.Error("From should not be a field")
	}
	if If.IsField() {
		t.Error("If should not be a field")
	}
	if In.IsField() {
		t.Error("In should not be a field")
	}
	if Loop.IsField() {
		t.Error("Loop should not be a field")
	}
	if Panic.IsField() {
		t.Error("Panic should not be a field")
	}
	if Pass.IsField() {
		t.Error("Pass should not be a field")
	}
	if Recover.IsField() {
		t.Error("Recover should not be a field")
	}
	if Redo.IsField() {
		t.Error("Redo should not be a field")
	}
	if Retry.IsField() {
		t.Error("Retry should not be a field")
	}
	if Return.IsField() {
		t.Error("Return should not be a field")
	}
	if Self.IsField() {
		t.Error("Self should not be a field")
	}
	if Then.IsField() {
		t.Error("Then should not be a field")
	}
	if Try.IsField() {
		t.Error("Try should not be a field")
	}
	if Unless.IsField() {
		t.Error("Unless should not be a field")
	}
	if Until.IsField() {
		t.Error("Until should not be a field")
	}
	if Where.IsField() {
		t.Error("Where should not be a field")
	}
	if While.IsField() {
		t.Error("While should not be a field")
	}
	if Bytes.IsField() {
		t.Error("Bytes should not be a field")
	}
	if Complex.IsField() {
		t.Error("Complex should not be a field")
	}
	if Float.IsField() {
		t.Error("Float should not be a field")
	}
	if Int.IsField() {
		t.Error("Int should not be a field")
	}
	if Nil.IsField() {
		t.Error("Nil should not be a field")
	}
	if StringDouble.IsField() {
		t.Error("StringDouble should not be a field")
	}
	if StringDoubleBlock.IsField() {
		t.Error("StringDoubleBlock should not be a field")
	}
	if StringSingle.IsField() {
		t.Error("StringSingle should not be a field")
	}
	if StringSingleBlock.IsField() {
		t.Error("StringSingleBlock should not be a field")
	}
	if Symbol.IsField() {
		t.Error("Symbol should not be a field")
	}
	if Add.IsField() {
		t.Error("Add should not be a field")
	}
	if Assign.IsField() {
		t.Error("Assign should not be a field")
	}
	if BitAnd.IsField() {
		t.Error("BitAnd should not be a field")
	}
	if BitOr.IsField() {
		t.Error("BitOr should not be a field")
	}
	if BitXor.IsField() {
		t.Error("BitXor should not be a field")
	}
	if Compare.IsField() {
		t.Error("Compare should not be a field")
	}
	if Concat.IsField() {
		t.Error("Concat should not be a field")
	}
	if Div.IsField() {
		t.Error("Div should not be a field")
	}
	if DivFloor.IsField() {
		t.Error("DivFloor should not be a field")
	}
	if Eq.IsField() {
		t.Error("Eq should not be a field")
	}
	if Exp.IsField() {
		t.Error("Exp should not be a field")
	}
	if Ge.IsField() {
		t.Error("Ge should not be a field")
	}
	if Gt.IsField() {
		t.Error("Gt should not be a field")
	}
	if Le.IsField() {
		t.Error("Le should not be a field")
	}
	if LogicAnd.IsField() {
		t.Error("LogicAnd should not be a field")
	}
	if LogicOr.IsField() {
		t.Error("LogicOr should not be a field")
	}
	if Lt.IsField() {
		t.Error("Lt should not be a field")
	}
	if Mod.IsField() {
		t.Error("Mod should not be a field")
	}
	if Mul.IsField() {
		t.Error("Mul should not be a field")
	}
	if Ne.IsField() {
		t.Error("Ne should not be a field")
	}
	if RangeExc.IsField() {
		t.Error("RangeExc should not be a field")
	}
	if RangeInc.IsField() {
		t.Error("RangeInc should not be a field")
	}
	if ShiftLeft.IsField() {
		t.Error("ShiftLeft should not be a field")
	}
	if ShiftRight.IsField() {
		t.Error("ShiftRight should not be a field")
	}
	if Sub.IsField() {
		t.Error("Sub should not be a field")
	}
	if BitNot.IsField() {
		t.Error("BitNot should not be a field")
	}
	if Block.IsField() {
		t.Error("Block should not be a field")
	}
	if BlockBang.IsField() {
		t.Error("BlockBang should not be a field")
	}
	if LogicNot.IsField() {
		t.Error("LogicNot should not be a field")
	}
	if Neg.IsField() {
		t.Error("Neg should not be a field")
	}
	if Pos.IsField() {
		t.Error("Pos should not be a field")
	}
	if Splat.IsField() {
		t.Error("Splat should not be a field")
	}
	if SplatMap.IsField() {
		t.Error("SplatMap should not be a field")
	}
	if Arrow.IsField() {
		t.Error("Arrow should not be a field")
	}
	if Bang.IsField() {
		t.Error("Bang should not be a field")
	}
	if Colon.IsField() {
		t.Error("Colon should not be a field")
	}
	if Comma.IsField() {
		t.Error("Comma should not be a field")
	}
	if Cont.IsField() {
		t.Error("Cont should not be a field")
	}
	if Dot.IsField() {
		t.Error("Dot should not be a field")
	}
	if LBrace.IsField() {
		t.Error("LBrace should not be a field")
	}
	if LBrack.IsField() {
		t.Error("LBrack should not be a field")
	}
	if LParen.IsField() {
		t.Error("LParen should not be a field")
	}
	if Pipe.IsField() {
		t.Error("Pipe should not be a field")
	}
	if PipeTo.IsField() {
		t.Error("PipeTo should not be a field")
	}
	if RBrace.IsField() {
		t.Error("RBrace should not be a field")
	}
	if RBrack.IsField() {
		t.Error("RBrack should not be a field")
	}
	if Rparen.IsField() {
		t.Error("Rparen should not be a field")
	}
	if SplatAll.IsField() {
		t.Error("SplatAll should not be a field")
	}
}

func TestIsHidden(t *testing.T) {
	if Illegal.IsHidden() {
		t.Error("Illegal should not be hidden")
	}
	if Comment.IsHidden() {
		t.Error("Comment should not be hidden")
	}
	if Space.IsHidden() {
		t.Error("Space should not be hidden")
	}
	if Newline.IsHidden() {
		t.Error("Newline should not be hidden")
	}
	if Const.IsHidden() {
		t.Error("Const should not be hidden")
	}
	if !ConstHidden.IsHidden() {
		t.Error("ConstHidden should be hidden")
	}
	if Field.IsHidden() {
		t.Error("Field should not be hidden")
	}
	if FieldStatic.IsHidden() {
		t.Error("FieldStaticshould not be hidden")
	}
	if Sink.IsHidden() {
		t.Error("Sink should not be hidden")
	}
	if Var.IsHidden() {
		t.Error("Var should not be hidden")
	}
	if !VarHidden.IsHidden() {
		t.Error("VarHidden should be hidden")
	}
	if Always.IsHidden() {
		t.Error("Always should not be hidden")
	}
	if As.IsHidden() {
		t.Error("As should not be hidden")
	}
	if Break.IsHidden() {
		t.Error("Break should not be hidden")
	}
	if Class.IsHidden() {
		t.Error("Class should not be hidden")
	}
	if Def.IsHidden() {
		t.Error("Def should not be hidden")
	}
	if Do.IsHidden() {
		t.Error("Do should not be hidden")
	}
	if Elif.IsHidden() {
		t.Error("Elif should not be hidden")
	}
	if Else.IsHidden() {
		t.Error("Else should not be hidden")
	}
	if For.IsHidden() {
		t.Error("For should not be hidden")
	}
	if From.IsHidden() {
		t.Error("From should not be hidden")
	}
	if If.IsHidden() {
		t.Error("If should not be hidden")
	}
	if In.IsHidden() {
		t.Error("In should not be hidden")
	}
	if Loop.IsHidden() {
		t.Error("Loop should not be hidden")
	}
	if Panic.IsHidden() {
		t.Error("Panic should not be hidden")
	}
	if Pass.IsHidden() {
		t.Error("Pass should not be hidden")
	}
	if Recover.IsHidden() {
		t.Error("Recover should not be hidden")
	}
	if Redo.IsHidden() {
		t.Error("Redo should not be hidden")
	}
	if Retry.IsHidden() {
		t.Error("Retry should not be hidden")
	}
	if Return.IsHidden() {
		t.Error("Return should not be hidden")
	}
	if Self.IsHidden() {
		t.Error("Self should not be hidden")
	}
	if Then.IsHidden() {
		t.Error("Then should not be hidden")
	}
	if Try.IsHidden() {
		t.Error("Try should not be hidden")
	}
	if Unless.IsHidden() {
		t.Error("Unless should not be hidden")
	}
	if Until.IsHidden() {
		t.Error("Until should not be hidden")
	}
	if Where.IsHidden() {
		t.Error("Where should not be hidden")
	}
	if While.IsHidden() {
		t.Error("While should not be hidden")
	}
	if Bytes.IsHidden() {
		t.Error("Bytes should not be hidden")
	}
	if Complex.IsHidden() {
		t.Error("Complex should not be hidden")
	}
	if Float.IsHidden() {
		t.Error("Float should not be hidden")
	}
	if Int.IsHidden() {
		t.Error("Int should not be hidden")
	}
	if Nil.IsHidden() {
		t.Error("Nil should not be hidden")
	}
	if StringDouble.IsHidden() {
		t.Error("StringDouble should not be hidden")
	}
	if StringDoubleBlock.IsHidden() {
		t.Error("StringDoubleBlock should not be hidden")
	}
	if StringSingle.IsHidden() {
		t.Error("StringSingle should not be hidden")
	}
	if StringSingleBlock.IsHidden() {
		t.Error("StringSingleBlock should not be hidden")
	}
	if Symbol.IsHidden() {
		t.Error("Symbol should not be hidden")
	}
	if Add.IsHidden() {
		t.Error("Add should not be hidden")
	}
	if Assign.IsHidden() {
		t.Error("Assign should not be hidden")
	}
	if BitAnd.IsHidden() {
		t.Error("BitAnd should not be hidden")
	}
	if BitOr.IsHidden() {
		t.Error("BitOr should not be hidden")
	}
	if BitXor.IsHidden() {
		t.Error("BitXor should not be hidden")
	}
	if Compare.IsHidden() {
		t.Error("Compare should not be hidden")
	}
	if Concat.IsHidden() {
		t.Error("Concat should not be hidden")
	}
	if Div.IsHidden() {
		t.Error("Div should not be hidden")
	}
	if DivFloor.IsHidden() {
		t.Error("DivFloor should not be hidden")
	}
	if Eq.IsHidden() {
		t.Error("Eq should not be hidden")
	}
	if Exp.IsHidden() {
		t.Error("Exp should not be hidden")
	}
	if Ge.IsHidden() {
		t.Error("Ge should not be hidden")
	}
	if Gt.IsHidden() {
		t.Error("Gt should not be hidden")
	}
	if Le.IsHidden() {
		t.Error("Le should not be hidden")
	}
	if LogicAnd.IsHidden() {
		t.Error("LogicAnd should not be hidden")
	}
	if LogicOr.IsHidden() {
		t.Error("LogicOr should not be hidden")
	}
	if Lt.IsHidden() {
		t.Error("Lt should not be hidden")
	}
	if Mod.IsHidden() {
		t.Error("Mod should not be hidden")
	}
	if Mul.IsHidden() {
		t.Error("Mul should not be hidden")
	}
	if Ne.IsHidden() {
		t.Error("Ne should not be hidden")
	}
	if RangeExc.IsHidden() {
		t.Error("RangeExc should not be hidden")
	}
	if RangeInc.IsHidden() {
		t.Error("RangeInc should not be hidden")
	}
	if ShiftLeft.IsHidden() {
		t.Error("ShiftLeft should not be hidden")
	}
	if ShiftRight.IsHidden() {
		t.Error("ShiftRight should not be hidden")
	}
	if Sub.IsHidden() {
		t.Error("Sub should not be hidden")
	}
	if BitNot.IsHidden() {
		t.Error("BitNot should not be hidden")
	}
	if Block.IsHidden() {
		t.Error("Block should not be hidden")
	}
	if BlockBang.IsHidden() {
		t.Error("BlockBang should not be hidden")
	}
	if LogicNot.IsHidden() {
		t.Error("LogicNot should not be hidden")
	}
	if Neg.IsHidden() {
		t.Error("Neg should not be hidden")
	}
	if Pos.IsHidden() {
		t.Error("Pos should not be hidden")
	}
	if Splat.IsHidden() {
		t.Error("Splat should not be hidden")
	}
	if SplatMap.IsHidden() {
		t.Error("SplatMap should not be hidden")
	}
	if Arrow.IsHidden() {
		t.Error("Arrow should not be hidden")
	}
	if Bang.IsHidden() {
		t.Error("Bang should not be hidden")
	}
	if Colon.IsHidden() {
		t.Error("Colon should not be hidden")
	}
	if Comma.IsHidden() {
		t.Error("Comma should not be hidden")
	}
	if Cont.IsHidden() {
		t.Error("Cont should not be hidden")
	}
	if Dot.IsHidden() {
		t.Error("Dot should not be hidden")
	}
	if LBrace.IsHidden() {
		t.Error("LBrace should not be hidden")
	}
	if LBrack.IsHidden() {
		t.Error("LBrack should not be hidden")
	}
	if LParen.IsHidden() {
		t.Error("LParen should not be hidden")
	}
	if Pipe.IsHidden() {
		t.Error("Pipe should not be hidden")
	}
	if PipeTo.IsHidden() {
		t.Error("PipeTo should not be hidden")
	}
	if RBrace.IsHidden() {
		t.Error("RBrace should not be hidden")
	}
	if RBrack.IsHidden() {
		t.Error("RBrack should not be hidden")
	}
	if Rparen.IsHidden() {
		t.Error("Rparen should not be hidden")
	}
	if SplatAll.IsHidden() {
		t.Error("SplatAll should not be hidden")
	}
}
