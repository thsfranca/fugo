package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/thsfranca/fugo/core"
	"github.com/thsfranca/fugo/lang/expr"
	"github.com/thsfranca/fugo/lang/util"
)

type Def struct{}

// Parse parses def special form
// (def x) or (def x initexpr) or (def x "docstring" initexpr)
func (d *Def) Parse(forms []antlr.Tree) {
	if len(forms) > 4 {
		panic("too many arguments to def")
	}

	if len(forms) < 3 {
		panic("too few arguments to def")
	}

	if util.SecondForm(forms).GetStop().GetTokenType() != core.FugoLexerSYMBOL {
		panic("first argument to def must be a Symbol")
	}

	expr.DefExpr(util.SecondForm(forms).GetText(), util.LastForm(forms).GetText())
}
