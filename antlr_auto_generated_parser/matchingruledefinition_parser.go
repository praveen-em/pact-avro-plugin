// Code generated from MatchingRuleDefinition.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MatchingRuleDefinition

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MatchingRuleDefinitionParser struct {
	*antlr.BaseParser
}

var MatchingRuleDefinitionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchingruledefinitionParserInit() {
	staticData := &MatchingRuleDefinitionParserStaticData
	staticData.LiteralNames = []string{
		"", "'matching'", "'notEmpty'", "'eachKey'", "'eachValue'", "'equalTo'",
		"'type'", "'number'", "'integer'", "'decimal'", "'datetime'", "'date'",
		"'time'", "'regex'", "'include'", "'boolean'", "'semver'", "'contentType'",
		"'null'", "", "", "'('", "')'", "", "", "','", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "INTEGER_LITERAL", "DECIMAL_LITERAL", "LEFT_BRACKET", "RIGHT_BRACKET",
		"STRING_LITERAL", "BOOLEAN_LITERAL", "COMMA", "DOLLAR", "WS",
	}
	staticData.RuleNames = []string{
		"matchingDefinition", "matchingDefinitionExp", "matchingRule", "primitiveValue",
		"string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 94, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 41, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 84, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 90,
		8, 3, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 4, 1, 0, 5, 6, 1, 0,
		19, 20, 1, 0, 10, 12, 2, 0, 18, 18, 23, 23, 105, 0, 10, 1, 0, 0, 0, 2,
		40, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 89, 1, 0, 0, 0, 8, 91, 1, 0, 0, 0,
		10, 15, 3, 2, 1, 0, 11, 12, 5, 25, 0, 0, 12, 14, 3, 2, 1, 0, 13, 11, 1,
		0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16,
		18, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 0, 0, 1, 19, 1, 1, 0, 0,
		0, 20, 21, 5, 1, 0, 0, 21, 22, 5, 21, 0, 0, 22, 23, 3, 4, 2, 0, 23, 24,
		5, 22, 0, 0, 24, 41, 1, 0, 0, 0, 25, 26, 5, 2, 0, 0, 26, 27, 5, 21, 0,
		0, 27, 28, 3, 6, 3, 0, 28, 29, 5, 22, 0, 0, 29, 41, 1, 0, 0, 0, 30, 31,
		5, 3, 0, 0, 31, 32, 5, 21, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 22, 0,
		0, 34, 41, 1, 0, 0, 0, 35, 36, 5, 4, 0, 0, 36, 37, 5, 21, 0, 0, 37, 38,
		3, 2, 1, 0, 38, 39, 5, 22, 0, 0, 39, 41, 1, 0, 0, 0, 40, 20, 1, 0, 0, 0,
		40, 25, 1, 0, 0, 0, 40, 30, 1, 0, 0, 0, 40, 35, 1, 0, 0, 0, 41, 3, 1, 0,
		0, 0, 42, 43, 7, 0, 0, 0, 43, 44, 5, 25, 0, 0, 44, 84, 3, 6, 3, 0, 45,
		46, 5, 7, 0, 0, 46, 47, 5, 25, 0, 0, 47, 84, 7, 1, 0, 0, 48, 49, 5, 8,
		0, 0, 49, 50, 5, 25, 0, 0, 50, 84, 5, 19, 0, 0, 51, 52, 5, 9, 0, 0, 52,
		53, 5, 25, 0, 0, 53, 84, 5, 20, 0, 0, 54, 55, 7, 2, 0, 0, 55, 56, 5, 25,
		0, 0, 56, 57, 3, 8, 4, 0, 57, 58, 5, 25, 0, 0, 58, 59, 3, 8, 4, 0, 59,
		84, 1, 0, 0, 0, 60, 61, 5, 13, 0, 0, 61, 62, 5, 25, 0, 0, 62, 63, 3, 8,
		4, 0, 63, 64, 5, 25, 0, 0, 64, 65, 3, 8, 4, 0, 65, 84, 1, 0, 0, 0, 66,
		67, 5, 14, 0, 0, 67, 68, 5, 25, 0, 0, 68, 84, 3, 8, 4, 0, 69, 70, 5, 15,
		0, 0, 70, 71, 5, 25, 0, 0, 71, 84, 5, 24, 0, 0, 72, 73, 5, 16, 0, 0, 73,
		74, 5, 25, 0, 0, 74, 84, 3, 8, 4, 0, 75, 76, 5, 17, 0, 0, 76, 77, 5, 25,
		0, 0, 77, 78, 3, 8, 4, 0, 78, 79, 5, 25, 0, 0, 79, 80, 3, 8, 4, 0, 80,
		84, 1, 0, 0, 0, 81, 82, 5, 26, 0, 0, 82, 84, 3, 8, 4, 0, 83, 42, 1, 0,
		0, 0, 83, 45, 1, 0, 0, 0, 83, 48, 1, 0, 0, 0, 83, 51, 1, 0, 0, 0, 83, 54,
		1, 0, 0, 0, 83, 60, 1, 0, 0, 0, 83, 66, 1, 0, 0, 0, 83, 69, 1, 0, 0, 0,
		83, 72, 1, 0, 0, 0, 83, 75, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 5, 1, 0,
		0, 0, 85, 90, 3, 8, 4, 0, 86, 90, 5, 20, 0, 0, 87, 90, 5, 19, 0, 0, 88,
		90, 5, 24, 0, 0, 89, 85, 1, 0, 0, 0, 89, 86, 1, 0, 0, 0, 89, 87, 1, 0,
		0, 0, 89, 88, 1, 0, 0, 0, 90, 7, 1, 0, 0, 0, 91, 92, 7, 3, 0, 0, 92, 9,
		1, 0, 0, 0, 4, 15, 40, 83, 89,
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

// MatchingRuleDefinitionParserInit initializes any static state used to implement MatchingRuleDefinitionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMatchingRuleDefinitionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchingRuleDefinitionParserInit() {
	staticData := &MatchingRuleDefinitionParserStaticData
	staticData.once.Do(matchingruledefinitionParserInit)
}

// NewMatchingRuleDefinitionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMatchingRuleDefinitionParser(input antlr.TokenStream) *MatchingRuleDefinitionParser {
	MatchingRuleDefinitionParserInit()
	this := new(MatchingRuleDefinitionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MatchingRuleDefinitionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MatchingRuleDefinition.g4"

	return this
}

// MatchingRuleDefinitionParser tokens.
const (
	MatchingRuleDefinitionParserEOF             = antlr.TokenEOF
	MatchingRuleDefinitionParserT__0            = 1
	MatchingRuleDefinitionParserT__1            = 2
	MatchingRuleDefinitionParserT__2            = 3
	MatchingRuleDefinitionParserT__3            = 4
	MatchingRuleDefinitionParserT__4            = 5
	MatchingRuleDefinitionParserT__5            = 6
	MatchingRuleDefinitionParserT__6            = 7
	MatchingRuleDefinitionParserT__7            = 8
	MatchingRuleDefinitionParserT__8            = 9
	MatchingRuleDefinitionParserT__9            = 10
	MatchingRuleDefinitionParserT__10           = 11
	MatchingRuleDefinitionParserT__11           = 12
	MatchingRuleDefinitionParserT__12           = 13
	MatchingRuleDefinitionParserT__13           = 14
	MatchingRuleDefinitionParserT__14           = 15
	MatchingRuleDefinitionParserT__15           = 16
	MatchingRuleDefinitionParserT__16           = 17
	MatchingRuleDefinitionParserT__17           = 18
	MatchingRuleDefinitionParserINTEGER_LITERAL = 19
	MatchingRuleDefinitionParserDECIMAL_LITERAL = 20
	MatchingRuleDefinitionParserLEFT_BRACKET    = 21
	MatchingRuleDefinitionParserRIGHT_BRACKET   = 22
	MatchingRuleDefinitionParserSTRING_LITERAL  = 23
	MatchingRuleDefinitionParserBOOLEAN_LITERAL = 24
	MatchingRuleDefinitionParserCOMMA           = 25
	MatchingRuleDefinitionParserDOLLAR          = 26
	MatchingRuleDefinitionParserWS              = 27
)

// MatchingRuleDefinitionParser rules.
const (
	MatchingRuleDefinitionParserRULE_matchingDefinition    = 0
	MatchingRuleDefinitionParserRULE_matchingDefinitionExp = 1
	MatchingRuleDefinitionParserRULE_matchingRule          = 2
	MatchingRuleDefinitionParserRULE_primitiveValue        = 3
	MatchingRuleDefinitionParserRULE_string                = 4
)

// IMatchingDefinitionContext is an interface to support dynamic dispatch.
type IMatchingDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMatchingDefinitionExp() []IMatchingDefinitionExpContext
	MatchingDefinitionExp(i int) IMatchingDefinitionExpContext
	EOF() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMatchingDefinitionContext differentiates from other interfaces.
	IsMatchingDefinitionContext()
}

type MatchingDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchingDefinitionContext() *MatchingDefinitionContext {
	var p = new(MatchingDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinition
	return p
}

func InitEmptyMatchingDefinitionContext(p *MatchingDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinition
}

func (*MatchingDefinitionContext) IsMatchingDefinitionContext() {}

func NewMatchingDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchingDefinitionContext {
	var p = new(MatchingDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinition

	return p
}

func (s *MatchingDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchingDefinitionContext) AllMatchingDefinitionExp() []IMatchingDefinitionExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchingDefinitionExpContext); ok {
			len++
		}
	}

	tst := make([]IMatchingDefinitionExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchingDefinitionExpContext); ok {
			tst[i] = t.(IMatchingDefinitionExpContext)
			i++
		}
	}

	return tst
}

func (s *MatchingDefinitionContext) MatchingDefinitionExp(i int) IMatchingDefinitionExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchingDefinitionExpContext); ok {
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

	return t.(IMatchingDefinitionExpContext)
}

func (s *MatchingDefinitionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserEOF, 0)
}

func (s *MatchingDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MatchingRuleDefinitionParserCOMMA)
}

func (s *MatchingDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserCOMMA, i)
}

func (s *MatchingDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchingDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchingDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.EnterMatchingDefinition(s)
	}
}

