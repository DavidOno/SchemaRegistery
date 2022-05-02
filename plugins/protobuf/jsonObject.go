package main

import "fmt"

type JsonObject struct {
	Name     string
	Elements []Json
}

func (jsonObject JsonObject) Append() {
	jsonDoc += fmt.Sprintf("\n\"%s\": {", jsonObject.Name)
	for index, element := range jsonObject.Elements {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append()
	}
	jsonDoc += "}"
}
