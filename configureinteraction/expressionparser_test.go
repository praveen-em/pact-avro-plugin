package configureinteraction

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
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
			name: "MatchingTypeNumber",
			input: "matching(type, 99)",
			want: output{map[string]any{"default": 99}, 99, map[string]any{"default": "type"}, map[string]any{}, nil},
		},
		{
			name: "MatchingTypeBoolean",
			input: "matching(type, true)",
			want: output{map[string]any{"default": true}, true, map[string]any{"default": "type"}, map[string]any{}, nil},
		},
	
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := output{}
			got.exampleValueMap, got.exampleValue, got.matchTypeMap, got.matchTypeConfigMap, got.err = parseExpression(test.input)
			log.Println("result: ", got.exampleValueMap, got.exampleValue, got.matchTypeMap, got.matchTypeConfigMap, got.err)
			require.NoError(t, got.err)
			assert.Equal(t, test.want.exampleValueMap, got.exampleValueMap)
			assert.Equal(t, test.want.exampleValue, got.exampleValue)
			assert.Equal(t, test.want.matchTypeMap, got.matchTypeMap)
			assert.Equal(t, test.want.matchTypeConfigMap, got.matchTypeConfigMap)
		})
	}
}