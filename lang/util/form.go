package util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/thsfranca/fugo/core"
)

func FirstForm(forms []antlr.Tree) *core.FormContext {
	return forms[0].(*core.FormContext)
}

func SecondForm(forms []antlr.Tree) *core.FormContext {
	return forms[1].(*core.FormContext)
}

func LastForm(forms []antlr.Tree) *core.FormContext {
	return forms[len(forms)-1].(*core.FormContext)
}
