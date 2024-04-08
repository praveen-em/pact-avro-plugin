package configureinteraction

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchingRulesBuilder(t *testing.T) {
	t.Parallel() //marks MatchingRulesBuilder as capable of running in parallel with other tests

	type in struct {
		exampleValueMap map[string]any
		matchTypeMap map[string]any
		matchTypeConfigMap map[string]any
	}

	type out struct {
		rules string
		err error
	}
	
	//List of test cases (table driven tests)
	tests := []struct {
		name string
		input in
		want out
	}{
		{
			name: "RulesType",
			input: in{exampleValueMap: map[string]any{"default":"value1"}, matchTypeMap: map[string]any{"default": "type"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"type", "values":{"match":"type"}}]` , err: nil},			
		},
		{
			name: "RulesEqualTo",
			input: in{exampleValueMap: map[string]any{"default":"value1"}, matchTypeMap: map[string]any{"default": "equalTo"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"equalTo", "values":{"match":"equalTo"}}]` , err: nil},			
		},
		{
			name: "RulesDate",
			input: in{exampleValueMap: map[string]any{"default":"2023-Jul-08"}, matchTypeMap: map[string]any{"default": "date"}, matchTypeConfigMap: map[string]any{"default": "2006-01-02"} },
			want: out{rules:`[{"type":"date","values":{"format":"2006-01-02","match":"date"}}]` , err: nil},			
		},
		{
			name: "RulesDateTime",
			input: in{exampleValueMap: map[string]any{"default":"2023-07-21 16:44:32"}, matchTypeMap: map[string]any{"default": "datetime"}, matchTypeConfigMap: map[string]any{"default": "2006-01-02 15:04:05"} },
			want: out{rules:`[{"type":"datetime","values":{"format":"2006-01-02 15:04:05","match":"datetime"}}]` , err: nil},			
		},
		{
			name: "RulesTime",
			input: in{exampleValueMap: map[string]any{"default":"16:44:32"}, matchTypeMap: map[string]any{"default": "time"}, matchTypeConfigMap: map[string]any{"default": "15:04:05"} },
			want: out{rules:`[{"type":"time","values":{"format":"15:04:05","match":"time"}}]` , err: nil},			
		},
		{
			name: "RulesRegex",
			input: in{exampleValueMap: map[string]any{"default":"v2"}, matchTypeMap: map[string]any{"default": "regex"}, matchTypeConfigMap: map[string]any{"default": "v[a-zA-z0-9]+"} },
			want: out{rules:`[{"type":"regex","values":{"regex":"v[a-zA-z0-9]+","match":"regex"}}]` , err: nil},			
		},
		{
			name: "RulesBoolean",
			input: in{exampleValueMap: map[string]any{"default":true}, matchTypeMap: map[string]any{"default": "boolean"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"boolean","values":{"match":"boolean"}}]` , err: nil},			
		},
		{
			name: "RulesContentType",
			input: in{exampleValueMap: map[string]any{"default":"{}"}, matchTypeMap: map[string]any{"default": "contentType"}, matchTypeConfigMap: map[string]any{"default":"application/avro"} },
			want: out{rules:`[{"type":"contentType","values":{"match":"contentType","value":"application/avro"}}]` , err: nil},			
		},
		{
			name: "RulesEachValueType",
			input: in{exampleValueMap: map[string]any{"eachValue":100}, matchTypeMap: map[string]any{"eachValue": "type"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"eachValue","values":{"rules":[{"match":"type"}],"value":100}}]` , err: nil},			
		},
		//TODO: Needs rethink on want for RulesEachValueReference
		{
			name: "RulesEachValueReference",
			input: in{exampleValueMap: map[string]any{"reference": true, "eachValue":"subProducts"}, matchTypeMap: map[string]any{"reference": "values", "eachValue":"$"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"values","values":{"match":"values"}}]` , err: nil},			
		},
		{
			name: "RulesEachKeyEachValue",
			input: in{exampleValueMap: map[string]any{"eachKey":"key1", "eachValue":"value1"}, matchTypeMap: map[string]any{"eachKey":"type", "eachValue":"notEmpty"}, matchTypeConfigMap: map[string]any{} },
			want: out{rules:`[{"type":"eachKey","values":{"rules":[{"match":"type"}],"value":"key1"}},{"type":"eachValue","values":{"rules":[{"match":"notEmpty"}],"value":"value1"}}]` , err: nil},			
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T){
			//Arrange
			got := out{}

			//Act
			rules, err := buildMatchingRules(test.input.matchTypeMap, test.input.matchTypeConfigMap, test.input.exampleValueMap)			 
			got.err = err

			bytes, err := json.Marshal(rules.GetRule())			
			if err != nil {
				log.Println("Error while json marshalling, ", err)				
			} else {
				got.rules = string(bytes)
			}
			// log.Println("got: ",  got.rules)
			// log.Println("want: ", test.want.rules)

			//Assert
			require.ErrorIs(t, got.err, test.want.err)				
			assert.JSONEq(t, test.want.rules , got.rules) //sometimes fail, same problem as https://github.com/stretchr/testify/issues/1025. needs fixing.
		})
	}
	
}


