package configureinteraction

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpresionParser(t *testing.T) {

	type output struct {
		exampleValueMap    map[string]any
		exampleValue       any
		matchTypeMap       map[string]any
		matchTypeConfigMap map[string]any
		err                error
	}

	tests := []struct {
		name string
		input string
		want output	
	}{
		{
			name: "NotEmpty",
			input: "notEmpty('value1')",
			want: output{map[string]any{"default": "value1"}, "value1", map[string]any{"default": "notEmpty"}, map[string]any{}, nil},
		},
		{
			name: "MatchingTypeString",
			input: "matching(type, 'value1')",
			want: output{map[string]any{"default": "value1"}, "value1", map[string]any{"default": "type"}, map[string]any{}, nil},
		},
		{
			name: "MatchingTypeInteger",
			input: "matching(type, 99)",
			want: output{map[string]any{"default": 99}, 99, map[string]any{"default": "type"}, map[string]any{}, nil},
		},
		{
			name: "MatchingTypeBoolean",
			input: "matching(type, true)",
			want: output{map[string]any{"default": true}, true, map[string]any{"default": "type"}, map[string]any{}, nil},
		},
		// {
		// 	name: "MatchingTypeDecimal",
		// 	input: "matching(type, 99.99)",
		// 	want: output{map[string]any{"default": 99.99}, 99.99, map[string]any{"default": "type"}, map[string]any{}, nil},
		// },
		{
			name: "MatchingEqualToString",
			input: "matching(equalTo, 'value1')",
			want: output{map[string]any{"default": "value1"}, "value1", map[string]any{"default": "equalTo"}, map[string]any{}, nil},
		},
		{
			name: "MatchingEqualToInteger",
			input: "matching(equalTo, 99)",
			want: output{map[string]any{"default": 99}, 99, map[string]any{"default": "equalTo"}, map[string]any{}, nil},
		},
		{
			name: "MatchingEqualToBoolean",
			input: "matching(equalTo, false)",
			want: output{map[string]any{"default": false}, false, map[string]any{"default": "equalTo"}, map[string]any{}, nil},
		},
		// {
		// 	name: "MatchingEqualToDecimal",
		// 	input: "matching(equalTo, 0.01)",
		// 	want: output{map[string]any{"default": 0.01}, 0.01, map[string]any{"default": "equalTo"}, map[string]any{}, nil},
		// },
		{
			name: "MatchingNumberInteger",
			input: "matching(number, 99)",
			want: output{map[string]any{"default": 99}, 99, map[string]any{"default": "number"}, map[string]any{}, nil},
		},
		// {
		// 	name: "MatchingNumberDecimal",
		// 	input: "matching(number, 99.10)",
		// 	want: output{map[string]any{"default": 99.10}, 99.10, map[string]any{"default": "number"}, map[string]any{}, nil},
		// },
		{
			name: "MatchingDate1",
			input: "matching(date, '2006-01-02', '2023-07-08')",
			want: output{map[string]any{"default": "2023-07-08"}, time.Time(time.Date(2023, time.July, 8, 0, 0, 0, 0, time.UTC)), map[string]any{"default": "date"}, map[string]any{"default": "2006-01-02"}, nil},
		},
		{
			name: "MatchingDate2",
			input: "matching(date, '2006-Jan-02', '2023-Jul-08')",
			want: output{map[string]any{"default": "2023-Jul-08"}, time.Time(time.Date(2023, time.July, 8, 0, 0, 0, 0, time.UTC)), map[string]any{"default": "date"}, map[string]any{"default": "2006-Jan-02"}, nil},
		},
		{
			name: "MatchingDateTime",
			input: "matching(datetime, '2006-01-02 15:04:05', '2023-07-21 16:44:32')",
			want: output{map[string]any{"default": "2023-07-21 16:44:32"}, time.Date(2023, time.July, 21, 16, 44, 32, 0, time.UTC), map[string]any{"default": "datetime"}, map[string]any{"default": "2006-01-02 15:04:05"}, nil},
		},
		{
			name: "MatchingDateTimeMicro",
			input: "matching(datetime, '2006-01-02 15:04:05.00000', '2023-07-21 16:44:32.12345')",
			want: output{map[string]any{"default": "2023-07-21 16:44:32.12345"}, time.Date(2023, time.July, 21, 16, 44, 32, 123450000, time.UTC), map[string]any{"default": "datetime"}, map[string]any{"default": "2006-01-02 15:04:05.00000"}, nil},
		},
		{
			name: "MatchingTime",
			input: "matching(time, '15:04:05', '16:44:32')",
			want: output{map[string]any{"default": "16:44:32"}, time.Date(0, time.January, 1, 16, 44, 32, 0, time.UTC), map[string]any{"default": "time"}, map[string]any{"default": "15:04:05"}, nil},
		},
		{
			name: "MatchingRegex",
			input: "matching(regex, 'v[a-zA-z0-9]+', 'v1')",
			want: output{map[string]any{"default": "v1"}, "v1", map[string]any{"default": "regex"}, map[string]any{"default":"v[a-zA-z0-9]+"}, nil},
		},
		{
			name: "MatchingInclude",
			input: "matching(include, 'testing')",
			want: output{map[string]any{"default": "testing"}, "testing", map[string]any{"default": "include"}, map[string]any{}, nil},
		},
		{
			name: "MatchingBoolean",
			input: "matching(boolean, true)",
			want: output{map[string]any{"default": true}, true, map[string]any{"default": "boolean"}, map[string]any{}, nil},
		},
		{
			name: "MatchingSemver",
			input: "matching(semver, '1.0.0')",
			want: output{map[string]any{"default": "1.0.0"}, "1.0.0", map[string]any{"default": "semver"}, map[string]any{}, nil},
		},
		{
			name: "MatchingContentType",
			input: "matching(contentType, 'application/avro', '{}')",
			want: output{map[string]any{"default": "{}"}, "{}", map[string]any{"default": "contentType"}, map[string]any{"default":"application/avro"}, nil},
		},
		{
			name: "eachValueMatchingReference",
			input: "eachValue(matching($'subProducts'))",
			want: output{map[string]any{"reference": true, "eachValue":"subProducts"}, "subProducts", map[string]any{"reference": "values", "eachValue":"$"}, map[string]any{}, nil},
		},
		{
			name: "eachValueNotEmpty",
			input: "eachValue(notEmpty('value1'))",
			want: output{map[string]any{"eachValue":"value1"}, []interface {}{"value1"}, map[string]any{"eachValue":"notEmpty"}, map[string]any{}, nil},
		},
		{
			name: "eachKeyEachValue",
			input: "eachKey(matching(type, 'key1')), eachValue(notEmpty('value1'))",
			want: output{map[string]any{"eachKey":"key1", "eachValue":"value1"}, map[string]interface{}{"key1":"value1"}, map[string]any{"eachKey":"type", "eachValue":"notEmpty"}, map[string]any{}, nil},
		},
	
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := output{}
			got.exampleValueMap, got.exampleValue, got.matchTypeMap, got.matchTypeConfigMap, got.err = parseExpression(test.input)
			log.Println("result: ", got.exampleValueMap, got.exampleValue, got.matchTypeMap, got.matchTypeConfigMap, got.err)
			require.ErrorIs(t, got.err, test.want.err)
			assert.Equal(t, test.want.exampleValueMap, got.exampleValueMap)
			assert.Equal(t, test.want.exampleValue, got.exampleValue)
			assert.Equal(t, test.want.matchTypeMap, got.matchTypeMap)
			assert.Equal(t, test.want.matchTypeConfigMap, got.matchTypeConfigMap)
		})
	}
}