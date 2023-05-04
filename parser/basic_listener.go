// Code generated from parser/BASIC.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // BASIC

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BASICListener is a complete listener for a parse tree produced by BASICParser.
type BASICListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
