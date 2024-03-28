package main

import (
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/praveen-em/pact-avro-plugin/antlr_auto_generated_parser"
)

const (
	FIRST = 0
	THIRD = 2
)

type MatchingRuleDefinitionListener struct {
	*parser.BaseMatchingRuleDefinitionListener
	exampleValue map[string]interface{}
	matchType map[string]interface{}
	matchTypeConfig map[string]interface{}
	key string
}

func(l *MatchingRuleDefinitionListener) getParsedData() (matchType map[string]interface{}, matchTypeConfig map[string]interface{}, exampleValue map[string]interface{}) {
	return l.matchType, l.matchTypeConfig, l.exampleValue
}

func (l *MatchingRuleDefinitionListener) EnterMatchingDefinitionExp(ctx *parser.MatchingDefinitionExpContext) {
	switch ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText() {
	case "notEmpty":
		l.matchType[l.key] = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()
	case "eachKey":
		l.key = "eachKey"
	case "eachValue":
		l.key = "eachValue"
	}
}

func (l *MatchingRuleDefinitionListener) ExitMatchingRule(ctx *parser.MatchingRuleContext) {

	l.matchType[l.key] = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()

	if (l.matchType[l.key] == "regex" || l.matchType[l.key] == "contentType" || ctx.GetMatcherType() != nil) && (ctx.GetChildCount() >= 3)  {
		l.matchTypeConfig[l.key] = strings.Trim(ctx.GetChild(THIRD).(*parser.StringContext).GetText(), "'") 
	}

	// TODO: currently handles for avro bytes.decimal. needs to handle prmitive decimal. Maybe move typeConversion out to a common helper function outside of this.
	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValue[l.key], ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValue[l.key], err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValue[l.key], err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}

	if ctx.DOLLAR() != nil {

		l.key = "reference"
		l.exampleValue[l.key] = true
		l.matchType[l.key] = "values"
	}
}

func (l *MatchingRuleDefinitionListener) ExitString(ctx *parser.StringContext) {
	l.exampleValue[l.key] = strings.Trim(ctx.GetText(), "'")
}

func (l *MatchingRuleDefinitionListener) ExitPrimitiveValue(ctx *parser.PrimitiveValueContext) {
	// TODO: currently handles for avro bytes.decimal. needs to handle prmitive decimal. Maybe move typeConversion out to a common helper function outside of this.
	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValue[l.key], ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValue[l.key], err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValue[l.key], err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}

}

