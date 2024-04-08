// Code generated from MatchingRuleDefinition.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MatchingRuleDefinitionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MatchingRuleDefinitionLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchingruledefinitionlexerLexerInit() {
	staticData := &MatchingRuleDefinitionLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "INTEGER_LITERAL", "DECIMAL_LITERAL", "DIGIT", "LEFT_BRACKET",
		"RIGHT_BRACKET", "STRING_LITERAL", "BOOLEAN_LITERAL", "COMMA", "DOLLAR",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 253, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 3, 18, 196, 8, 18, 1, 18, 4,
		18, 199, 8, 18, 11, 18, 12, 18, 200, 1, 19, 3, 19, 204, 8, 19, 1, 19, 4,
		19, 207, 8, 19, 11, 19, 12, 19, 208, 1, 19, 1, 19, 4, 19, 213, 8, 19, 11,
		19, 12, 19, 214, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		5, 23, 225, 8, 23, 10, 23, 12, 23, 228, 9, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 241, 8, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 4, 27, 248, 8, 27, 11, 27, 12, 27, 249,
		1, 27, 1, 27, 0, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 0, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25,
		53, 26, 55, 27, 1, 0, 3, 1, 0, 48, 57, 1, 0, 39, 39, 3, 0, 9, 10, 13, 13,
		32, 32, 259, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 66, 1, 0, 0, 0, 5, 75, 1, 0, 0, 0, 7,
		83, 1, 0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 101, 1, 0, 0, 0, 13, 106, 1, 0,
		0, 0, 15, 113, 1, 0, 0, 0, 17, 121, 1, 0, 0, 0, 19, 129, 1, 0, 0, 0, 21,
		138, 1, 0, 0, 0, 23, 143, 1, 0, 0, 0, 25, 148, 1, 0, 0, 0, 27, 154, 1,
		0, 0, 0, 29, 162, 1, 0, 0, 0, 31, 170, 1, 0, 0, 0, 33, 177, 1, 0, 0, 0,
		35, 189, 1, 0, 0, 0, 37, 195, 1, 0, 0, 0, 39, 203, 1, 0, 0, 0, 41, 216,
		1, 0, 0, 0, 43, 218, 1, 0, 0, 0, 45, 220, 1, 0, 0, 0, 47, 222, 1, 0, 0,
		0, 49, 240, 1, 0, 0, 0, 51, 242, 1, 0, 0, 0, 53, 244, 1, 0, 0, 0, 55, 247,
		1, 0, 0, 0, 57, 58, 5, 109, 0, 0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 116,
		0, 0, 60, 61, 5, 99, 0, 0, 61, 62, 5, 104, 0, 0, 62, 63, 5, 105, 0, 0,
		63, 64, 5, 110, 0, 0, 64, 65, 5, 103, 0, 0, 65, 2, 1, 0, 0, 0, 66, 67,
		5, 110, 0, 0, 67, 68, 5, 111, 0, 0, 68, 69, 5, 116, 0, 0, 69, 70, 5, 69,
		0, 0, 70, 71, 5, 109, 0, 0, 71, 72, 5, 112, 0, 0, 72, 73, 5, 116, 0, 0,
		73, 74, 5, 121, 0, 0, 74, 4, 1, 0, 0, 0, 75, 76, 5, 101, 0, 0, 76, 77,
		5, 97, 0, 0, 77, 78, 5, 99, 0, 0, 78, 79, 5, 104, 0, 0, 79, 80, 5, 75,
		0, 0, 80, 81, 5, 101, 0, 0, 81, 82, 5, 121, 0, 0, 82, 6, 1, 0, 0, 0, 83,
		84, 5, 101, 0, 0, 84, 85, 5, 97, 0, 0, 85, 86, 5, 99, 0, 0, 86, 87, 5,
		104, 0, 0, 87, 88, 5, 86, 0, 0, 88, 89, 5, 97, 0, 0, 89, 90, 5, 108, 0,
		0, 90, 91, 5, 117, 0, 0, 91, 92, 5, 101, 0, 0, 92, 8, 1, 0, 0, 0, 93, 94,
		5, 101, 0, 0, 94, 95, 5, 113, 0, 0, 95, 96, 5, 117, 0, 0, 96, 97, 5, 97,
		0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 84, 0, 0, 99, 100, 5, 111, 0, 0,
		100, 10, 1, 0, 0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 121, 0, 0, 103,
		104, 5, 112, 0, 0, 104, 105, 5, 101, 0, 0, 105, 12, 1, 0, 0, 0, 106, 107,
		5, 110, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 109, 0, 0, 109, 110,
		5, 98, 0, 0, 110, 111, 5, 101, 0, 0, 111, 112, 5, 114, 0, 0, 112, 14, 1,
		0, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5, 110, 0, 0, 115, 116, 5, 116,
		0, 0, 116, 117, 5, 101, 0, 0, 117, 118, 5, 103, 0, 0, 118, 119, 5, 101,
		0, 0, 119, 120, 5, 114, 0, 0, 120, 16, 1, 0, 0, 0, 121, 122, 5, 100, 0,
		0, 122, 123, 5, 101, 0, 0, 123, 124, 5, 99, 0, 0, 124, 125, 5, 105, 0,
		0, 125, 126, 5, 109, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 108, 0,
		0, 128, 18, 1, 0, 0, 0, 129, 130, 5, 100, 0, 0, 130, 131, 5, 97, 0, 0,
		131, 132, 5, 116, 0, 0, 132, 133, 5, 101, 0, 0, 133, 134, 5, 116, 0, 0,
		134, 135, 5, 105, 0, 0, 135, 136, 5, 109, 0, 0, 136, 137, 5, 101, 0, 0,
		137, 20, 1, 0, 0, 0, 138, 139, 5, 100, 0, 0, 139, 140, 5, 97, 0, 0, 140,
		141, 5, 116, 0, 0, 141, 142, 5, 101, 0, 0, 142, 22, 1, 0, 0, 0, 143, 144,
		5, 116, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 109, 0, 0, 146, 147,
		5, 101, 0, 0, 147, 24, 1, 0, 0, 0, 148, 149, 5, 114, 0, 0, 149, 150, 5,
		101, 0, 0, 150, 151, 5, 103, 0, 0, 151, 152, 5, 101, 0, 0, 152, 153, 5,
		120, 0, 0, 153, 26, 1, 0, 0, 0, 154, 155, 5, 105, 0, 0, 155, 156, 5, 110,
		0, 0, 156, 157, 5, 99, 0, 0, 157, 158, 5, 108, 0, 0, 158, 159, 5, 117,
		0, 0, 159, 160, 5, 100, 0, 0, 160, 161, 5, 101, 0, 0, 161, 28, 1, 0, 0,
		0, 162, 163, 5, 98, 0, 0, 163, 164, 5, 111, 0, 0, 164, 165, 5, 111, 0,
		0, 165, 166, 5, 108, 0, 0, 166, 167, 5, 101, 0, 0, 167, 168, 5, 97, 0,
		0, 168, 169, 5, 110, 0, 0, 169, 30, 1, 0, 0, 0, 170, 171, 5, 115, 0, 0,
		171, 172, 5, 101, 0, 0, 172, 173, 5, 109, 0, 0, 173, 174, 5, 118, 0, 0,
		174, 175, 5, 101, 0, 0, 175, 176, 5, 114, 0, 0, 176, 32, 1, 0, 0, 0, 177,
		178, 5, 99, 0, 0, 178, 179, 5, 111, 0, 0, 179, 180, 5, 110, 0, 0, 180,
		181, 5, 116, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 110, 0, 0, 183,
		184, 5, 116, 0, 0, 184, 185, 5, 84, 0, 0, 185, 186, 5, 121, 0, 0, 186,
		187, 5, 112, 0, 0, 187, 188, 5, 101, 0, 0, 188, 34, 1, 0, 0, 0, 189, 190,
		5, 110, 0, 0, 190, 191, 5, 117, 0, 0, 191, 192, 5, 108, 0, 0, 192, 193,
		5, 108, 0, 0, 193, 36, 1, 0, 0, 0, 194, 196, 5, 45, 0, 0, 195, 194, 1,
		0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 1, 0, 0, 0, 197, 199, 3, 41, 20,
		0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200,
		201, 1, 0, 0, 0, 201, 38, 1, 0, 0, 0, 202, 204, 5, 45, 0, 0, 203, 202,
		1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 207, 3, 41,
		20, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0,
		208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 5, 46, 0, 0, 211,
		213, 3, 41, 20, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 212,
		1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 40, 1, 0, 0, 0, 216, 217, 7, 0,
		0, 0, 217, 42, 1, 0, 0, 0, 218, 219, 5, 40, 0, 0, 219, 44, 1, 0, 0, 0,
		220, 221, 5, 41, 0, 0, 221, 46, 1, 0, 0, 0, 222, 226, 5, 39, 0, 0, 223,
		225, 8, 1, 0, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 226, 1, 0,
		0, 0, 229, 230, 5, 39, 0, 0, 230, 48, 1, 0, 0, 0, 231, 232, 5, 116, 0,
		0, 232, 233, 5, 114, 0, 0, 233, 234, 5, 117, 0, 0, 234, 241, 5, 101, 0,
		0, 235, 236, 5, 102, 0, 0, 236, 237, 5, 97, 0, 0, 237, 238, 5, 108, 0,
		0, 238, 239, 5, 115, 0, 0, 239, 241, 5, 101, 0, 0, 240, 231, 1, 0, 0, 0,
		240, 235, 1, 0, 0, 0, 241, 50, 1, 0, 0, 0, 242, 243, 5, 44, 0, 0, 243,
		52, 1, 0, 0, 0, 244, 245, 5, 36, 0, 0, 245, 54, 1, 0, 0, 0, 246, 248, 7,
		2, 0, 0, 247, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 247, 1, 0, 0,
		0, 249, 250, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 6, 27, 0, 0, 252,
		56, 1, 0, 0, 0, 9, 0, 195, 200, 203, 208, 214, 226, 240, 249, 1, 6, 0,
		0,
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

