package configureinteraction

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)



func TestMatchingRulesBuilder(t *testing.T) {

	type in struct {
		exampleValue map[string]any
		matchType map[string]any
		matchTypeConfig map[string]any
	}

	type out struct {
		rules string
		err error
	}
	
	tests := []struct {
		name string
		input in
		want out
	}{
		{
			name: "RulesType",
			input: in{exampleValue: map[string]any{"default":"value1"}, matchType: map[string]any{"default": "type"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`[{"type":"type", "values":{"match":"type"}}]` , err: nil},
			
		},
		{
			name: "RulesEqualTo",
			input: in{exampleValue: map[string]any{"default":"value1"}, matchType: map[string]any{"default": "equalTo"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`[{"type":"equalTo", "values":{"match":"equalTo"}}]` , err: nil},			
		},
		{
			name: "RulesDate",
			input: in{exampleValue: map[string]any{"default":"2023-Jul-08"}, matchType: map[string]any{"default": "date"}, matchTypeConfig: map[string]any{"default": "2006-01-02"} },
			want: out{rules:`[{"type":"date","values":{"format":"2006-01-02","match":"date"}}]` , err: nil},			
		},
		{
			name: "RulesDateTime",
			input: in{exampleValue: map[string]any{"default":"2023-07-21 16:44:32"}, matchType: map[string]any{"default": "datetime"}, matchTypeConfig: map[string]any{"default": "2006-01-02 15:04:05"} },
			want: out{rules:`[{"type":"datetime","values":{"format":"2006-01-02 15:04:05","match":"datetime"}}]` , err: nil},			
		},
		{
			name: "RulesTime",
			input: in{exampleValue: map[string]any{"default":"16:44:32"}, matchType: map[string]any{"default": "time"}, matchTypeConfig: map[string]any{"default": "15:04:05"} },
			want: out{rules:`[{"type":"time","values":{"format":"15:04:05","match":"time"}}]` , err: nil},			
		},
		{
			name: "RulesRegex",
			input: in{exampleValue: map[string]any{"default":"v2"}, matchType: map[string]any{"default": "regex"}, matchTypeConfig: map[string]any{"default": "v[a-zA-z0-9]+"} },
			want: out{rules:`[{"type":"regex","values":{"regex":"v[a-zA-z0-9]+","match":"regex"}}]` , err: nil},			
		},
		{
			name: "RulesBoolean",
			input: in{exampleValue: map[string]any{"default":true}, matchType: map[string]any{"default": "boolean"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`[{"type":"boolean","values":{"match":"boolean"}}]` , err: nil},			
		},
		{
			name: "RulesContentType",
			input: in{exampleValue: map[string]any{"default":"{}"}, matchType: map[string]any{"default": "contentType"}, matchTypeConfig: map[string]any{"default":"application/avro"} },
			want: out{rules:`[{"type":"contentType","values":{"match":"contentType","value":"application/avro"}}]` , err: nil},			
		},
		{
			name: "RulesEachValueType",
			input: in{exampleValue: map[string]any{"eachValue":100}, matchType: map[string]any{"eachValue": "type"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`[{"type":"eachValue","values":{"rules":[{"match":"type"}],"value":100}}]` , err: nil},			
		},
		{
			name: "RulesEachValueReference",
			input: in{exampleValue: map[string]any{"reference": true, "eachValue":"subProducts"}, matchType: map[string]any{"reference": "values", "eachValue":"$"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`TODO:[{"type":"values","values":{"match":"values"}}]` , err: nil},			
		},
		{
			name: "RulesEachKeyEachValue",
			input: in{exampleValue: map[string]any{"eachKey":"key1", "eachValue":"value1"}, matchType: map[string]any{"eachKey":"type", "eachValue":"notEmpty"}, matchTypeConfig: map[string]any{} },
			want: out{rules:`[{"type":"eachKey","values":{"rules":[{"match":"type"}],"value":"key1"}},{"type":"eachValue","values":{"rules":[{"match":"notEmpty"}],"value":"value1"}}]` , err: nil},			
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := out{}
			matchingRules, err := buildMatchingRules(test.input.matchType, test.input.matchTypeConfig, test.input.exampleValue)			 
			got.err = err
			bytes, e := json.Marshal(matchingRules.GetRule())			
			if e != nil {
				log.Println("Error while json marshalling, ", e)				
			} else {
				got.rules = string(bytes)
			}			
			log.Println("got: ",  got.rules)
			log.Println("want: ", test.want.rules)
			require.ErrorIs(t, got.err, test.want.err)			
			assert.JSONEq(t, test.want.rules , got.rules)
		})
	}
	
}


