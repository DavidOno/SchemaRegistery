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
	for propKey, propValue := range props {
		properties := component[objectType].(map[string]interface{})
		properties["name"] = input["title"]
		properties["description"] = input["description"]
		properties["fields"] = make([]map[string]interface{}, 0)
		fields := properties["fields"].(map[string]interface{})
		fields["name"] = propKey
		fields["type"] = propValue.(map[string]interface{})["type"]
		fields["description"] = propValue.(map[string]interface{})["description"]

		// typeOfValue := reflect.TypeOf(propValue)
		// fmt.Println(typeOfValue)
		fmt.Println(propKey)
		fmt.Println(propValue)
		// switch typeOfValue.Kind() {
		// case reflect.Map:

		// if mv, ok := propValue.(MapType); ok {
		// 	// docMap[propKey] = doc.throughMap(mv)
		// 	fmt.Println(mv)
		// } else {
		// 	panic("error when casting to MapType")
		// }
		// case reflect.Array, reflect.Slice:
		// if mv, ok := propValue.(ArrayType); ok {
		// 	docMap[propKey] = doc.throughArray(mv)
		// } else {
		// 	panic("error when casting to ArrayType")
		// }
		// default:
		// docMap[propKey] = doc.processType(propValue)
		// }
	}
	components = append(components, component)
	ddm["components"] = components
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
