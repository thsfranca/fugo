package main

import (
	"github.com/thsfranca/fugo/core"
	"github.com/thsfranca/fugo/lang/parser"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := core.NewFugoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	fugoParser := core.NewFugoParser(stream)
	fugoParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := fugoParser.Form()
	treeShapeListener := parser.NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(treeShapeListener, tree)
}
