package configureinteraction

import (
	"errors"
	"log"

	"github.com/hamba/avro/v2"
	plugin "github.com/praveen-em/pact-avro-plugin/io_pact_plugin"
	"google.golang.org/protobuf/types/known/structpb"
)

type Configuration struct {
	Content       map[string]interface{}
	ContentBinary []byte
	Rules         map[string]*plugin.MatchingRules
	Schema        avro.Schema
}

var rules = make(map[string]*plugin.MatchingRules)

func ParseContentsConfig(ContentsConfig *structpb.Struct) (*Configuration, error) {
	records := ContentsConfig.GetFields()
	rulesPath := "$."

	if _, ok := records["pact:schema"]; !ok {
		err := errors.New("config item with key 'pact:schema' and the avro schema string is required")
		log.Println(err.Error())
		return nil, err
	}

	schema, err := avro.Parse(records["pact:schema"].GetStringValue())
	if err != nil {
		log.Println("ERROR parsing schema. ", err.Error())
		return nil, err
	}
	log.Println("Schema: ", schema)

	content, err := buildInteraction(records, schema, rules, rulesPath)
	if err != nil {
		return nil, err
	}

	contentBinary, err := convertNativeToAvroBinary(schema, content)
	if err != nil {
		return nil, err
	}

	var parsedData = &Configuration{content, contentBinary, rules, schema}
	return parsedData, err
}
