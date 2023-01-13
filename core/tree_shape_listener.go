package core

import (
	"github.com/thsfranca/fugo/core/parser"
)

type TreeShapeListener struct {
	*parser.BaseFugoListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterFn(ctx *parser.FnContext) {
	identifiers[ctx.IDENTIFIER(0).GetText()].(func(text string))("teste")
}
