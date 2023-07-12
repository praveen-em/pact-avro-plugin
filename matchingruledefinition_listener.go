package main

import (
	"log"
	"math/big"
	"strconv"
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
	if ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText() == "notEmpty" {		
		l.matchType = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()
		// log.Println("Match Type: ", l.matchType)
	}
}

func (l *MatchingRuleDefinitionListener) ExitMatchingRule(ctx *parser.MatchingRuleContext) {
	
	l.matchType = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()
	// log.Println("Match Type: ", l.matchType)

	if (l.matchType == "regex" || l.matchType == "contentType" || ctx.GetMatcherType() != nil) && (ctx.GetChildCount() >= 3)  {
		l.matchTypeConfig = strings.Trim(ctx.GetChild(THIRD).(*parser.StringContext).GetText(), "'") 
		// log.Println("Match Type Config: ", l.matchTypeConfig)
	}

	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValue, ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValue, err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
		// log.Println("INTEGER: ", l.exampleValue)
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValue, err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}
}

func (l *MatchingRuleDefinitionListener) ExitString(ctx *parser.StringContext) {
	l.exampleValue = strings.Trim(ctx.GetText(), "'")
	// log.Println("STRING: ", l.exampleValue)
}

func (l *MatchingRuleDefinitionListener) ExitPrimitiveValue(ctx *parser.PrimitiveValueContext) {
	// l.exampleValue = strings.Trim(ctx.GetText(), "'")

	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValue, ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValue, err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
		// log.Println("INTEGER: ", l.exampleValue)
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValue, err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}

}