func (s *MatchingDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.ExitMatchingDefinition(s)
	}
}

func (p *MatchingRuleDefinitionParser) MatchingDefinition() (localctx IMatchingDefinitionContext) {
	localctx = NewMatchingDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MatchingRuleDefinitionParserRULE_matchingDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.MatchingDefinitionExp()
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MatchingRuleDefinitionParserCOMMA {
		{
			p.SetState(11)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)
			p.MatchingDefinitionExp()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(MatchingRuleDefinitionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchingDefinitionExpContext is an interface to support dynamic dispatch.
type IMatchingDefinitionExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	MatchingRule() IMatchingRuleContext
	RIGHT_BRACKET() antlr.TerminalNode
	PrimitiveValue() IPrimitiveValueContext
	MatchingDefinitionExp() IMatchingDefinitionExpContext

	// IsMatchingDefinitionExpContext differentiates from other interfaces.
	IsMatchingDefinitionExpContext()
}

type MatchingDefinitionExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchingDefinitionExpContext() *MatchingDefinitionExpContext {
	var p = new(MatchingDefinitionExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinitionExp
	return p
}

func InitEmptyMatchingDefinitionExpContext(p *MatchingDefinitionExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinitionExp
}

func (*MatchingDefinitionExpContext) IsMatchingDefinitionExpContext() {}

func NewMatchingDefinitionExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchingDefinitionExpContext {
	var p = new(MatchingDefinitionExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingDefinitionExp

	return p
}

func (s *MatchingDefinitionExpContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchingDefinitionExpContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserLEFT_BRACKET, 0)
}

func (s *MatchingDefinitionExpContext) MatchingRule() IMatchingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchingRuleContext)
}

func (s *MatchingDefinitionExpContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserRIGHT_BRACKET, 0)
}

func (s *MatchingDefinitionExpContext) PrimitiveValue() IPrimitiveValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
}

func (s *MatchingDefinitionExpContext) MatchingDefinitionExp() IMatchingDefinitionExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchingDefinitionExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchingDefinitionExpContext)
}

