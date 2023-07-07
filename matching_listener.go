package main

import (
	"log"
	"strings"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/praveen-em/pact-foobar-plugin/antlr_auto_generated_parser"
)

const (
	FIRST = 0
	THIRD = 2
)

type MatchingRuleDefinitionListener struct {
	*parser.BaseMatchingRuleDefinitionListener
	exampleValue interface{}
	matchType string
	matchTypeConfig string
}

func (l *MatchingRuleDefinitionListener) getExampleValue() interface{} {
	return l.exampleValue
}

func (l *MatchingRuleDefinitionListener) getMatchType() string {
	return l.matchType
}

func (l *MatchingRuleDefinitionListener) getMatchTypeConfig() string {
	return l.matchTypeConfig
}

func(l *MatchingRuleDefinitionListener) getParsedData() (matchType string, matchTypeConfig string, exampleValue interface{}) {
	return l.matchType, l.matchTypeConfig, l.exampleValue
}

func (l *MatchingRuleDefinitionListener) ExitMatchingDefinitionExp(ctx *parser.MatchingDefinitionExpContext) {
	if ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText() == "notEmpty" {		
		l.matchType = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		log.Println("Match Type: ", l.matchType)
	}
}

func (l *MatchingRuleDefinitionListener) ExitMatchingRule(ctx *parser.MatchingRuleContext) {
	l.matchType = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()
	log.Println("Match Type: ", l.matchType)

	if (l.matchType == "regex" || l.matchType == "contentType" || ctx.GetMatcherType() != nil) && (ctx.GetChildCount() >= 3)  {
		l.matchTypeConfig = strings.Trim(ctx.GetChild(THIRD).(*parser.StringContext).GetText(), "'") 
		log.Println("Match Type Config: ", l.matchTypeConfig)
	}

	if ctx.DECIMAL_LITERAL() != nil {
		l.exampleValue = ctx.DECIMAL_LITERAL().GetText()
	}

	if ctx.INTEGER_LITERAL() != nil {
		l.exampleValue = ctx.INTEGER_LITERAL().GetText()
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		l.exampleValue = ctx.BOOLEAN_LITERAL().GetText()
	}
}

func (l *MatchingRuleDefinitionListener) ExitString(ctx *parser.StringContext) {
	// log.Println("Exit String: ", ctx.GetText())
	l.exampleValue = strings.Trim(ctx.GetText(), "'")
}

func (l *MatchingRuleDefinitionListener) ExitPrimitiveValue(ctx *parser.PrimitiveValueContext) {
	// log.Println("Exit Primitive: ", ctx.GetText())
	l.exampleValue = strings.Trim(ctx.GetText(), "'")
}
