package main

import (
	"github.com/arpitbbhayani/gbasic/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	lexer := parser.NewBASICLexer(antlr.NewInputStream("10 PRINT \"Hello, world!\""))
	parser := parser.NewBASICParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.Program()
}
