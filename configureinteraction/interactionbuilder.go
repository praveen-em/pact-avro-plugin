package configureinteraction

import (
	"reflect"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/hamba/avro/v2"
	plugin "github.com/praveen-em/pact-avro-plugin/io_pact_plugin"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/types/known/structpb"
)

var avroPrimitiveTypes = []string{"null", "boolean", "int", "long", "float", "double", "bytes", "string"}
var referenceList []string

func buildInteraction(records map[string]*structpb.Value, schema avro.Schema) (content map[string]interface{}, rules map[string]*plugin.MatchingRules, err error) {	
	rules = make(map[string]*plugin.MatchingRules)
	rulesPath := "$."
	content, err = buildInteractionWithRules(records, schema, rules, rulesPath)
	return content, rules, err
}

func buildInteractionWithRules(records map[string]*structpb.Value, schema avro.Schema, rules map[string]*plugin.MatchingRules, rulesPath string) (content map[string]interface{}, err error) {
	content = make(map[string]interface{})
	var (
		exampleValueMap    map[string]interface{}
		exampleValue       interface{}
		matchTypeMap       map[string]interface{}
		matchTypeConfigMap map[string]interface{}
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

				exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err = parseExpression(expression)
				if err != nil {
					return content, err
				}

				matchingRules, err := buildMatchingRules(matchTypeMap, matchTypeConfigMap, exampleValueMap)
				if err != nil {
					return content, err
				}

				switch matchTypeMap["reference"] {
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

				contentMap, err = buildInteractionWithRules(value.GetStructValue().Fields, schema, rules, rulesPath)
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
						contentMap, err = buildInteractionWithRules(valuesList[index].GetStructValue().Fields, schema, rules, rulesPath)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, contentMap)
						rulesPath = strings.TrimSuffix(rulesPath, key+".")

					case *structpb.Value_StringValue:
						expression := valuesList[index].GetStringValue()

						exampleValueMap, exampleValue, matchTypeMap, matchTypeConfigMap, err = parseExpression(expression)
						if err != nil {
							return content, err
						}

						contentList = append(contentList, exampleValue)
						matchingRules, err := buildMatchingRules(matchTypeMap, matchTypeConfigMap, exampleValueMap)
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