// MatchingRuleDefinitionLexerInit initializes any static state used to implement MatchingRuleDefinitionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMatchingRuleDefinitionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchingRuleDefinitionLexerInit() {
	staticData := &MatchingRuleDefinitionLexerLexerStaticData
	staticData.once.Do(matchingruledefinitionlexerLexerInit)
}

// NewMatchingRuleDefinitionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMatchingRuleDefinitionLexer(input antlr.CharStream) *MatchingRuleDefinitionLexer {
	MatchingRuleDefinitionLexerInit()
	l := new(MatchingRuleDefinitionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MatchingRuleDefinitionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MatchingRuleDefinition.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MatchingRuleDefinitionLexer tokens.
const (
	MatchingRuleDefinitionLexerT__0            = 1
	MatchingRuleDefinitionLexerT__1            = 2
	MatchingRuleDefinitionLexerT__2            = 3
	MatchingRuleDefinitionLexerT__3            = 4
	MatchingRuleDefinitionLexerT__4            = 5
	MatchingRuleDefinitionLexerT__5            = 6
	MatchingRuleDefinitionLexerT__6            = 7
	MatchingRuleDefinitionLexerT__7            = 8
	MatchingRuleDefinitionLexerT__8            = 9
	MatchingRuleDefinitionLexerT__9            = 10
	MatchingRuleDefinitionLexerT__10           = 11
	MatchingRuleDefinitionLexerT__11           = 12
	MatchingRuleDefinitionLexerT__12           = 13
	MatchingRuleDefinitionLexerT__13           = 14
	MatchingRuleDefinitionLexerT__14           = 15
	MatchingRuleDefinitionLexerT__15           = 16
	MatchingRuleDefinitionLexerT__16           = 17
	MatchingRuleDefinitionLexerT__17           = 18
	MatchingRuleDefinitionLexerINTEGER_LITERAL = 19
	MatchingRuleDefinitionLexerDECIMAL_LITERAL = 20
	MatchingRuleDefinitionLexerLEFT_BRACKET    = 21
	MatchingRuleDefinitionLexerRIGHT_BRACKET   = 22
	MatchingRuleDefinitionLexerSTRING_LITERAL  = 23
	MatchingRuleDefinitionLexerBOOLEAN_LITERAL = 24
	MatchingRuleDefinitionLexerCOMMA           = 25
	MatchingRuleDefinitionLexerDOLLAR          = 26
	MatchingRuleDefinitionLexerWS              = 27
)
