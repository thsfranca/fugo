package main

import (
	"github.com/thsfranca/fugo/core"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/thsfranca/fugo/core/parser"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := parser.NewFugoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	fugoParser := parser.NewFugoParser(stream)
	fugoParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := fugoParser.Program()
	treeShapeListener := core.NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(treeShapeListener, tree)
}
