package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/thsfranca/fugo/lang"
	"github.com/thsfranca/fugo/lang/types"
)

type Parser interface {
	Parse(forms []antlr.Tree)
}

func NewParser(symbol types.Symbol) Parser {
	switch symbol {
	case types.Symbol(lang.DEF):
		return &Def{}
	}
	return nil
}
