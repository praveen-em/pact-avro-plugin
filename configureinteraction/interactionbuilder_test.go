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
	t.Parallel() //marks InteractionBuilder as capable of running in parallel with other tests
	sPrefix := `{"name":"ProductEvent","type":"record","fields":[`
	sSuffix := `]}`

	type out struct {
		content string
		rules string
		err error
	}
	
	//List of test cases (table driven tests)
	tests := []struct {
		name string
		data string
		avroSchema string
		want out
	}{
		{
			name: "PayloadWithBoolean",
			data: `{"available":"matching(boolean, true)"}`,
			avroSchema: sPrefix+`{"name":"available","type":"boolean"}`+sSuffix,
			want: out{ content: `{"available":true}`,
						rules: `{"$.available":{"rule":[{"type":"boolean","values":{"match":"boolean"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEnum",
			data: `{"event":"matching(regex, '^(CREATED|UPDATED|DELETED)$', 'CREATED')"}`,
			avroSchema: sPrefix+`{"name":"event","type":{"name":"EventType","type":"enum","symbols":["CREATED","UPDATED","DELETED","UNKNOWN"]}}`+sSuffix,
			want: out{ content: `{"event":"CREATED"}`,
						rules: `{"$.event":{"rule":[{"type":"regex","values":{"match":"regex","regex":"^(CREATED|UPDATED|DELETED)$"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithLogicalDateSchemaUnion",
			data: `{"createdOn":"matching(date, '2006-01-02', '2023-07-08')" }`, 			
			avroSchema: sPrefix+`{"name":"createdOn","type":["null",{"type":"int","logicalType":"date"}] }`+sSuffix,
			want: out{content: `{"createdOn":"2023-07-08T00:00:00Z"}`, 
						rules: `{"$.createdOn":{"rule":[{"type":"date","values":{"format":"2006-01-02","match":"date"}}]}}` ,
						err: nil},			
		},
		{
			name: "PayloadWithDecimalSchemaUnion",
			data: `{"price":"matching(decimal, 99.89)"}`,
			avroSchema: sPrefix+`{"name":"price","type":["null",{"type":"bytes","logicalType":"decimal","precision":6,"scale":2}]}`+sSuffix,
			want: out{ content: `{"price":"9989/100"}`,
						rules: `{"$.price":{"rule":[{"type":"decimal","values":{"match":"decimal"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithRecordSchemaUnion",
			data: `{"location":{
				"doorNumber":"matching(integer, 29)",
				"postcode":"notEmpty('GU15 7SR')",
				"street":"notEmpty('cross street')"}}`,
			avroSchema: sPrefix+`{"name":"location","type":["null",{"name":"Address","type":"record","fields":[{"name":"doorNumber","type":"int"},{"name":"street","type":"string"},{"name":"postcode","type":"string"}]}]}`+sSuffix,
			want: out{ content: `{"location":{"Address":{"doorNumber":29,"postcode":"GU15 7SR","street":"cross street"}}}`,
						rules:`{"$.location.doorNumber":{"rule":[{"type":"integer","values":{"match":"integer"}}]},
								"$.location.postcode":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]},
								"$.location.street":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithArrayInt",
			data: `{"items":["matching(type,10)","matching(type,11)","matching(type, 12)"]}`,
			avroSchema: sPrefix+`{"name":"items","type":{"type":"array","items":"int"}}`+sSuffix,
			want: out{ content: `{"items":[10,11,12]}`,
						rules: `{"$.items":{"rule":[{"type":"type","values":{"match":"type"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithArrayIntSchemaUnion",
			data: `{"items":["matching(type,10)","matching(type,11)","matching(type, 12)"]}`,
			avroSchema: sPrefix+`{"name":"items","type":["string",{"type":"array","items":"int"}]}`+sSuffix,
			want: out{ content: `{"items":{"array":[10,11,12]}}`,
						rules: `{"$.items":{"rule":[{"type":"type","values":{"match":"type"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithStringSchemaUnion",
			data: `{"item":"matching(type,'item1')"}`,
			avroSchema: sPrefix+`{"name":"item","type":["string",{"type":"array","items":"int"}]}`+sSuffix,
			want: out{ content: `{"item":"item1"}`,
						rules: `{"$.item":{"rule":[{"type":"type","values":{"match":"type"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithArrayStringSchemaUnion",
			data: `{"items":["matching(type, 'item1')","matching(type, 'item2')","matching(type, 'item3')"]}`,
			avroSchema: sPrefix+`{"name":"items","type":["null",{"type":"array","items":"string"}]}`+sSuffix,
			want: out{ content: `{"items":{"array":["item1","item2","item3"]}}`,
						rules: `{"$.items":{"rule":[{"type":"type","values":{"match":"type"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithMapStringSchemaUnion",
			data: `{"items":{"key1":"matching(equalTo, 'value1')","key2":"notEmpty('value2')"}}`,
			avroSchema: sPrefix+`{"name":"items","type":["null",{"type":"map","values":"string"}]}`+sSuffix,
			want: out{ content: `{"items":{"map":{"key1":"value1","key2":"value2"}}}`,
						rules: `{"$.items.key1":{"rule":[{"type":"equalTo","values":{"match":"equalTo"}}]},
								 "$.items.key2":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEachValueSchemaUnion",
			data: `{"items":{"pact:match":"eachValue(matching(type, 100))"}}`,
			avroSchema: sPrefix+`{"name":"items","type":["null",{"type":"array","items":"int"}]}`+sSuffix,
			want: out{ content: `{"items":{"array":[100]}}`,
						rules: `{"$.items":{"rule":[{"type":"eachValue","values":{"rules":[{"match":"type"}],"value":100}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEachValue",
			data: `{"items":{"pact:match":"eachValue(matching(type, 100))"}}`,
			avroSchema: sPrefix+`{"name":"items","type":{"type":"array","items":"int"}}`+sSuffix,
			want: out{ content: `{"items":[100]}`,
						rules: `{"$.items":{"rule":[{"type":"eachValue","values":{"rules":[{"match":"type"}],"value":100}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEachKeyEachValueSchemaUnion",
			data: `{"items":{"pact:match":"eachKey(matching(type, 'key1')), eachValue(notEmpty('value1'))"}}`,
			avroSchema: sPrefix+`{"name":"items","type":["null",{"type":"map","values":"string"}]}`+sSuffix,
			want: out{ content: `{"items":{"map":{"key1":"value1"}}}`,
						rules: `{"$.items":{"rule":[{"type":"eachKey","values":{"rules":[{"match":"type"}],"value":"key1"}},
													{"type":"eachValue","values":{"rules":[{"match":"notEmpty"}],"value":"value1"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEachValueReferenceSchemaUnion",
			data: `{"items":{"pact:match":"eachValue(matching($'subProducts'))","subProducts":[{"id":"notEmpty(1)","name":"notEmpty('name1')"}]}}`,
			avroSchema: sPrefix+`{"name":"items","type":["int",{"type":"array","items":{"name":"subProduct","type":"record","fields":[{"name":"name","type":"string"},{"name":"id","type":"int"}]}}]}`+sSuffix,
			want: out{ content: `{"items":{"array":[{"id":1,"name":"name1"}]}}`,
						rules: `{"$.items":{"rule":[{"type":"values","values":{"match":"values"}}]},
								 "$.items.*":{"rule":[{"type":"type","values":{"match":"type"}}]},
								 "$.items.*.id":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]},
								 "$.items.*.name":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]}}`,
						err: nil},			
		},
		{
			name: "PayloadWithEachValueReference",
			data: `{"items":{"pact:match":"eachValue(matching($'subProducts'))","subProducts":[{"id":"notEmpty(1)"}]}}`,
			avroSchema: sPrefix+`{"name":"items","type":{"type":"array","items":{"name":"subProduct","type":"record","fields":[{"name":"id","type":"int"}]}}}`+sSuffix,
			want: out{ content: `{"items":[{"id":1}]}`,
						rules: `{"$.items":{"rule":[{"type":"values","values":{"match":"values"}}]},
								 "$.items.*":{"rule":[{"type":"type","values":{"match":"type"}}]},								
								 "$.items.*.id":{"rule":[{"type":"notEmpty","values":{"match":"notEmpty"}}]}}`,
						err: nil},			
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T){
			//Arrange
			got := out{}
			records := make(map[string]*structpb.Value)			
			err := json.Unmarshal([]byte(test.data), &records)
			if err != nil {
				log.Println("Error while json unmarshalling, ", err)				
			}
			schema := avroSchemaParse(test.avroSchema)

			//Act
			content, rules, err := buildInteraction(records, schema)		 			
			got.err = err
			got.content = jsonMarshal(content)			
			got.rules = jsonMarshal(rules)				

			// log.Println("got content: ",  got.content)
			// log.Println("want content: ", test.want.content)
			// log.Println("got rules: ",  got.rules)
			// log.Println("want rules: ", test.want.rules)

			//Assert
			require.ErrorIs(t, got.err, test.want.err)							
			assert.JSONEq(t, test.want.content, got.content)
			assert.JSONEq(t, test.want.rules , got.rules)  //sometimes fail, same problem as https://github.com/stretchr/testify/issues/1025. needs fixing.			
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


