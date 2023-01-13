package core

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TreeShapeListener struct {
	*antlr.BaseParseTreeListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterFunction(ctx antlr.ParserRuleContext) {
	fmt.Println("HERE")
}
