package main

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hamba/avro/v2"
	plugin "github.com/praveen-em/pact-foobar-plugin/io_pact_plugin"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/types/known/structpb"
)

type configuration struct {
	content       map[string]interface{}
	contentBinary []byte
	rules         map[string]*plugin.MatchingRules
	schema        avro.Schema
}

func parseContentsConfig(req *plugin.ConfigureInteractionRequest) (*configuration, error) {
	records := req.GetContentsConfig().GetFields()

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

	content, rules, err := parseRecords(records, schema)
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

func parseRecords(records map[string]*structpb.Value, schema avro.Schema) (content map[string]interface{}, rules map[string]*plugin.MatchingRules, err error) {
	content = make(map[string]interface{})
	rules = make(map[string]*plugin.MatchingRules)
	var (
		exampleValue    interface{}
		matchType       string
		matchTypeConfig string
	)

	for key, value := range records {
		log.Println("ContentsConfig interated:", key, value)

		if !strings.HasPrefix(key, "pact:") {
			switch value.Kind.(type) {
			case *structpb.Value_StringValue:
				expression := value.GetStringValue()

				content[key], matchType, matchTypeConfig, err = parseExpression(expression)
				if err != nil {
					return content, rules, err
				}

				matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
				if err != nil {
					return content, rules, err
				}

				rules["$."+key] = matchingRules

			case *structpb.Value_StructValue:
				var contentMap map[string]interface{}
				contentMap, _, err = parseRecords(value.GetStructValue().Fields, schema)
				if err != nil {
					return content, rules, err
				}

				if isUnion(key, schema) {
					content[key] = map[string]any{"map": contentMap}
				} else {
					content[key] = contentMap
				}

			case *structpb.Value_ListValue:
				var contentList []interface{}
				valuesList := value.GetListValue().Values
				for index := range valuesList {
					// log.Println("List value: ", valuesList[index])
					expression := valuesList[index].GetStringValue()

					exampleValue, matchType, matchTypeConfig, err = parseExpression(expression)
					if err != nil {
						return content, rules, err
					}

					contentList = append(contentList, exampleValue)
					matchingRules, err := buildMatchingRules(matchType, matchTypeConfig)
					if err != nil {
						return content, rules, err
					}
					rules["$."+key] = matchingRules
				}

				if isUnion(key, schema) {
					content[key] = map[string]any{"array": contentList}
				}
			}
		}
	}
	log.Println("Parsed content:", content)
	return content, rules, err
}

func isUnion(key string, schema avro.Schema) (bool) {
	extract := "fields.#(name=\"" + key + "\").type"
	result := gjson.Get(schema.String(), extract)
	return reflect.TypeOf(result.Value()).Kind() == reflect.Slice 
}

func buildMatchingRules(matchType string, matchTypeConfig string) (*plugin.MatchingRules, error) {
	matchingRules := &plugin.MatchingRules{}

	matchingRule, err := getMatchingRule(matchType, matchTypeConfig)
	if err != nil {
		return matchingRules, err
	}

	matchingRules.Rule = append(matchingRules.Rule, matchingRule)
	return matchingRules, err
}

func parseExpression(expression string) (exampleValue interface{}, matchType string, matchTypeConfig string, err error) {
	matchType, matchTypeConfig, exampleValue = parseMatchingRuleDefinition(expression)
	exampleValue, err = formatDateTime(matchType, matchTypeConfig, exampleValue)
	if err != nil {
		return exampleValue, matchType, matchTypeConfig, err
	}
	return exampleValue, matchType, matchTypeConfig, err
}

func getMatchingRule(matchType string, matchTypeConfig string) (matchingRule *plugin.MatchingRule, err error) {
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

	return binary, err
}

func formatDateTime(matchType string, matchTypeConfig string, exampleValue interface{}) (interface{}, error) {
	//TODO: cover all date types
	var err error
	switch {
	case matchType == "date":
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
