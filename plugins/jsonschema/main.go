package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var input map[string]interface{}
var ddm map[string]interface{}

type MapType map[string]interface{}
type ArrayType []interface{}

func main() {
	ddm = make(map[string]interface{})
	readFromFile()
	mapJSS2DomainDataModel()
	writeToFile(ddm)
}

func readFromFile() {
	file, err := os.ReadFile("./test.json")
	check(err)
	json.Unmarshal(file, &input)
}

func writeToFile(input map[string]interface{}) {
	f, err := os.Create("jss2dc_test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bytes, _ := json.MarshalIndent(input, "", "\t")
	_, err2 := f.Write(bytes)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func mapJSS2DomainDataModel() {
	mapName()
	mapSchemaSpec()
	mapDescription()
	mapComponents()
}

func mapComponents() {
	components := make([]map[string]interface{}, 0)

	props := input["properties"].(map[string]interface{})
	component := make(map[string]interface{})
	objectType := mapTypeOfObject(component)
	properties := component[objectType].(map[string]interface{})
	properties["name"] = input["title"]
	properties["description"] = input["description"]
	fields := make([]map[string]interface{}, 0)
	for propKey, propFields := range props {
		field := make(map[string]interface{})
		field["name"] = propKey
		field["type"] = mapType(propFields)
		field["description"] = getValueFromMap(propFields, "description")
		field["optional"] = mapIfPropertyIsOptional(propKey)
		field["min"] = getMinProperty(propFields)
		field["max"] = getMaxProperty(propFields)
		mapUniqueField(propFields, field)
		fields = append(fields, field)
	}
	properties["fields"] = fields
	components = append(components, component)
	ddm["components"] = components
}

func mapUniqueField(propFields interface{}, field map[string]interface{}) {
	if value, ok := propFields.(map[string]interface{})["uniqueItems"]; ok {
		field["unique"] = value
	}
}

func getMaxProperty(propFields interface{}) interface{} {
	if isArray(propFields) {
		return findMaxValueDefinitionForArray(propFields)
	} else {
		return findMaxValueDefinition(propFields)
	}
}

func getMinProperty(propFields interface{}) interface{} {
	if isArray(propFields) {
		return findMinValueDefinitionForArray(propFields)
	} else {
		return findMinValueDefinition(propFields)
	}
}

func findMaxValueDefinitionForArray(propFields interface{}) interface{} {
	properties := propFields.(map[string]interface{})
	if value, ok := properties["maxItems"]; ok {
		return fmt.Sprintf("%v", value)
	}
	return "*"
}

func findMinValueDefinitionForArray(propFields interface{}) interface{} {
	properties := propFields.(map[string]interface{})
	if value, ok := properties["minItems"]; ok {
		return fmt.Sprintf("%v", value)
	}
	return "0"
}

func findMaxValueDefinition(propFields interface{}) interface{} {
	properties := propFields.(map[string]interface{})
	if value, ok := properties["exclusiveMaximum"]; ok {
		return fmt.Sprintf("%v[", value)
	} else if value, ok := properties["maximum"]; ok {
		max := value.(int)
		return fmt.Sprintf("%v", max)
	}
	return nil
}

func findMinValueDefinition(propFields interface{}) interface{} {
	properties := propFields.(map[string]interface{})
	if value, ok := properties["exclusiveMinimum"]; ok {
		return fmt.Sprintf("]%v", value)
	} else if value, ok := properties["minimum"]; ok {
		min := value.(int)
		return fmt.Sprintf("%v", min)
	}
	return nil
}

func isArray(propFields interface{}) bool {
	typeOfProperty := getValueFromMap(propFields, "type")
	if typeOfProperty == "array" {
		return true
	}
	return false
}

func getValueFromMap(mapping interface{}, keys ...string) interface{} {
	result := mapping
	for _, key := range keys {
		result = result.(map[string]interface{})[key]
	}
	return result
}

func mapType(propValue interface{}) string {
	typeOfField := getValueFromMap(propValue, "type").(string)
	if typeOfField == "array" {
		return getValueFromMap(propValue, "items", "type").(string)
	} else {
		return typeOfField
	}
}

func mapIfPropertyIsOptional(propKey string) bool {
	listOfRequiredProps := input["required"].([]interface{})
	set := make(map[interface{}]bool)
	for _, v := range listOfRequiredProps {
		set[v] = true
	}
	return !set[propKey]
}

func mapTypeOfObject(component map[string]interface{}) string {
	if input["type"] == "object" {
		component["object"] = make(map[string]interface{})
		return "object"
	} else {
		panic("Could not determine type")
	}
}

func mapDescription() {
	ddm["description"] = input["description"]
}

func mapSchemaSpec() {
	ddm["schemaSpec"] = "json-schema"
}

func mapName() {
	ddm["name"] = "test" //TODO: Improve that
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
