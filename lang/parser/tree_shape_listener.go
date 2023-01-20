package parser

import (
	"github.com/thsfranca/fugo/core"
	"github.com/thsfranca/fugo/lang/types"
)

type TreeShapeListener struct {
	*core.BaseFugoListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterList_(ctx *core.List_Context) {
	parser := NewParser(types.Symbol(ctx.Forms().GetChildren()[0].(*core.FormContext).GetText()))
	parser.Parse(ctx.Forms().GetChildren())
}
