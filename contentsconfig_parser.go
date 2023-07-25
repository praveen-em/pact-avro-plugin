package main

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hamba/avro/v2"
	plugin "github.com/praveen-em/pact-avro-plugin/io_pact_plugin"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/types/known/structpb"
)

type configuration struct {
	content       map[string]interface{}
	contentBinary []byte
	rules         map[string]*plugin.MatchingRules
	schema        avro.Schema
}

var rules = make(map[string]*plugin.MatchingRules)

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
		// log.Println("ContentsConfig interated:", key, value)

		if !strings.HasPrefix(key, "pact:") {
			switch value.Kind.(type) {
			case *structpb.Value_StringValue:
				expression := value.GetStringValue()

				exampleValue, matchType, matchTypeConfig, err = parseExpression(expression)
				if err != nil {
					return content, err
				}
				content[key] = exampleValue

				matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
				if err != nil {
					return content, err
				}

				rules[rulesPath+key] = matchingRules

			case *structpb.Value_StructValue:
				var contentMap map[string]interface{}
				rulesPath = rulesPath + key + "."

				contentMap, err = iterateRecords(value.GetStructValue().Fields, schema, rules, rulesPath)
				if err != nil {
					return content, err
				}
				rulesPath = strings.TrimSuffix(rulesPath, key+".")

				if unionType, ok := isUnion(key, schema); ok {
					content[key] = map[string]any{unionType: contentMap}
				} else {
					content[key] = contentMap
				}

			case *structpb.Value_ListValue:
				var contentList []interface{}
				var contentMap map[string]interface{}
				valuesList := value.GetListValue().Values
				for index := range valuesList {

					switch valuesList[index].Kind.(type) {
					case *structpb.Value_StructValue:
						contentMap, err = iterateRecords(valuesList[index].GetStructValue().Fields, schema, rules, rulesPath)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, contentMap)

					case *structpb.Value_StringValue:
						expression := valuesList[index].GetStringValue()

						exampleValue, matchType, matchTypeConfig, err = parseExpression(expression)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, exampleValue)
					}

					matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
					if err != nil {
						return content, err
					}
					rules[rulesPath+key] = matchingRules

					if unionType, ok := isUnion(key, schema); ok {
						content[key] = map[string]any{unionType: contentList}
					}
				}
			}
		}
	}
	// log.Println("Parsed content:", content)
	// log.Println("Parsed Rules: ", rules)
	return content, err
}

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
	binary, err := avro.Marshal(schema, data)
	if err != nil {
		log.Println("ERROR while trying to marshal from Native to Binary. ", err.Error())
		return nil, err
	}

	log.Println("Native content: ", data)

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
