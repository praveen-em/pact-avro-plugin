package interaction

import (
	"log"
	"time"
)

func parseExpression(expression string) (exampleValueMap map[string]interface{}, exampleValue interface{}, matchTypeMap map[string]interface{}, matchTypeConfigMap map[string]interface{}, err error) {
	
	matchTypeMap, matchTypeConfigMap, exampleValueMap = parseMatchingRuleDefinition(expression)
	
	var (
		isMap       bool
		isReference bool
		eachKey     string
	)

	if v, ok := exampleValueMap["eachKey"]; ok {
		matchType, _ := matchTypeMap["eachKey"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["eachKey"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}
		isMap = true
		eachKey = v.(string)
	}

	if _, ok := exampleValueMap["reference"]; ok {
		isReference = true
	}

	if v, ok := exampleValueMap["default"]; ok {
		matchType, _ := matchTypeMap["default"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["default"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}
		exampleValue = v
	}

	if v, ok := exampleValueMap["eachValue"]; ok {
		matchType, _ := matchTypeMap["eachValue"].(string)
		matchTypeConfig, _ := matchTypeConfigMap["eachValue"].(string)
		v, err = convertToNativeType(matchType, matchTypeConfig, v)
		if err != nil {
			return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
		}

		switch {
		case isMap:
			m := make(map[string]any)
			m[eachKey] = v
			exampleValue = m

		case isReference:
			exampleValue = v

		default:
			var s []interface{}
			s = append(s, v)
			exampleValue = s
		}

	}
	return exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err
}

func convertToNativeType(matchType string, matchTypeConfig string, exampleValue interface{}) (interface{}, error) {
	//TODO: cover all data types
	var err error
	switch matchType {

	case "date":
		{
			date, err := time.Parse(matchTypeConfig, exampleValue.(string))
			if err != nil {
				log.Println("ERROR While converting to native date format. ", err)
				return nil, err
			} else {
				exampleValue = date
			}
		}
	}
	return exampleValue, err
}