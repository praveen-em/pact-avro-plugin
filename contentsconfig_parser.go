package main

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hamba/avro/v2"
	plugin "github.com/praveen-em/pact-avro-plugin/io_pact_plugin"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/structpb"
)

type configuration struct {
	content       map[string]interface{}
	contentBinary []byte
	rules         map[string]*plugin.MatchingRules
	schema        avro.Schema
}

var rules = make(map[string]*plugin.MatchingRules)
var avroPrimitiveTypes = []string{"null", "boolean", "int", "long", "float", "double", "bytes", "string"}
var referenceList []string

func parseContentsConfig(req *plugin.ConfigureInteractionRequest) (*configuration, error) {
	records := req.GetContentsConfig().GetFields()
	rulesPath := "$."

	if _, ok := records["pact:schema"]; !ok {
		err := errors.New("Config item with key 'pact:schema' and the Avro schema string is required")
		log.Println(err.Error())
		return nil, err
	}

	schema, err := avro.Parse(records["pact:schema"].GetStringValue())
	if err != nil {
		log.Println("ERROR parsing schema. ", err.Error())
		return nil, err
	}
	log.Println("Schema: ", schema)

	content, err := iterateRecords(records, schema, rules, rulesPath)
	if err != nil {
		return nil, err
	}

	contentBinary, err := convertNativeToAvroBinary(schema, content)
	if err != nil {
		return nil, err
	}

	var parsedData = &configuration{content, contentBinary, rules, schema}
	return parsedData, err
}

func iterateRecords(records map[string]*structpb.Value, schema avro.Schema, rules map[string]*plugin.MatchingRules, rulesPath string) (content map[string]interface{}, err error) {
	content = make(map[string]interface{})
	var (
		exampleValue    interface{}
		matchType       string
		matchTypeConfig string
	)

	if rulesPath == "" {
		rulesPath = "$."
	}

	for key, value := range records {
		// log.Println("ContentsConfig iterated:", key, value)
		var isPayload bool
		if !strings.HasPrefix(key, "pact:") || key == "pact:match" {
			isPayload = true
		}

		if isPayload {
			switch value.Kind.(type) {
			case *structpb.Value_StringValue:
				expression := value.GetStringValue()

				exampleValue, matchType, matchTypeConfig, err = parseExpression(expression)
				if err != nil {
					return content, err
				}

				matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
				if err != nil {
					return content, err
				}

				switch matchType {
				case "values":
					if s, ok := exampleValue.(string); ok {
						referenceList = append(referenceList, s)
						rules[strings.TrimSuffix(rulesPath, ".")] = matchingRules
					}

				default:
					content[key] = exampleValue
					if unionType, ok := isUnion(key, schema); ok {
						switch {
						case !slices.Contains(avroPrimitiveTypes, unionType):
							content[key] = map[string]any{unionType: exampleValue}
						}
					}
					rules[rulesPath+key] = matchingRules

				}

			case *structpb.Value_StructValue:
				var contentMap map[string]interface{}
				var contentGeneric interface{}
				rulesPath = rulesPath + key + "."

				contentMap, err = iterateRecords(value.GetStructValue().Fields, schema, rules, rulesPath)
				if err != nil {
					return content, err
				}

				contentGeneric = contentMap
				if v, ok := contentMap["pact:match"]; ok {
					contentGeneric = v
				}

				if len(referenceList) > 0 {
					for k, v := range contentMap {
						if slices.Contains(referenceList, k) {
							contentGeneric = v
							break
						}
					}
				}

				rulesPath = strings.TrimSuffix(rulesPath, key+".")

				if unionType, ok := isUnion(key, schema); ok {
					content[key] = map[string]any{unionType: contentGeneric}
				} else {
					content[key] = contentGeneric
				}

			case *structpb.Value_ListValue:
				var contentList []interface{}
				var contentMap map[string]interface{}
				valuesList := value.GetListValue().Values

				for index := range valuesList {

					switch valuesList[index].Kind.(type) {
					case *structpb.Value_StructValue:
						rulesPath = rulesPath + key + "."
						contentMap, err = iterateRecords(valuesList[index].GetStructValue().Fields, schema, rules, rulesPath)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, contentMap)
						rulesPath = strings.TrimSuffix(rulesPath, key+".")

					case *structpb.Value_StringValue:
						expression := valuesList[index].GetStringValue()

						exampleValue, matchType, matchTypeConfig, err = parseExpression(expression)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, exampleValue)
						matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
						if err != nil {
							return content, err
						}
						rules[rulesPath+key] = matchingRules
					}

					if unionType, ok := isUnion(key, schema); ok {
						content[key] = map[string]any{unionType: contentList}
					} else {
						content[key] = contentList
					}
				}
			}
		}
	}
	// log.Println("Parsed content:", content)
	// log.Println("Parsed Rules: ", rules)
	return content, err
}

