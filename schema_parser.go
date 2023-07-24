package main

// import "log"

// type Field struct {
// 	Name string      `json:"name"`
// 	Type interface{} `json:"type"`
// 	Fields []interface{} `json:"fields"`
// 	Symbols []interface{} `json:"symbols"`
// 	Items interface{}	`json:"items"`
// 	Values interface{}	`json:"values"` 
// 	Size int		`json:"size"`
// }

// func traverseFields(fields []Field, path string) {
// 	for _, field := range fields {
// 		switch fieldType := field.Type.(type) {
// 		case string:			
// 			switch field.Type {
// 			case "record":
// 				traverseFields(parseFields(field.Fields), path+field.Name+".")	
// 			case "array":
// 				switch field.Items.(type) {
// 				case string:
// 					log.Println("Path:",path, ",Name:",field.Name, ",Type:",field.Type, ",Items:", field.Items, ",Values:", field.Values, "Symbols:", field.Symbols)
// 				case []interface{}:					
// 					traverseFields(parseFields(field.Items.([]interface{})), path+field.Name+".")	
// 				case map[string]interface{}:
// 					traverseFields(parseMapFields(field.Items.(map[string]interface{})), path+field.Name+".")
// 				}
// 			case "values":
// 				switch field.Values.(type) {
// 				case string:
// 					log.Println("Path:",path, ",Name:",field.Name, ",Type:",field.Type, ",Items:", field.Items, ",Values:", field.Values, "Symbols:", field.Symbols)
// 				case []interface{}:					
// 					traverseFields(parseFields(field.Items.([]interface{})), path+field.Name+".")	
// 				case map[string]interface{}:
// 					traverseFields(parseMapFields(field.Items.(map[string]interface{})), path+field.Name+".")
// 				}
// 			// case "enum":
// 			// 	log.Println("Path:",path, ",Name:",field.Name, ",Type:",field.Type, ",Items:", field.Items, ",Values:", field.Values, "Symbols:", field.Symbols)
// 			default:
// 				log.Println("Path:",path, ",Name:",field.Name, ",Type:",field.Type, ",Items:", field.Items, ",Values:", field.Values, "Symbols:", field.Symbols)
// 			}
// 		case []interface{}:
// 			for _, subType := range fieldType {
// 				switch subType := subType.(type) {
// 				case string:
// 					log.Println("Path:",path, ",Name:",field.Name, ",Type:",field.Type, ",Items:", field.Items, ",Values:", field.Values, "Symbols:", field.Symbols)
// 				case map[string]interface{}:
// 					traverseFields(parseMapFields(subType), path+field.Name+".")	
// 				}
// 			}
// 		case map[string]interface{}:
// 			traverseFields(parseMapFields(fieldType), path+field.Name+".")			
// 		}
// 	}
// }

// func parseFields(fields []interface{}) []Field {
// 	var parsedFields []Field
// 	for _, field := range fields {
// 		if fieldMap, ok := field.(map[string]interface{}); ok {
// 			name, _ := fieldMap["name"].(string)
// 			items, _ := fieldMap["items"].(string)
// 			fieldsArray, _ := fieldMap["fields"].([]interface{})
// 			symbols, _ := fieldMap["symbols"].([]interface{})
// 			values, _ := fieldMap["values"].(string)
// 			parsedFields = append(parsedFields, Field{
// 				Name: name,
// 				Type: fieldMap["type"],
// 				Fields: fieldsArray,
// 				Symbols: symbols,
// 				Items: items,
// 				Values: values,
// 			})
// 		}
// 	}
// 	// log.Println("PARSED fields:", parsedFields)
// 	return parsedFields
// }

// func parseMapFields(fields map[string]interface{}) []Field {
// 	var parsedMapFields []Field

// 	name, _ := fields["name"].(string)	
// 	items := fields["items"]
// 	fieldsArray, _ := fields["fields"].([]interface{})
// 	symbols, _ := fields["symbols"].([]interface{})
// 	values := fields["values"]
// 	size, _ := fields["size"].(int)
// 	parsedMapFields = append(parsedMapFields, Field{
// 		Name: name,
// 		Type: fields["type"],
// 		Fields: fieldsArray,
// 		Symbols: symbols,
// 		Items: items,
// 		Values: values,
// 		Size: size,
// 	})
// 	// log.Println("PARSED fields:", parsedMapFields)
// 	return parsedMapFields
// }


// func findField(avroMap map[string]interface{}, fieldName string) interface{} {
// 	for key, value := range avroMap {
// 		log.Println("Key Value:", key, value)
// 		if key == fieldName {
// 			return value
// 		}

// 		if nestedMap, ok := value.(map[string]interface{}); ok {
// 			if field := findField(nestedMap, fieldName); field != nil {
// 				return field
// 			}
// 		}
// 	}

// 	return nil
// }
