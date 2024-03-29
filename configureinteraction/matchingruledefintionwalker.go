package configureinteraction

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/praveen-em/pact-avro-plugin/antlr_auto_generated_parser"
)

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
	listener.exampleValue = make(map[string]interface{})
	listener.matchType = make(map[string]interface{})
	listener.matchTypeConfig = make(map[string]interface{})
	listener.key = "default"
}