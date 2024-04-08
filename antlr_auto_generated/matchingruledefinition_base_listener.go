// Code generated from MatchingRuleDefinition.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MatchingRuleDefinition

import "github.com/antlr4-go/antlr/v4"

// BaseMatchingRuleDefinitionListener is a complete listener for a parse tree produced by MatchingRuleDefinitionParser.
type BaseMatchingRuleDefinitionListener struct{}

var _ MatchingRuleDefinitionListener = &BaseMatchingRuleDefinitionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMatchingRuleDefinitionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMatchingRuleDefinitionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMatchingDefinition is called when production matchingDefinition is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterMatchingDefinition(ctx *MatchingDefinitionContext) {
}

// ExitMatchingDefinition is called when production matchingDefinition is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitMatchingDefinition(ctx *MatchingDefinitionContext) {}

// EnterMatchingDefinitionExp is called when production matchingDefinitionExp is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterMatchingDefinitionExp(ctx *MatchingDefinitionExpContext) {
}

// ExitMatchingDefinitionExp is called when production matchingDefinitionExp is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitMatchingDefinitionExp(ctx *MatchingDefinitionExpContext) {
}

// EnterMatchingRule is called when production matchingRule is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterMatchingRule(ctx *MatchingRuleContext) {}

// ExitMatchingRule is called when production matchingRule is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitMatchingRule(ctx *MatchingRuleContext) {}

// EnterPrimitiveValue is called when production primitiveValue is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterPrimitiveValue(ctx *PrimitiveValueContext) {}

// ExitPrimitiveValue is called when production primitiveValue is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitPrimitiveValue(ctx *PrimitiveValueContext) {}

// EnterString is called when production string is entered.
func (s *BaseMatchingRuleDefinitionListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseMatchingRuleDefinitionListener) ExitString(ctx *StringContext) {}
