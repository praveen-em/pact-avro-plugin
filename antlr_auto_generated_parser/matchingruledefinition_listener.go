// Code generated from MatchingRuleDefinition.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MatchingRuleDefinition

import "github.com/antlr4-go/antlr/v4"

// MatchingRuleDefinitionListener is a complete listener for a parse tree produced by MatchingRuleDefinitionParser.
type MatchingRuleDefinitionListener interface {
	antlr.ParseTreeListener

	// EnterMatchingDefinition is called when entering the matchingDefinition production.
	EnterMatchingDefinition(c *MatchingDefinitionContext)

	// EnterMatchingDefinitionExp is called when entering the matchingDefinitionExp production.
	EnterMatchingDefinitionExp(c *MatchingDefinitionExpContext)

	// EnterMatchingRule is called when entering the matchingRule production.
	EnterMatchingRule(c *MatchingRuleContext)

	// EnterPrimitiveValue is called when entering the primitiveValue production.
	EnterPrimitiveValue(c *PrimitiveValueContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitMatchingDefinition is called when exiting the matchingDefinition production.
	ExitMatchingDefinition(c *MatchingDefinitionContext)

	// ExitMatchingDefinitionExp is called when exiting the matchingDefinitionExp production.
	ExitMatchingDefinitionExp(c *MatchingDefinitionExpContext)

	// ExitMatchingRule is called when exiting the matchingRule production.
	ExitMatchingRule(c *MatchingRuleContext)

	// ExitPrimitiveValue is called when exiting the primitiveValue production.
	ExitPrimitiveValue(c *PrimitiveValueContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}
