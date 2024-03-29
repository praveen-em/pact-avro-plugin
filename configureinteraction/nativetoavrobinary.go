package configureinteraction

import (
	"encoding/json"
	"log"

	"github.com/hamba/avro/v2"
)

func convertNativeToAvroBinary(schema avro.Schema, data map[string]interface{}) ([]byte, error) {
	//TODO: remove this temp experiment
	//temp start
	var tempData map[string]interface{}
	err := json.Unmarshal([]byte(schema.String()), &tempData)
	if err != nil {
		log.Println("ERROR while marshalling json", err)
		return nil, err
	}

	if fields, ok := tempData["fields"].([]interface{}); ok {
		parseSchemaFields(parseFields(fields), "", false)
	}
	log.Println("Schema Fields: ", schemaTypesContainer)
	log.Println("Union Fields: ", schemaUnionContainer)
	//temp end

	log.Println("Native content: ", data)

	binary, err := avro.Marshal(schema, data)
	if err != nil {
		log.Println("ERROR while trying to marshal from Native to Binary. ", err.Error())
		return nil, err
	}

	return binary, err
}
