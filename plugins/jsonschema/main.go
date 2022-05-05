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
	fmt.Println(props)
	component := make(map[string]interface{})
	fmt.Println(len(component))
	objectType := mapTypeOfObject(component)
	fmt.Println(len(component))
	properties := component[objectType].(map[string]interface{})
	properties["name"] = input["title"]
	properties["description"] = input["description"]
	fields := make([]map[string]interface{}, 0)
	for propKey, propValue := range props {
		field := make(map[string]interface{})
		field["name"] = propKey
		field["type"] = mapType(propValue)
		field["description"] = getValueFromMap(propValue, "description")
		field["optional"] = mapIfPropertyIsOptional(propKey)
		fields = append(fields, field)
	}
	properties["fields"] = fields
	components = append(components, component)
	ddm["components"] = components
}

func getValueFromMap(mapping interface{}, key string) interface{} {
	return mapping.(map[string]interface{})[key]
}

func mapType(propValue interface{}) string {
	typeOfField := getValueFromMap(propValue, "type").(string)
	if typeOfField == "array" {
		return getValueFromMap(getValueFromMap(propValue, "items"), "type").(string)
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
