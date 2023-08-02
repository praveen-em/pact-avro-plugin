package main

import (
	"strings"
)

type Field struct {
	Name        string        `json:"name"`
	Type        interface{}   `json:"type"`
	LogicalType string        `json:"logicalType"`
	Fields      []interface{} `json:"fields"`
	Symbols     []interface{} `json:"symbols"`
	Items       interface{}   `json:"items"`
	Values      interface{}   `json:"values"`
	Size        int           `json:"size"`
}

var schemaTypesContainer = make(map[string]any)
var schemaUnionContainer = make(map[string][]any)

func iterateSchemaFields(fields []Field, path string, isUnion bool) {
	for _, field := range fields {
		switch fieldType := field.Type.(type) {
		case string:
			switch field.Type {
			case "record":
				addToTypesContainer(path+field.Name, "record", false)
				if isUnion {
					addToTypesContainer(path, "record", true)
				}
				iterateSchemaFields(parseFields(field.Fields), path, false)
			case "array":
				switch field.Items.(type) {
				case string:
					// log.Println("Path:", path, ",Name:", field.Name, ",Type:", field.Type, ",Items:", field.Items)
					addToTypesContainer(path+field.Name, field.Type.(string)+"."+field.Items.(string)+getLogicalType(field.LogicalType), isUnion)
				case []interface{}:
					iterateSchemaFields(parseFields(field.Items.([]interface{})), path, false)
				case map[string]interface{}:
					if isUnion {
						addToTypesContainer(path, "array.record", true)
					}
					iterateSchemaFields(parseMapFields(field.Items.(map[string]interface{})), path, false)
				}
			case "map":
				switch field.Values.(type) {
				case string:
					// log.Println("Path:", path, ",Name:", field.Name, ",Type:", field.Type, ",Values:", field.Values)
					addToTypesContainer(path+field.Name, field.Type.(string)+"."+field.Values.(string)+getLogicalType(field.LogicalType), isUnion)
				case []interface{}:
					iterateSchemaFields(parseFields(field.Items.([]interface{})), path, false)
				case map[string]interface{}:
					if isUnion {
						addToTypesContainer(path, "map.record", true)
					}
					iterateSchemaFields(parseMapFields(field.Items.(map[string]interface{})), path, false)
				}
			case "enum":
				// log.Println("Path:", path, ",Name:", field.Name, ",Type:", field.Type)
				addToTypesContainer(path, field.Type.(string), isUnion)
			default:
				// log.Println("Path:", path, ",Name:", field.Name, ",Type:", field.Type)
				addToTypesContainer(path+field.Name, field.Type.(string)+getLogicalType(field.LogicalType), isUnion)
			}
		case []interface{}:
			addToTypesContainer(path+field.Name, "union", false)
			for _, subType := range fieldType {
				switch subType := subType.(type) {
				case string:
					// log.Println("Path:", path, ",Name:", "union."+field.Name, ",Type:", subType)
					addToTypesContainer(path+field.Name, subType, true)
				case []interface{}:
					iterateSchemaFields(parseFields(subType), path+field.Name+".", true)
				case map[string]interface{}:
					iterateSchemaFields(parseMapFields(subType), path+field.Name+".", true)
				}
			}
		case map[string]interface{}:
			iterateSchemaFields(parseMapFields(fieldType), path+field.Name+".", false)
		}
	}
}

func addToTypesContainer(fieldName string, fieldType string, isUnion bool) {
	if isUnion {
		schemaUnionContainer[strings.TrimSuffix(fieldName, ".")] = append(schemaUnionContainer[strings.TrimSuffix(fieldName, ".")], fieldType)
	} else {
		schemaTypesContainer[strings.TrimSuffix(fieldName, ".")] = fieldType
	}
}

func getLogicalType(logicalType string) string {
	if logicalType != "" {
		logicalType = "." + logicalType
	}
	return logicalType
}

func parseFields(fields []interface{}) []Field {
	var parsedFields []Field
	for _, field := range fields {
		if fieldMap, ok := field.(map[string]interface{}); ok {
			name, _ := fieldMap["name"].(string)
			items, _ := fieldMap["items"].(string)
			fieldsArray, _ := fieldMap["fields"].([]interface{})
			symbols, _ := fieldMap["symbols"].([]interface{})
			values, _ := fieldMap["values"].(string)
			logicalType, _ := fieldMap["logicalType"].(string)
			parsedFields = append(parsedFields, Field{
				Name:        name,
				Type:        fieldMap["type"],
				LogicalType: logicalType,
				Fields:      fieldsArray,
				Symbols:     symbols,
				Items:       items,
				Values:      values,
			})
		}
	}
	// log.Println("PARSED fields:", parsedFields)
	return parsedFields
}

func parseMapFields(fields map[string]interface{}) []Field {
	var parsedMapFields []Field

	name, _ := fields["name"].(string)
	items := fields["items"]
	fieldsArray, _ := fields["fields"].([]interface{})
	symbols, _ := fields["symbols"].([]interface{})
	values := fields["values"]
	size, _ := fields["size"].(int)
	logicalType, _ := fields["logicalType"].(string)
	parsedMapFields = append(parsedMapFields, Field{
		Name:        name,
		Type:        fields["type"],
		LogicalType: logicalType,
		Fields:      fieldsArray,
		Symbols:     symbols,
		Items:       items,
		Values:      values,
		Size:        size,
	})
	// log.Println("PARSED fields:", parsedMapFields)
	return parsedMapFields
}
