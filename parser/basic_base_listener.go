// Code generated from parser/BASIC.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // BASIC

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseBASICListener is a complete listener for a parse tree produced by BASICParser.
type BaseBASICListener struct{}

var _ BASICListener = &BaseBASICListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBASICListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBASICListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBASICListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBASICListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBASICListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBASICListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBASICListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBASICListener) ExitStatement(ctx *StatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseBASICListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseBASICListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBASICListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBASICListener) ExitExpression(ctx *ExpressionContext) {}
