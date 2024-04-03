package configureinteraction

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/hamba/avro/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestInteractionBuilder(t *testing.T) {
		
	type out struct {
		content string
		rules string
		err error
	}
	
	tests := []struct {
		name string
		data string
		avroSchema string
		want out
	}{
		{
			name: "InteractionRecordTypeBoolean",
			data: `{"available":"matching(boolean, true)"}`,
			avroSchema: `{"name":"testEvent","type":"record","fields":[{"name":"available","type":"boolean"}]}`,
			want: out{ content: `{"available":true}`,
						rules:`{"$.available":{"rule":[{"type":"boolean","values":{"match":"boolean"}}]}}`,
						err: nil},			
		},
		{
			name: "InteractionRecordTypeLogicalDate",
			data: `{"createdOn":"matching(date, '2006-01-02', '2023-07-08')" }`, 			
			avroSchema: `{"name":"testEvent","type":"record","fields":[{"name":"createdOn","type":["null",{"type":"int","logicalType":"date"}] }]}`,
			want: out{content: `{"createdOn":"2023-07-08T00:00:00Z"}`, 
						rules:`{"$.createdOn":{"rule":[{"type":"date","values":{"format":"2006-01-02","match":"date"}}]}}` ,
						err: nil},			
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			got := out{}

			records := make(map[string]*structpb.Value)			
			err := json.Unmarshal([]byte(test.data), &records)
			if err != nil {
				log.Println("Error while json unmarshalling, ", err)				
			}

			schema := avroSchemaParse(test.avroSchema)
			content, rules, err := buildInteraction(records, schema)		 			
			got.err = err
			require.ErrorIs(t, got.err, test.want.err)		
			
			got.content = jsonMarshal(content)			
			got.rules = jsonMarshal(rules)			
						
			// log.Println("got content: ",  got.content)
			// log.Println("want content: ", test.want.content)

			// log.Println("got rules: ",  got.rules)
			// log.Println("want rules: ", test.want.rules)
									
			assert.JSONEq(t, test.want.content, got.content)
			assert.JSONEq(t, test.want.rules , got.rules)
		})
	}
	
}

func jsonMarshal(input any) (output string) {
	bytes, err := json.Marshal(input)			
	if err != nil {
		log.Println("Error while json marshalling, ", err)				
	} 
	return string(bytes)	
}

func avroSchemaParse(inputSchema string) (output avro.Schema) {
	output, err := avro.Parse(inputSchema)
	if err != nil {
		log.Println("Error while avro schema parsing, ", err)				
	}
	return output
}


