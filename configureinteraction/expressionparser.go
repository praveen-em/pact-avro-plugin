package configureinteraction

import (
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/praveen-em/pact-avro-plugin/antlr_auto_generated"
)

const (
	FIRST = 0
	THIRD = 2
)

type MatchingRuleDefinitionListener struct {
	*parser.BaseMatchingRuleDefinitionListener
	exampleValueMap map[string]interface{}
	matchTypeMap map[string]interface{}
	matchTypeConfigMap map[string]interface{}
	key string
}

func parseExpression(expression string) (exampleValueMap map[string]interface{}, exampleValue interface{}, matchTypeMap map[string]interface{}, matchTypeConfigMap map[string]interface{}, err error) {		
	var (
		isMap       bool
		isReference bool
		eachKey     string
	)

	matchTypeMap, matchTypeConfigMap, exampleValueMap = parseMatchingRuleDefinition(expression)		

	if v, ok := exampleValueMap["eachKey"]; ok {
		matchType, _ := matchTypeMap["eachKey"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["eachKey"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}
		isMap = true
		eachKey = v.(string)
	}

	if _, ok := exampleValueMap["reference"]; ok {
		isReference = true
	}

	if v, ok := exampleValueMap["default"]; ok {
		matchType, _ := matchTypeMap["default"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["default"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}
		exampleValue = v
	}

	if v, ok := exampleValueMap["eachValue"]; ok {
		matchType, _ := matchTypeMap["eachValue"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["eachValue"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}

		switch {
		case isMap:
			m := make(map[string]any)
			m[eachKey] = v
			exampleValue = m
		case isReference:
			exampleValue = v
		default:
			var s []interface{}
			s = append(s, v)
			exampleValue = s
		}
	}
	return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
}

func convertToNativeType(matchType string, matchTypeConfig string, exampleValue interface{}) (interface{}, error) {
	//TODO: cover all data types
	var err error
	switch matchType {
	case "date", "datetime", "time":
		{
			date, err := time.Parse(matchTypeConfig, exampleValue.(string))
			if err != nil {
				log.Println("ERROR While converting to native date format. ", err)
				return nil, err
			} else {
				exampleValue = date
			}
		}
	}
	return exampleValue, err
}

// This function makes use of ANTLR generated lexer and parser functions.
func parseMatchingRuleDefinition(expression string) (matchType map[string]interface{}, matchTypeConfig map[string]interface{}, exampleValue map[string]interface{}) {
	is := antlr.NewInputStream(expression)
	lexer := parser.NewMatchingRuleDefinitionLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMatchingRuleDefinitionParser(stream)
	tree := p.MatchingDefinition()
	listener := new(MatchingRuleDefinitionListener)
	initListener(listener)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.getParsedData()
}

func initListener(listener *MatchingRuleDefinitionListener) {
	listener.exampleValueMap = make(map[string]interface{})
	listener.matchTypeMap = make(map[string]interface{})
	listener.matchTypeConfigMap = make(map[string]interface{})
	listener.key = "default"
}

func(l *MatchingRuleDefinitionListener) getParsedData() (matchTypeMap map[string]interface{}, matchTypeConfigMap map[string]interface{}, exampleValueMap map[string]interface{}) {
	return l.matchTypeMap, l.matchTypeConfigMap, l.exampleValueMap
}

func (l *MatchingRuleDefinitionListener) EnterMatchingDefinitionExp(ctx *parser.MatchingDefinitionExpContext) {
	switch ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText() {
	case "notEmpty":
		l.matchTypeMap[l.key] = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()
	case "eachKey":
		l.key = "eachKey"
	case "eachValue":
		l.key = "eachValue"
	}
}

func (l *MatchingRuleDefinitionListener) ExitMatchingRule(ctx *parser.MatchingRuleContext) {
	l.matchTypeMap[l.key] = ctx.GetChild(FIRST).(*antlr.TerminalNodeImpl).GetText()

	if (l.matchTypeMap[l.key] == "regex" || l.matchTypeMap[l.key] == "contentType" || ctx.GetMatcherType() != nil) && (ctx.GetChildCount() >= 3)  {
		l.matchTypeConfigMap[l.key] = strings.Trim(ctx.GetChild(THIRD).(*parser.StringContext).GetText(), "'") 
	}

	// TODO: currently handles for avro bytes.decimal. needs to handle prmitive decimal. Maybe move typeConversion out to a common helper function outside of this.
	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValueMap[l.key], ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValueMap[l.key], err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValueMap[l.key], err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}

	if ctx.DOLLAR() != nil {

		l.key = "reference"
		l.exampleValueMap[l.key] = true
		l.matchTypeMap[l.key] = "values"
	}
}

func (l *MatchingRuleDefinitionListener) ExitString(ctx *parser.StringContext) {
	l.exampleValueMap[l.key] = strings.Trim(ctx.GetText(), "'")
}

func (l *MatchingRuleDefinitionListener) ExitPrimitiveValue(ctx *parser.PrimitiveValueContext) {
	// TODO: currently handles for avro bytes.decimal. needs to handle prmitive decimal. Maybe move typeConversion out to a common helper function outside of this.
	if ctx.DECIMAL_LITERAL() != nil {
		r := new(big.Rat)
		var ok bool
		if l.exampleValueMap[l.key], ok = r.SetString(ctx.DECIMAL_LITERAL().GetText()); !ok {
			log.Println("ERROR while converting string to decimal")
		}
	}

	if ctx.INTEGER_LITERAL() != nil {
		var err error		
		l.exampleValueMap[l.key], err = strconv.Atoi(ctx.INTEGER_LITERAL().GetText())
		if (err != nil) {
			log.Println("ERROR while converting string to integer. ", err)
		}
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		var err error
		l.exampleValueMap[l.key], err = strconv.ParseBool(ctx.BOOLEAN_LITERAL().GetText()) 
		if (err != nil) {
			log.Println("ERROR while converting string to boolean. ", err)
		}
	}
}