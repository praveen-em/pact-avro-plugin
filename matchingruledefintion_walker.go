package main

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/praveen-em/pact-avro-plugin/antlr_auto_generated_parser"
)

func parseMatchingRuleDefinition(expression string) (matchType string, matchTypeConfig string, exampleValue interface{}) {
	is := antlr.NewInputStream(expression)
	lexer := parser.NewMatchingRuleDefinitionLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMatchingRuleDefinitionParser(stream)
	tree := p.MatchingDefinition()
	var listener MatchingRuleDefinitionListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	return listener.getParsedData()
}