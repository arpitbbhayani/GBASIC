// Code generated from parser/BASIC.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BASICLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var basiclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func basiclexerLexerInit() {
	staticData := &basiclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'PRINT'",
	}
	staticData.symbolicNames = []string{
		"", "PRINT", "NUMBER", "STRING", "SPACES", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"ESC_SEQ", "DIGIT", "SPACE", "PRINT", "NUMBER", "STRING", "SPACES",
		"WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 54, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 32, 8, 4,
		11, 4, 12, 4, 33, 1, 5, 1, 5, 1, 5, 5, 5, 39, 8, 5, 10, 5, 12, 5, 42, 9,
		5, 1, 5, 1, 5, 1, 6, 4, 6, 47, 8, 6, 11, 6, 12, 6, 48, 1, 7, 1, 7, 1, 7,
		1, 7, 0, 0, 8, 1, 0, 3, 0, 5, 0, 7, 1, 9, 2, 11, 3, 13, 4, 15, 5, 1, 0,
		4, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114,
		116, 116, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32,
		32, 54, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 20, 1, 0, 0, 0, 5, 22,
		1, 0, 0, 0, 7, 24, 1, 0, 0, 0, 9, 31, 1, 0, 0, 0, 11, 35, 1, 0, 0, 0, 13,
		46, 1, 0, 0, 0, 15, 50, 1, 0, 0, 0, 17, 18, 5, 92, 0, 0, 18, 19, 7, 0,
		0, 0, 19, 2, 1, 0, 0, 0, 20, 21, 7, 1, 0, 0, 21, 4, 1, 0, 0, 0, 22, 23,
		5, 32, 0, 0, 23, 6, 1, 0, 0, 0, 24, 25, 5, 80, 0, 0, 25, 26, 5, 82, 0,
		0, 26, 27, 5, 73, 0, 0, 27, 28, 5, 78, 0, 0, 28, 29, 5, 84, 0, 0, 29, 8,
		1, 0, 0, 0, 30, 32, 3, 3, 1, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0,
		33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 10, 1, 0, 0, 0, 35, 40, 5,
		34, 0, 0, 36, 39, 3, 1, 0, 0, 37, 39, 8, 2, 0, 0, 38, 36, 1, 0, 0, 0, 38,
		37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0,
		0, 41, 43, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 34, 0, 0, 44, 12,
		1, 0, 0, 0, 45, 47, 3, 5, 2, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 14, 1, 0, 0, 0, 50, 51, 7,
		3, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 7, 0, 0, 53, 16, 1, 0, 0, 0, 5,
		0, 33, 38, 40, 48, 1, 6, 0, 0,
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

// BASICLexerInit initializes any static state used to implement BASICLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBASICLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BASICLexerInit() {
	staticData := &basiclexerLexerStaticData
	staticData.once.Do(basiclexerLexerInit)
}

// NewBASICLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBASICLexer(input antlr.CharStream) *BASICLexer {
	BASICLexerInit()
	l := new(BASICLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &basiclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "BASIC.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BASICLexer tokens.
const (
	BASICLexerPRINT      = 1
	BASICLexerNUMBER     = 2
	BASICLexerSTRING     = 3
	BASICLexerSPACES     = 4
	BASICLexerWHITESPACE = 5
)
