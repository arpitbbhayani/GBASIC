// Code generated from parser/BASIC.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // BASIC

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BASICParser struct {
	*antlr.BaseParser
}

var basicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func basicParserInit() {
	staticData := &basicParserStaticData
	staticData.literalNames = []string{
		"", "'PRINT'",
	}
	staticData.symbolicNames = []string{
		"", "PRINT", "NUMBER", "STRING", "SPACES", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"program", "statement", "printStatement", "expression",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 27, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5, 0,
		10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 1, 1, 0,
		2, 3, 23, 0, 11, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 18, 1, 0, 0, 0, 6, 24,
		1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11,
		9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13, 11, 1, 0, 0,
		0, 14, 15, 5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 17, 3, 4, 2, 0, 17, 3, 1,
		0, 0, 0, 18, 19, 5, 2, 0, 0, 19, 20, 5, 4, 0, 0, 20, 21, 5, 1, 0, 0, 21,
		22, 5, 4, 0, 0, 22, 23, 3, 6, 3, 0, 23, 5, 1, 0, 0, 0, 24, 25, 7, 0, 0,
		0, 25, 7, 1, 0, 0, 0, 1, 11,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BASICParserInit initializes any static state used to implement BASICParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBASICParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BASICParserInit() {
	staticData := &basicParserStaticData
	staticData.once.Do(basicParserInit)
}

// NewBASICParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBASICParser(input antlr.TokenStream) *BASICParser {
	BASICParserInit()
	this := new(BASICParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &basicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BASIC.g4"

	return this
}

// BASICParser tokens.
const (
	BASICParserEOF        = antlr.TokenEOF
	BASICParserPRINT      = 1
	BASICParserNUMBER     = 2
	BASICParserSTRING     = 3
	BASICParserSPACES     = 4
	BASICParserWHITESPACE = 5
)

// BASICParser rules.
const (
	BASICParserRULE_program        = 0
	BASICParserRULE_statement      = 1
	BASICParserRULE_printStatement = 2
	BASICParserRULE_expression     = 3
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BASICParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BASICParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BASICParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BASICParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BASICParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BASICParserNUMBER {
		{
			p.SetState(8)
			p.Statement()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(BASICParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrintStatement() IPrintStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BASICParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BASICParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *BASICParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BASICParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.PrintStatement()
	}

	return localctx
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	AllSPACES() []antlr.TerminalNode
	SPACES(i int) antlr.TerminalNode
	PRINT() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BASICParserRULE_printStatement
	return p
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BASICParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BASICParserNUMBER, 0)
}

func (s *PrintStatementContext) AllSPACES() []antlr.TerminalNode {
	return s.GetTokens(BASICParserSPACES)
}

func (s *PrintStatementContext) SPACES(i int) antlr.TerminalNode {
	return s.GetToken(BASICParserSPACES, i)
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(BASICParserPRINT, 0)
}

func (s *PrintStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (p *BASICParser) PrintStatement() (localctx IPrintStatementContext) {
	this := p
	_ = this

	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BASICParserRULE_printStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(BASICParserNUMBER)
	}
	{
		p.SetState(19)
		p.Match(BASICParserSPACES)
	}
	{
		p.SetState(20)
		p.Match(BASICParserPRINT)
	}
	{
		p.SetState(21)
		p.Match(BASICParserSPACES)
	}
	{
		p.SetState(22)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BASICParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BASICParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BASICParserNUMBER, 0)
}

func (s *ExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(BASICParserSTRING, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BASICListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BASICParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BASICParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BASICParserNUMBER || _la == BASICParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