// TODO: optimize
func isUnion(key string, schema avro.Schema) (unionType string, ok bool) {
	query := "fields.#(name=\"" + key + "\").type"
	result := searchJson(schema.String(), query)
	if (result.Exists()) && (reflect.TypeOf(result.Value()).Kind()) == reflect.Slice {
		query = "fields.#(name=\"" + key + "\").type.#.type"
		result := searchJson(schema.String(), query)
		if result.Exists() {
			resultSanitized := strings.Trim(result.String(), "[]\"")
			switch resultSanitized {
			case "record", "enum", "fixed":
				query = "fields.#(name=\"" + key + "\").type.#.name"
				result := searchJson(schema.String(), query)
				if result.Exists() {
					resultSanitized = strings.Trim(result.String(), "[]\"")
				}
			}
			return resultSanitized, true
		}
	}
	return "", false
}

func searchJson(json string, query string) gjson.Result {
	return gjson.Get(json, query)
}

func parseExpression(expression string) (exampleValue interface{}, matchType string, matchTypeConfig string, err error) {
	matchType, matchTypeConfig, exampleValue = parseMatchingRuleDefinition(expression)
	var (
		isMap       bool
		isReference bool
		eachKey     string
	)
	if exampleValueMap, ok := exampleValue.(map[string]interface{}); ok {
		if v, ok := exampleValueMap["eachKey"]; ok {
			isMap = true
			eachKey = v.(string)
		}

		if _, ok := exampleValueMap["reference"]; ok {
			isReference = true
		}

		if v, ok := exampleValueMap["default"]; ok {
			exampleValue = v
		}

		if v, ok := exampleValueMap["eachValue"]; ok {
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
	}

	exampleValue, err = convertToNativeType(matchType, matchTypeConfig, exampleValue)
	if err != nil {
		return exampleValue, matchType, matchTypeConfig, err
	}

	return exampleValue, matchType, matchTypeConfig, err
}

func buildMatchingRules(matchType string, matchTypeConfig string) (*plugin.MatchingRules, error) {
	matchingRules := &plugin.MatchingRules{}

	matchingRule, err := buildMatchingRule(matchType, matchTypeConfig)
	if err != nil {
		return matchingRules, err
	}

	matchingRules.Rule = append(matchingRules.Rule, matchingRule)
	return matchingRules, err
}

func buildMatchingRule(matchType string, matchTypeConfig string) (matchingRule *plugin.MatchingRule, err error) {
	matchingRule = &plugin.MatchingRule{}
	matchingRule.Type = matchType
	values := make(map[string]interface{})
	values["match"] = matchType
	if matchTypeConfig != "" {
		matchTypeConfigKey := "format"
		switch {
		case matchType == "regex":
			matchTypeConfigKey = "regex"
		case matchType == "datetime" || matchType == "date" || matchType == "time":
			matchTypeConfigKey = "format"
		case matchType == "contentType":
			matchTypeConfigKey = "value"
		}
		values[matchTypeConfigKey] = matchTypeConfig
	}

	matchingRule.Values, err = structpb.NewStruct(values)
	if err != nil {
		log.Println("ERROR while converting from native go to structpb format. ", err)
		return matchingRule, err
	}
	return matchingRule, err

}

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
		iterateSchemaFields(parseFields(fields), "", false)
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
