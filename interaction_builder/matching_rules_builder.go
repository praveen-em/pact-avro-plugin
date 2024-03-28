package interaction

import (
	"log"
	plugin "github.com/praveen-em/pact-avro-plugin/io_pact_plugin"
	"google.golang.org/protobuf/types/known/structpb"
)

func buildMatchingRules(matchTypeMap map[string]interface{}, matchTypeConfigMap map[string]interface{}, exampleValueMap map[string]interface{}) (*plugin.MatchingRules, error) {
	matchingRules := &plugin.MatchingRules{}	
	var err error

	for k, v := range exampleValueMap {
		var matchType, matchTypeConfig string
		switch k {
		case "eachKey", "eachValue":
			if _ , ok := exampleValueMap["reference"]; !ok {
				eachKeyEachValueMap := make(map[string]interface{})
				eachKeyEachValueMap[k] = v
				matchType, _ = matchTypeMap[k].(string)
				matchTypeConfig, _ := matchTypeConfigMap[k].(string)
				matchingRule, err := buildMatchingRule(matchType, matchTypeConfig, eachKeyEachValueMap)
				if err != nil {
					return matchingRules, err
				}
				matchingRules.Rule = append(matchingRules.Rule, matchingRule)
			}

		default: 
			matchType, _ = matchTypeMap[k].(string)
			matchTypeConfig, _ = matchTypeConfigMap[k].(string)
			matchingRule, err := buildMatchingRule(matchType, matchTypeConfig, nil)
			if err != nil {
				return matchingRules, err
			}
			matchingRules.Rule = append(matchingRules.Rule, matchingRule)
		}

	}
	
	return matchingRules, err
}

func buildMatchingRule(matchType string, matchTypeConfig string, eachKeyEachValueMap map[string]interface{}) (matchingRule *plugin.MatchingRule, err error) {
	valuesInner := make(map[string]interface{})
	valuesOuter := make(map[string]interface{})
	var list []interface{}
	matchingRule = &plugin.MatchingRule{}
	matchingRule.Type = matchType

	valuesInner["match"] = matchType
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
		valuesInner[matchTypeConfigKey] = matchTypeConfig
	}

	if v, ok := eachKeyEachValueMap["eachKey"]; ok {
		matchingRule.Type = "eachKey"
		valuesOuter["value"] = v
		valuesOuter["rules"] = append(list, valuesInner)
	} else if v, ok := eachKeyEachValueMap["eachValue"]; ok {
		matchingRule.Type = "eachValue"
		valuesOuter["value"] = v
		valuesOuter["rules"] = append(list, valuesInner)
	} else {
		valuesOuter = valuesInner	
	}

	matchingRule.Values, err = structpb.NewStruct(valuesOuter)
	if err != nil {
		log.Println("ERROR while converting from native go to structpb format. ", err)
		return matchingRule, err
	}
	// log.Println("Matching Rule: ", matchingRule)
	return matchingRule, err

}