func (s *MatchingDefinitionExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchingDefinitionExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchingDefinitionExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.EnterMatchingDefinitionExp(s)
	}
}

func (s *MatchingDefinitionExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.ExitMatchingDefinitionExp(s)
	}
}

func (p *MatchingRuleDefinitionParser) MatchingDefinitionExp() (localctx IMatchingDefinitionExpContext) {
	localctx = NewMatchingDefinitionExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MatchingRuleDefinitionParserRULE_matchingDefinitionExp)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MatchingRuleDefinitionParserT__0:
		{
			p.SetState(20)
			p.Match(MatchingRuleDefinitionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.Match(MatchingRuleDefinitionParserLEFT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.MatchingRule()
		}
		{
			p.SetState(23)
			p.Match(MatchingRuleDefinitionParserRIGHT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__1:
		{
			p.SetState(25)
			p.Match(MatchingRuleDefinitionParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(MatchingRuleDefinitionParserLEFT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.PrimitiveValue()
		}
		{
			p.SetState(28)
			p.Match(MatchingRuleDefinitionParserRIGHT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__2:
		{
			p.SetState(30)
			p.Match(MatchingRuleDefinitionParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(MatchingRuleDefinitionParserLEFT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.MatchingDefinitionExp()
		}
		{
			p.SetState(33)
			p.Match(MatchingRuleDefinitionParserRIGHT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__3:
		{
			p.SetState(35)
			p.Match(MatchingRuleDefinitionParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(MatchingRuleDefinitionParserLEFT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.MatchingDefinitionExp()
		}
		{
			p.SetState(38)
			p.Match(MatchingRuleDefinitionParserRIGHT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchingRuleContext is an interface to support dynamic dispatch.
type IMatchingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMatcherType returns the matcherType token.
	GetMatcherType() antlr.Token

	// SetMatcherType sets the matcherType token.
	SetMatcherType(antlr.Token)

	// Getter signatures
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PrimitiveValue() IPrimitiveValueContext
	DECIMAL_LITERAL() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	AllString_() []IStringContext
	String_(i int) IStringContext
	BOOLEAN_LITERAL() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode

	// IsMatchingRuleContext differentiates from other interfaces.
	IsMatchingRuleContext()
}

type MatchingRuleContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	matcherType antlr.Token
}

func NewEmptyMatchingRuleContext() *MatchingRuleContext {
	var p = new(MatchingRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingRule
	return p
}

func InitEmptyMatchingRuleContext(p *MatchingRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingRule
}

func (*MatchingRuleContext) IsMatchingRuleContext() {}

func NewMatchingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchingRuleContext {
	var p = new(MatchingRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchingRuleDefinitionParserRULE_matchingRule

	return p
}

func (s *MatchingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchingRuleContext) GetMatcherType() antlr.Token { return s.matcherType }

func (s *MatchingRuleContext) SetMatcherType(v antlr.Token) { s.matcherType = v }

func (s *MatchingRuleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MatchingRuleDefinitionParserCOMMA)
}

func (s *MatchingRuleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserCOMMA, i)
}

func (s *MatchingRuleContext) PrimitiveValue() IPrimitiveValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
}

func (s *MatchingRuleContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserDECIMAL_LITERAL, 0)
}

func (s *MatchingRuleContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserINTEGER_LITERAL, 0)
}

func (s *MatchingRuleContext) AllString_() []IStringContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringContext); ok {
			len++
		}
	}

	tst := make([]IStringContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringContext); ok {
			tst[i] = t.(IStringContext)
			i++
		}
	}

	return tst
}

func (s *MatchingRuleContext) String_(i int) IStringContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
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

	return t.(IStringContext)
}

func (s *MatchingRuleContext) BOOLEAN_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserBOOLEAN_LITERAL, 0)
}

func (s *MatchingRuleContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserDOLLAR, 0)
}

func (s *MatchingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchingRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.EnterMatchingRule(s)
	}
}

func (s *MatchingRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.ExitMatchingRule(s)
	}
}

func (p *MatchingRuleDefinitionParser) MatchingRule() (localctx IMatchingRuleContext) {
	localctx = NewMatchingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MatchingRuleDefinitionParserRULE_matchingRule)
	var _la int

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MatchingRuleDefinitionParserT__4, MatchingRuleDefinitionParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MatchingRuleDefinitionParserT__4 || _la == MatchingRuleDefinitionParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(43)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.PrimitiveValue()
		}

	case MatchingRuleDefinitionParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(MatchingRuleDefinitionParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MatchingRuleDefinitionParserINTEGER_LITERAL || _la == MatchingRuleDefinitionParserDECIMAL_LITERAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MatchingRuleDefinitionParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(MatchingRuleDefinitionParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(MatchingRuleDefinitionParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Match(MatchingRuleDefinitionParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(MatchingRuleDefinitionParserDECIMAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__9, MatchingRuleDefinitionParserT__10, MatchingRuleDefinitionParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MatchingRuleContext).matcherType = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MatchingRuleContext).matcherType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(55)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.String_()
		}
		{
			p.SetState(57)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.String_()
		}

	case MatchingRuleDefinitionParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Match(MatchingRuleDefinitionParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.String_()
		}
		{
			p.SetState(63)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.String_()
		}

	case MatchingRuleDefinitionParserT__13:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(66)
			p.Match(MatchingRuleDefinitionParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.String_()
		}

	case MatchingRuleDefinitionParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(69)
			p.Match(MatchingRuleDefinitionParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(MatchingRuleDefinitionParserBOOLEAN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserT__15:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(72)
			p.Match(MatchingRuleDefinitionParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.String_()
		}

	case MatchingRuleDefinitionParserT__16:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(75)
			p.Match(MatchingRuleDefinitionParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.String_()
		}
		{
			p.SetState(78)
			p.Match(MatchingRuleDefinitionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.String_()
		}

	case MatchingRuleDefinitionParserDOLLAR:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(81)
			p.Match(MatchingRuleDefinitionParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.String_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveValueContext is an interface to support dynamic dispatch.
type IPrimitiveValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext
	DECIMAL_LITERAL() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	BOOLEAN_LITERAL() antlr.TerminalNode

	// IsPrimitiveValueContext differentiates from other interfaces.
	IsPrimitiveValueContext()
}

type PrimitiveValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveValueContext() *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_primitiveValue
	return p
}

func InitEmptyPrimitiveValueContext(p *PrimitiveValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_primitiveValue
}

func (*PrimitiveValueContext) IsPrimitiveValueContext() {}

func NewPrimitiveValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchingRuleDefinitionParserRULE_primitiveValue

	return p
}

func (s *PrimitiveValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveValueContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *PrimitiveValueContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserDECIMAL_LITERAL, 0)
}

func (s *PrimitiveValueContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserINTEGER_LITERAL, 0)
}

func (s *PrimitiveValueContext) BOOLEAN_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserBOOLEAN_LITERAL, 0)
}

func (s *PrimitiveValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.EnterPrimitiveValue(s)
	}
}

func (s *PrimitiveValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.ExitPrimitiveValue(s)
	}
}

func (p *MatchingRuleDefinitionParser) PrimitiveValue() (localctx IPrimitiveValueContext) {
	localctx = NewPrimitiveValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MatchingRuleDefinitionParserRULE_primitiveValue)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MatchingRuleDefinitionParserT__17, MatchingRuleDefinitionParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.String_()
		}

	case MatchingRuleDefinitionParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(MatchingRuleDefinitionParserDECIMAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(MatchingRuleDefinitionParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MatchingRuleDefinitionParserBOOLEAN_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.Match(MatchingRuleDefinitionParserBOOLEAN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchingRuleDefinitionParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchingRuleDefinitionParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MatchingRuleDefinitionParserSTRING_LITERAL, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchingRuleDefinitionListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *MatchingRuleDefinitionParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MatchingRuleDefinitionParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MatchingRuleDefinitionParserT__17 || _la == MatchingRuleDefinitionParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
